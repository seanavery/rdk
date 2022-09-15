package datamanager_test

import (
	"context"
	"image"
	"io"
	"io/fs"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/edaniels/golog"
	"github.com/edaniels/gostream"
	"github.com/pkg/errors"
	v1 "go.viam.com/api/app/datasync/v1"
	"go.viam.com/test"
	"go.viam.com/utils/rpc"

	"go.viam.com/rdk/components/arm"
	"go.viam.com/rdk/components/camera"
	"go.viam.com/rdk/config"
	commonpb "go.viam.com/rdk/proto/api/common/v1"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/robot"
	"go.viam.com/rdk/services/datamanager"
	"go.viam.com/rdk/services/datamanager/datacapture"
	"go.viam.com/rdk/services/datamanager/datasync"
	"go.viam.com/rdk/services/datamanager/internal"
	"go.viam.com/rdk/testutils/inject"
	rutils "go.viam.com/rdk/utils"
)

const (
	captureWaitTime = time.Millisecond * 25
	syncWaitTime    = time.Millisecond * 100
)

var (
	// Robot config which specifies data manager service.
	configPath = "services/datamanager/data/fake_robot_with_data_manager.json"

	// 0.0041 mins is 246 milliseconds, this is the interval waiting time in the config file used for testing.
	configSyncIntervalMins = 0.0041

	syncIntervalMins   = 0.0041 // 250ms
	captureDir         = "/tmp/capture"
	armDir             = captureDir + "/arm/arm1/GetEndPosition"
	emptyFileBytesSize = 30 // size of leading metadata message
	testSvcName1       = "svc1"
	testSvcName2       = "svc2"
)

const (
	testDataManagerServiceName = "DataManager1"
)

// readDir filters out folders from a slice of FileInfos.
func readDir(t *testing.T, dir string) ([]fs.DirEntry, error) {
	t.Helper()
	filesAndFolders, err := os.ReadDir(dir)
	if err != nil {
		t.Log(err)
	}
	var onlyFiles []fs.DirEntry
	for _, s := range filesAndFolders {
		if !s.IsDir() {
			onlyFiles = append(onlyFiles, s)
		}
	}
	return onlyFiles, err
}

func resetFolder(t *testing.T, path string) {
	t.Helper()
	if err := os.RemoveAll(path); err != nil {
		t.Log(err)
	}
}

func getInjectedRobotWithArm(armKey string) *inject.Robot {
	r := &inject.Robot{}
	rs := map[resource.Name]interface{}{}
	injectedArm := &inject.Arm{}
	injectedArm.GetEndPositionFunc = func(ctx context.Context, extra map[string]interface{}) (*commonpb.Pose, error) {
		return &commonpb.Pose{X: 1, Y: 2, Z: 3}, nil
	}
	rs[arm.Named(armKey)] = injectedArm
	r.MockResourcesFromMap(rs)
	return r
}

func getInjectedRobotWithCamera(t *testing.T) *inject.Robot {
	t.Helper()
	r := &inject.Robot{}
	rs := map[resource.Name]interface{}{}

	img := image.NewNRGBA64(image.Rect(0, 0, 4, 4))
	injectCamera := &inject.Camera{}
	var imageReleasedMu sync.Mutex
	injectCamera.StreamFunc = func(ctx context.Context, errHandlers ...gostream.ErrorHandler) (gostream.VideoStream, error) {
		return gostream.NewEmbeddedVideoStreamFromReader(gostream.VideoReaderFunc(func(ctx context.Context) (image.Image, func(), error) {
			imageReleasedMu.Lock()
			time.Sleep(10 * time.Nanosecond)
			imageReleasedMu.Unlock()
			return img, func() {}, nil
		})), nil
	}

	rs[camera.Named("c1")] = injectCamera
	r.MockResourcesFromMap(rs)
	return r
}

