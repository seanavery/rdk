// Package fake implements a fake slam service.
package fake

import (
	"context"

	"github.com/edaniels/golog"
	"go.opencensus.io/trace"

	"go.viam.com/rdk/components/generic"
	"go.viam.com/rdk/config"
	"go.viam.com/rdk/registry"
	"go.viam.com/rdk/resource"
	"go.viam.com/rdk/services/slam"
	"go.viam.com/rdk/spatialmath"
)

var model = resource.NewDefaultModel("fake")

const datasetDirectory = "slam/example_cartographer_outputs/viam-office-02-22-1"

func init() {
	registry.RegisterService(
		slam.Subtype,
		model,
		registry.Service{
			Constructor: func(
				ctx context.Context,
				_ registry.Dependencies,
				config config.Service,
				logger golog.Logger,
			) (interface{}, error) {
				return &SLAM{Name: config.Name, logger: logger, dataCount: -1}, nil
			},
		},
	)
}

var _ = slam.Service(&SLAM{})

// SLAM is a fake slam that returns generic data.
type SLAM struct {
	generic.Echo
	Name      string
	dataCount int
	logger    golog.Logger
}

func (slamSvc *SLAM) getCount() int {
	if slamSvc.dataCount < 0 {
		return 0
	}
	return slamSvc.dataCount
}

// GetPosition returns a Pose and a component reference string of the robot's current location according to SLAM.
func (slamSvc *SLAM) GetPosition(ctx context.Context, name string) (spatialmath.Pose, string, error) {
	ctx, span := trace.StartSpan(ctx, "slam::fake::GetPosition")
	defer span.End()
	return fakeGetPosition(ctx, datasetDirectory, slamSvc)
}

// GetPointCloudMapStream returns a callback function which will return the next chunk of the current pointcloud
// map.
func (slamSvc *SLAM) GetPointCloudMapStream(ctx context.Context, name string) (func() ([]byte, error), error) {
	ctx, span := trace.StartSpan(ctx, "slam::fake::GetPointCloudMapStream")
	defer span.End()
	slamSvc.incrementDataCount()
	return fakeGetPointCloudMapStream(ctx, datasetDirectory, slamSvc)
}

// GetInternalStateStream returns a callback function which will return the next chunk of the current internal
// state of the slam algo.
func (slamSvc *SLAM) GetInternalStateStream(ctx context.Context, name string) (func() ([]byte, error), error) {
	ctx, span := trace.StartSpan(ctx, "slam::fake::GetInternalStateStream")
	defer span.End()
	return fakeGetInternalStateStream(ctx, datasetDirectory, slamSvc)
}

// incrementDataCount is not thread safe but that is ok as we only intend a single user to be interacting
// with it at a time.
func (slamSvc *SLAM) incrementDataCount() {
	slamSvc.dataCount = ((slamSvc.dataCount + 1) % maxDataCount)
}