func newTestDataManager(t *testing.T, localArmKey, remoteArmKey string) internal.DMService {
	t.Helper()
	dmCfg := &datamanager.Config{}
	cfgService := config.Service{
		Type:                "data_manager",
		ConvertedAttributes: dmCfg,
	}
	logger := golog.NewTestLogger(t)

	// Create local robot with injected arm.
	r := getInjectedRobotWithArm(localArmKey)

	// If passed, create remote robot with an injected arm.
	if remoteArmKey != "" {
		remoteRobot := getInjectedRobotWithArm(remoteArmKey)

		r.RemoteByNameFunc = func(name string) (robot.Robot, bool) {
			return remoteRobot, true
		}
	}

	svc, err := datamanager.New(context.Background(), r, cfgService, logger)
	if err != nil {
		t.Log(err)
	}
	return svc.(internal.DMService)
}

func setupConfig(t *testing.T, relativePath string) *config.Config {
	t.Helper()
	logger := golog.NewTestLogger(t)
	testCfg, err := config.Read(context.Background(), rutils.ResolveFile(relativePath), logger)
	test.That(t, err, test.ShouldBeNil)
	testCfg.Cloud = &config.Cloud{ID: "part_id"}
	return testCfg
}

func TestNewDataManager(t *testing.T) {
	dmsvc := newTestDataManager(t, "arm1", "")
	testCfg := setupConfig(t, configPath)

	// Empty config at initialization.
	captureDir := "/tmp/capture"
	defer resetFolder(t, captureDir)
	err := dmsvc.Update(context.Background(), testCfg)
	test.That(t, err, test.ShouldBeNil)
	captureTime := time.Millisecond * 100
	time.Sleep(captureTime)

	err = dmsvc.Close(context.Background())
	test.That(t, err, test.ShouldBeNil)

	// Check that a collector wrote to file.
	armDir := captureDir + "/arm/arm1/GetEndPosition"
	filesInArmDir, err := readDir(t, armDir)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(filesInArmDir), test.ShouldEqual, 1)
	oldInfo, err := filesInArmDir[0].Info()
	test.That(t, err, test.ShouldBeNil)
	oldSize := oldInfo.Size()
	test.That(t, oldSize, test.ShouldBeGreaterThan, emptyFileBytesSize)

	// Check that dummy tags "a" and "b" are being wrote to metadata.
	captureFileName := filesInArmDir[0].Name()
	file, err := os.Open(armDir + "/" + captureFileName)
	test.That(t, err, test.ShouldBeNil)
	md, err := datacapture.ReadDataCaptureMetadata(file)
	test.That(t, md.Tags[0], test.ShouldEqual, "a")
	test.That(t, md.Tags[1], test.ShouldEqual, "b")

	// When Close returns all background processes in svc should be closed, but still sleep for 100ms to verify
	// that there's not a resource leak causing writes to still happens after Close() returns.
	time.Sleep(captureTime)
	test.That(t, err, test.ShouldBeNil)
	filesInArmDir, err = readDir(t, armDir)
	test.That(t, err, test.ShouldBeNil)
	newInfo, err := filesInArmDir[0].Info()
	test.That(t, err, test.ShouldBeNil)
	newSize := newInfo.Size()
	test.That(t, oldSize, test.ShouldEqual, newSize)
}

func TestCaptureDisabled(t *testing.T) {
	// Empty config at initialization.
	captureDir := "/tmp/capture"
	dmsvc := newTestDataManager(t, "arm1", "")
	// Set capture parameters in Update.
	testCfg := setupConfig(t, configPath)
	dmCfg, err := getDataManagerConfig(testCfg)
	test.That(t, err, test.ShouldBeNil)

	defer resetFolder(t, captureDir)
	err = dmsvc.Update(context.Background(), testCfg)
	test.That(t, err, test.ShouldBeNil)
	time.Sleep(captureWaitTime)

	// Call Update with a disabled capture and give the collector time to write to file.
	dmCfg.CaptureDisabled = true
	err = dmsvc.Update(context.Background(), testCfg)
	test.That(t, err, test.ShouldBeNil)
	time.Sleep(captureWaitTime)

	// Verify that the collector wrote to its file.
	armDir := captureDir + "/arm/arm1/GetEndPosition"
	filesInArmDir, err := readDir(t, armDir)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(filesInArmDir), test.ShouldEqual, 1)
	info, err := filesInArmDir[0].Info()
	test.That(t, err, test.ShouldBeNil)
	test.That(t, info.Size(), test.ShouldBeGreaterThan, emptyFileBytesSize)

	// Re-enable capture.
	dmCfg.CaptureDisabled = false
	err = dmsvc.Update(context.Background(), testCfg)
	test.That(t, err, test.ShouldBeNil)
	time.Sleep(captureWaitTime)

	// Close service.
	err = dmsvc.Close(context.Background())
	test.That(t, err, test.ShouldBeNil)

	// Verify that started collection began in a new file when it was re-enabled.
	filesInArmDir, err = readDir(t, armDir)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(filesInArmDir), test.ShouldEqual, 2)

	// Verify that something different was written to both files.
	test.That(t, filesInArmDir[0], test.ShouldNotEqual, filesInArmDir[1])
	info, err = filesInArmDir[1].Info()
	test.That(t, err, test.ShouldBeNil)
	test.That(t, info.Size(), test.ShouldBeGreaterThan, emptyFileBytesSize)
}

func TestNewRemoteDataManager(t *testing.T) {
	// Empty config at initialization.
	captureDir := "/tmp/capture"
	dmsvc := newTestDataManager(t, "localArm", "remoteArm")

	// Set capture parameters in Update.
	conf := setupConfig(t, "services/datamanager/data/fake_robot_with_remote_and_data_manager.json")
	defer resetFolder(t, captureDir)
	err := dmsvc.Update(context.Background(), conf)
	test.That(t, err, test.ShouldBeNil)
	time.Sleep(captureWaitTime)

	// Verify that after close is called, the collector is no longer writing.
	err = dmsvc.Close(context.Background())
	test.That(t, err, test.ShouldBeNil)

	// Verify that the local and remote collectors wrote to their files.
	localArmDir := captureDir + "/arm/localArm/GetEndPosition"
	filesInLocalArmDir, err := readDir(t, localArmDir)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(filesInLocalArmDir), test.ShouldEqual, 1)
	info, err := filesInLocalArmDir[0].Info()
	test.That(t, err, test.ShouldBeNil)
	test.That(t, info.Size(), test.ShouldBeGreaterThan, 0)

	remoteArmDir := captureDir + "/arm/remoteArm/GetEndPosition"
	filesInRemoteArmDir, err := readDir(t, remoteArmDir)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(filesInRemoteArmDir), test.ShouldEqual, 1)
	info, err = filesInRemoteArmDir[0].Info()
	test.That(t, err, test.ShouldBeNil)
	test.That(t, info.Size(), test.ShouldBeGreaterThan, 0)
}

// Validates that if the datamanager/robot die unexpectedly, that previously captured but not synced files are still
// synced at start up.
func TestRecoversAfterKilled(t *testing.T) {
	// Register mock datasync service with a mock server.
	rpcServer, mockService := buildAndStartLocalServer(t)
	defer func() {
		err := rpcServer.Stop()
		test.That(t, err, test.ShouldBeNil)
	}()

	dirs, numArbitraryFilesToSync, err := populateAdditionalSyncPaths()
	defer func() {
		for _, dir := range dirs {
			resetFolder(t, dir)
		}
	}()
	defer resetFolder(t, captureDir)
	defer resetFolder(t, armDir)
	if err != nil {
		t.Error("unable to generate arbitrary data files and create directory structure for additionalSyncPaths")
	}

	testCfg := setupConfig(t, configPath)
	dmCfg, err := getDataManagerConfig(testCfg)
	test.That(t, err, test.ShouldBeNil)
	dmCfg.SyncIntervalMins = configSyncIntervalMins
	dmCfg.AdditionalSyncPaths = dirs

	// Initialize the data manager and update it with our config.
	dmsvc := newTestDataManager(t, "arm1", "")
	dmsvc.SetSyncerConstructor(getTestSyncerConstructor(t, rpcServer))
	dmsvc.SetWaitAfterLastModifiedSecs(10)
	err = dmsvc.Update(context.TODO(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// We set sync_interval_mins to be about 250ms in the config, so wait 150ms so data is captured but not synced.
	time.Sleep(time.Millisecond * 150)

	// Simulate turning off the service.
	err = dmsvc.Close(context.TODO())
	test.That(t, err, test.ShouldBeNil)

	// Validate nothing has been synced yet.
	test.That(t, len(mockService.getUploadedFiles()), test.ShouldEqual, 0)

	// Turn the service back on.
	dmsvc = newTestDataManager(t, "arm1", "")
	dmsvc.SetSyncerConstructor(getTestSyncerConstructor(t, rpcServer))
	dmsvc.SetWaitAfterLastModifiedSecs(0)
	err = dmsvc.Update(context.TODO(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// Validate that the previously captured file was uploaded at startup.
	time.Sleep(syncWaitTime)
	err = dmsvc.Close(context.TODO())
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(mockService.getUploadedFiles()), test.ShouldEqual, 1+numArbitraryFilesToSync)
}

// Validates that if the robot config file specifies a directory path in additionalSyncPaths that does not exist,
// that directory is created (and can be synced on subsequent iterations of syncing).
func TestCreatesAdditionalSyncPaths(t *testing.T) {
	td := "additional_sync_path_dir"
	// Once testing is complete, remove contents from data capture dirs.
	defer resetFolder(t, captureDir)
	defer resetFolder(t, armDir)
	defer resetFolder(t, td)

	// Register mock datasync service with a mock server.
	rpcServer, _ := buildAndStartLocalServer(t)
	defer func() {
		err := rpcServer.Stop()
		test.That(t, err, test.ShouldBeNil)
	}()

	testCfg := setupConfig(t, configPath)
	dmCfg, err := getDataManagerConfig(testCfg)
	test.That(t, err, test.ShouldBeNil)
	dmCfg.SyncIntervalMins = syncIntervalMins
	dmCfg.AdditionalSyncPaths = []string{td}

	// Initialize the data manager and update it with our config. The call to Update(ctx, conf) should create the
	// arbitrary sync paths directory it in the file system.
	dmsvc := newTestDataManager(t, "arm1", "")
	dmsvc.SetSyncerConstructor(getTestSyncerConstructor(t, rpcServer))
	dmsvc.SetWaitAfterLastModifiedSecs(0)
	err = dmsvc.Update(context.TODO(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// Validate the "additional_sync_path_dir" was created. Wait some time to ensure it would have been created.
	time.Sleep(syncWaitTime)
	_ = dmsvc.Close(context.TODO())
	_, err = os.Stat(td)
	test.That(t, errors.Is(err, nil), test.ShouldBeTrue)
}

// Generates and populates a directory structure of files that contain arbitrary file data. Used to simulate testing
// syncing of data in the service's additional_sync_paths.
//nolint
func populateAdditionalSyncPaths() ([]string, int, error) {
	var additionalSyncPaths []string
	numArbitraryFilesToSync := 0

	// Generate additional_sync_paths "dummy" dirs & files.
	for i := 0; i < 2; i++ {
		// Create a temp dir that will be in additional_sync_paths.
		td, err := os.MkdirTemp("", "additional_sync_path_dir_")
		if err != nil {
			return []string{}, 0, errors.New("cannot create temporary dir to simulate additional_sync_paths in data manager service config")
		}
		additionalSyncPaths = append(additionalSyncPaths, td)

		// Make the first dir empty.
		if i == 0 {
			continue
		} else {
			// Make the dirs that will contain two file.
			for i := 0; i < 2; i++ {
				// Generate data that will be in a temp file.
				fileData := []byte("This is file data. It will be stored in a directory included in the user's specified additional sync paths. Hopefully it is uploaded from the robot to the cloud!")

				// Create arbitrary file that will be in the temp dir generated above.
				tf, err := os.CreateTemp(td, "arbitrary_file_")
				if err != nil {
					return nil, 0, errors.New("cannot create temporary file to simulate uploading from data manager service")
				}

				// Write data to the temp file.
				if _, err := tf.Write(fileData); err != nil {
					return nil, 0, errors.New("cannot write arbitrary data to temporary file")
				}

				// Increment number of files to be synced.
				numArbitraryFilesToSync++
			}
		}
	}
	return additionalSyncPaths, numArbitraryFilesToSync, nil
}

func noRepeatedElements(slice []string) bool {
	visited := make(map[string]bool, 0)
	for i := 0; i < len(slice); i++ {
		if visited[slice[i]] {
			return false
		}
		visited[slice[i]] = true
	}
	return true
}

// Validates that manual syncing works for a datamanager.
func TestManualSync(t *testing.T) {
	// Register mock datasync service with a mock server.
	rpcServer, mockService := buildAndStartLocalServer(t)
	defer func() {
		err := rpcServer.Stop()
		test.That(t, err, test.ShouldBeNil)
	}()

	dirs, numArbitraryFilesToSync, err := populateAdditionalSyncPaths()
	defer func() {
		for _, dir := range dirs {
			resetFolder(t, dir)
		}
	}()
	defer resetFolder(t, captureDir)
	defer resetFolder(t, armDir)
	if err != nil {
		t.Error("unable to generate arbitrary data files and create directory structure for additionalSyncPaths")
	}
	testCfg := setupConfig(t, configPath)
	dmCfg, err := getDataManagerConfig(testCfg)
	test.That(t, err, test.ShouldBeNil)
	dmCfg.SyncIntervalMins = configSyncIntervalMins
	dmCfg.AdditionalSyncPaths = dirs

	// Initialize the data manager and update it with our config.
	dmsvc := newTestDataManager(t, "arm1", "")
	dmsvc.SetSyncerConstructor(getTestSyncerConstructor(t, rpcServer))
	dmsvc.SetWaitAfterLastModifiedSecs(0)
	err = dmsvc.Update(context.TODO(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// Run and upload files.
	err = dmsvc.Sync(context.Background())
	test.That(t, err, test.ShouldBeNil)
	time.Sleep(syncWaitTime)

	// Verify that one data capture file was uploaded, two additional_sync_paths files were uploaded,
	// and that no two uploaded files are the same.
	test.That(t, len(mockService.getUploadedFiles()), test.ShouldEqual, numArbitraryFilesToSync+1)
	test.That(t, noRepeatedElements(mockService.getUploadedFiles()), test.ShouldBeTrue)

	// Sync again and verify it synced the second data capture file, but also validate that it didn't attempt to resync
	// any files that were previously synced.
	err = dmsvc.Sync(context.Background())
	test.That(t, err, test.ShouldBeNil)
	time.Sleep(syncWaitTime)
	_ = dmsvc.Close(context.TODO())
	test.That(t, len(mockService.getUploadedFiles()), test.ShouldEqual, numArbitraryFilesToSync+2)
	test.That(t, noRepeatedElements(mockService.getUploadedFiles()), test.ShouldBeTrue)
}

// Validates that scheduled syncing works for a datamanager.
func TestScheduledSync(t *testing.T) {
	// Register mock datasync service with a mock server.
	rpcServer, mockService := buildAndStartLocalServer(t)
	defer func() {
		err := rpcServer.Stop()
		test.That(t, err, test.ShouldBeNil)
	}()

	dirs, numArbitraryFilesToSync, err := populateAdditionalSyncPaths()
	defer func() {
		for _, dir := range dirs {
			_ = os.RemoveAll(dir)
		}
	}()
	defer resetFolder(t, captureDir)
	defer resetFolder(t, armDir)
	if err != nil {
		t.Error("unable to generate arbitrary data files and create directory structure for additionalSyncPaths")
	}
	// Use config with 250ms sync interval.
	testCfg := setupConfig(t, configPath)
	dmCfg, err := getDataManagerConfig(testCfg)
	test.That(t, err, test.ShouldBeNil)
	dmCfg.SyncIntervalMins = configSyncIntervalMins
	dmCfg.AdditionalSyncPaths = dirs

	// Make the captureDir where we're logging data for our arm.
	captureDir := "/tmp/capture"
	armDir := captureDir + "/arm/arm1/GetEndPosition"

	// Clear the capture dir after we're done.
	defer resetFolder(t, armDir)

	// Initialize the data manager and update it with our config.
	dmsvc := newTestDataManager(t, "arm1", "")
	dmsvc.SetSyncerConstructor(getTestSyncerConstructor(t, rpcServer))
	dmsvc.SetWaitAfterLastModifiedSecs(0)
	err = dmsvc.Update(context.TODO(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// We set sync_interval_mins to be about 250ms in the config, so wait 600ms (more than two iterations of syncing)
	// for the additional_sync_paths files to sync AND for TWO data capture files to sync.
	time.Sleep(time.Millisecond * 600)
	_ = dmsvc.Close(context.TODO())

	// Verify that the additional_sync_paths files AND the TWO data capture files were uploaded.
	test.That(t, len(mockService.getUploadedFiles()), test.ShouldEqual, numArbitraryFilesToSync+2)
	test.That(t, noRepeatedElements(mockService.getUploadedFiles()), test.ShouldBeTrue)
}

// Validates that we can attempt a scheduled and manual syncDataCaptureFiles at the same time without duplicating files
// or running into errors.
func TestManualAndScheduledSync(t *testing.T) {
	// Register mock datasync service with a mock server.
	rpcServer, mockService := buildAndStartLocalServer(t)
	defer func() {
		err := rpcServer.Stop()
		test.That(t, err, test.ShouldBeNil)
	}()

	dirs, numArbitraryFilesToSync, err := populateAdditionalSyncPaths()
	defer func() {
		for _, dir := range dirs {
			resetFolder(t, dir)
		}
	}()
	defer resetFolder(t, captureDir)
	defer resetFolder(t, armDir)
	if err != nil {
		t.Error("unable to generate arbitrary data files and create directory structure for additionalSyncPaths")
	}
	testCfg := setupConfig(t, configPath)
	dmCfg, err := getDataManagerConfig(testCfg)
	test.That(t, err, test.ShouldBeNil)
	dmCfg.SyncIntervalMins = configSyncIntervalMins
	dmCfg.AdditionalSyncPaths = dirs

	// Make the captureDir where we're logging data for our arm.
	captureDir := "/tmp/capture"
	armDir := captureDir + "/arm/arm1/GetEndPosition"
	defer resetFolder(t, armDir)

	// Initialize the data manager and update it with our config.
	dmsvc := newTestDataManager(t, "arm1", "")
	dmsvc.SetSyncerConstructor(getTestSyncerConstructor(t, rpcServer))
	dmsvc.SetWaitAfterLastModifiedSecs(0)
	err = dmsvc.Update(context.TODO(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// Perform a manual and scheduled syncDataCaptureFiles at approximately the same time, then close the svc.
	time.Sleep(time.Millisecond * 250)
	err = dmsvc.Sync(context.TODO())
	test.That(t, err, test.ShouldBeNil)
	time.Sleep(syncWaitTime)
	_ = dmsvc.Close(context.TODO())

	// Verify that two data capture files were uploaded, two additional_sync_paths files were uploaded,
	// and that no two uploaded files are the same.
	test.That(t, len(mockService.getUploadedFiles()), test.ShouldEqual, numArbitraryFilesToSync+2)
	test.That(t, noRepeatedElements(mockService.getUploadedFiles()), test.ShouldBeTrue)

	// We've uploaded (and thus deleted) the first two files and should now be collecting a single new one.
	filesInArmDir, err := readDir(t, armDir)
	if err != nil {
		t.Fatalf("failed to list files in armDir")
	}
	test.That(t, len(filesInArmDir), test.ShouldEqual, 1)
}

func TestRegisteredReconfigurable(t *testing.T) {
	s := registry.ResourceSubtypeLookup(datamanager.Subtype)
	test.That(t, s, test.ShouldNotBeNil)
	r := s.Reconfigurable
	test.That(t, r, test.ShouldNotBeNil)
}

func TestWrapWithReconfigurable(t *testing.T) {
	svc := &mock{name: testSvcName1}
	reconfSvc1, err := datamanager.WrapWithReconfigurable(svc)
	test.That(t, err, test.ShouldBeNil)

	_, err = datamanager.WrapWithReconfigurable(nil)
	test.That(t, err, test.ShouldBeError, datamanager.NewUnimplementedInterfaceError(nil))

	reconfSvc2, err := datamanager.WrapWithReconfigurable(reconfSvc1)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, reconfSvc2, test.ShouldEqual, reconfSvc1)
}

func TestReconfigurable(t *testing.T) {
	actualSvc1 := &mock{name: testSvcName1}
	reconfSvc1, err := datamanager.WrapWithReconfigurable(actualSvc1)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, reconfSvc1, test.ShouldNotBeNil)

	actualArm2 := &mock{name: testSvcName2}
	reconfSvc2, err := datamanager.WrapWithReconfigurable(actualArm2)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, reconfSvc2, test.ShouldNotBeNil)
	test.That(t, actualSvc1.reconfCount, test.ShouldEqual, 0)

	err = reconfSvc1.Reconfigure(context.Background(), reconfSvc2)
	test.That(t, err, test.ShouldBeNil)
	test.That(t, reconfSvc1, test.ShouldResemble, reconfSvc2)
	test.That(t, actualSvc1.reconfCount, test.ShouldEqual, 1)

	err = reconfSvc1.Reconfigure(context.Background(), nil)
	test.That(t, err, test.ShouldNotBeNil)
	test.That(t, err, test.ShouldBeError, rutils.NewUnexpectedTypeError(reconfSvc1, nil))
}

type mock struct {
	datamanager.Service
	name        string
	reconfCount int
}

func (m *mock) Close(_ context.Context) error {
	m.reconfCount++
	return nil
}

func TestSyncDisabled(t *testing.T) {
	// Register mock datasync service with a mock server.
	rpcServer, mockService := buildAndStartLocalServer(t)
	defer func() {
		err := rpcServer.Stop()
		test.That(t, err, test.ShouldBeNil)
	}()

	testCfg := setupConfig(t, configPath)
	dmCfg, err := getDataManagerConfig(testCfg)
	test.That(t, err, test.ShouldBeNil)
	dmCfg.SyncIntervalMins = syncIntervalMins

	// Make the captureDir where we're logging data for our arm.
	captureDir := "/tmp/capture"
	armDir := captureDir + "/arm/arm1/"
	defer resetFolder(t, armDir)

	// Initialize the data manager and update it with our config.
	dmsvc := newTestDataManager(t, "arm1", "")
	dmsvc.SetSyncerConstructor(getTestSyncerConstructor(t, rpcServer))
	err = dmsvc.Update(context.TODO(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// We set sync_interval_mins to be about 250ms in the config, so wait 150ms so data is captured but not synced.
	time.Sleep(time.Millisecond * 150)

	// Simulate disabling sync.
	dmCfg.ScheduledSyncDisabled = true
	err = dmsvc.Update(context.Background(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// Validate nothing has been synced yet.
	test.That(t, len(mockService.getUploadedFiles()), test.ShouldEqual, 0)

	// Re-enable sync.
	dmCfg.ScheduledSyncDisabled = false
	err = dmsvc.Update(context.Background(), testCfg)
	test.That(t, err, test.ShouldBeNil)

	// We set sync_interval_mins to be about 250ms in the config, so wait 600ms and ensure three files were uploaded:
	// one from file immediately uploaded when sync was re-enabled and two after.
	time.Sleep(time.Millisecond * 600)
	err = dmsvc.Close(context.TODO())
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(mockService.getUploadedFiles()), test.ShouldEqual, 3)
}

func TestGetDurationFromHz(t *testing.T) {
	test.That(t, datamanager.GetDurationFromHz(0.1), test.ShouldEqual, time.Second*10)
	test.That(t, datamanager.GetDurationFromHz(0.5), test.ShouldEqual, time.Second*2)
	test.That(t, datamanager.GetDurationFromHz(1), test.ShouldEqual, time.Second)
	test.That(t, datamanager.GetDurationFromHz(1000), test.ShouldEqual, time.Millisecond)
	test.That(t, datamanager.GetDurationFromHz(0), test.ShouldEqual, 0)
}

func TestAdditionalParamsInConfig(t *testing.T) {
	conf := setupConfig(t, "services/datamanager/data/robot_with_cam_capture.json")
	r := getInjectedRobotWithCamera(t)

	dmCfg := &datamanager.Config{}
	cfgService := config.Service{
		Type:                "data_manager",
		ConvertedAttributes: dmCfg,
	}
	logger := golog.NewTestLogger(t)
	svc, err := datamanager.New(context.Background(), r, cfgService, logger)
	if err != nil {
		t.Log(err)
	}

	dmsvc := svc.(internal.DMService)

	defer resetFolder(t, captureDir)

	err = dmsvc.Update(context.Background(), conf)
	test.That(t, err, test.ShouldBeNil)
	time.Sleep(captureWaitTime)

	filesInCamDir, err := readDir(t, captureDir+"/camera/c1/Next")
	test.That(t, err, test.ShouldBeNil)
	test.That(t, len(filesInCamDir), test.ShouldEqual, 1)
	info, err := filesInCamDir[0].Info()
	test.That(t, err, test.ShouldBeNil)
	test.That(t, info.Size(), test.ShouldBeGreaterThan, emptyFileBytesSize)

	// Verify that after close is called, the collector is no longer writing.
	err = dmsvc.Close(context.Background())
	test.That(t, err, test.ShouldBeNil)
	err = r.Close(context.Background())
	test.That(t, err, test.ShouldBeNil)
}

func getDataManagerConfig(config *config.Config) (*datamanager.Config, error) {
	svcConfig, ok, err := datamanager.GetServiceConfig(config)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.New("failed to get service config")
	}
	return svcConfig, nil
}

type mockDataSyncServiceServer struct {
	uploadedFiles *[]string
	lock          *sync.Mutex
	v1.UnimplementedDataSyncServiceServer
}

func (m mockDataSyncServiceServer) getUploadedFiles() []string {
	(*m.lock).Lock()
	defer (*m.lock).Unlock()
	return *m.uploadedFiles
}

func (m mockDataSyncServiceServer) Upload(stream v1.DataSyncService_UploadServer) error {
	var fileName string
	for {
		ur, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return err
		}
		if ur.GetMetadata() != nil {
			fileName = ur.GetMetadata().GetFileName()
		}
	}
	(*m.lock).Lock()
	*m.uploadedFiles = append(*m.uploadedFiles, fileName)
	(*m.lock).Unlock()
	return nil
}

//nolint:thelper
func buildAndStartLocalServer(t *testing.T) (rpc.Server, *mockDataSyncServiceServer) {
	logger, _ := golog.NewObservedTestLogger(t)
	rpcServer, err := rpc.NewServer(logger, rpc.WithUnauthenticated())
	test.That(t, err, test.ShouldBeNil)
	mockService := mockDataSyncServiceServer{
		uploadedFiles:                      &[]string{},
		lock:                               &sync.Mutex{},
		UnimplementedDataSyncServiceServer: v1.UnimplementedDataSyncServiceServer{},
	}
	err = rpcServer.RegisterServiceServer(
		context.Background(),
		&v1.DataSyncService_ServiceDesc,
		mockService,
		v1.RegisterDataSyncServiceHandlerFromEndpoint,
	)
	test.That(t, err, test.ShouldBeNil)

	// Stand up the server. Defer stopping the server.
	go func() {
		err := rpcServer.Start()
		test.That(t, err, test.ShouldBeNil)
	}()
	return rpcServer, &mockService
}

func getLocalServerConn(rpcServer rpc.Server, logger golog.Logger) (rpc.ClientConn, error) {
	return rpc.DialDirectGRPC(
		context.Background(),
		rpcServer.InternalAddr().String(),
		logger,
		rpc.WithInsecure(),
	)
}

//nolint:thelper
func getTestSyncerConstructor(t *testing.T, server rpc.Server) datasync.ManagerConstructor {
	return func(logger golog.Logger, cfg *config.Config) (datasync.Manager, error) {
		conn, err := getLocalServerConn(server, logger)
		test.That(t, err, test.ShouldBeNil)
		client := datasync.NewClient(conn)
		return datasync.NewManager(logger, cfg.Cloud.ID, client, conn)
	}
}
