// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/api/component/v1/imu.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// AngularVelocity contains angular velocity in deg/s across x/y/z axes.
type AngularVelocity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Velocity in deg/s across the x-axis
	XDegsPerSec float64 `protobuf:"fixed64,1,opt,name=x_degs_per_sec,json=xDegsPerSec,proto3" json:"x_degs_per_sec,omitempty"`
	// Velocity in deg/s across the y-axis
	YDegsPerSec float64 `protobuf:"fixed64,2,opt,name=y_degs_per_sec,json=yDegsPerSec,proto3" json:"y_degs_per_sec,omitempty"`
	// Velocity in deg/s across the z-axis
	ZDegsPerSec float64 `protobuf:"fixed64,3,opt,name=z_degs_per_sec,json=zDegsPerSec,proto3" json:"z_degs_per_sec,omitempty"`
}

func (x *AngularVelocity) Reset() {
	*x = AngularVelocity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AngularVelocity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AngularVelocity) ProtoMessage() {}

func (x *AngularVelocity) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AngularVelocity.ProtoReflect.Descriptor instead.
func (*AngularVelocity) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{0}
}

func (x *AngularVelocity) GetXDegsPerSec() float64 {
	if x != nil {
		return x.XDegsPerSec
	}
	return 0
}

func (x *AngularVelocity) GetYDegsPerSec() float64 {
	if x != nil {
		return x.YDegsPerSec
	}
	return 0
}

func (x *AngularVelocity) GetZDegsPerSec() float64 {
	if x != nil {
		return x.ZDegsPerSec
	}
	return 0
}

// EulerAngles are three angles used to represent the rotation of an object in 3D Euclidean space
// The Tait–Bryan angle formalism is used, with rotations around three distinct axes in the z-y′-x″ sequence.
type EulerAngles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Rotation in deg around the x-axis
	RollDeg float64 `protobuf:"fixed64,1,opt,name=roll_deg,json=rollDeg,proto3" json:"roll_deg,omitempty"`
	// Rotation in deg around the y-axis
	PitchDeg float64 `protobuf:"fixed64,2,opt,name=pitch_deg,json=pitchDeg,proto3" json:"pitch_deg,omitempty"`
	// Rotation in deg around the z-axis
	YawDeg float64 `protobuf:"fixed64,3,opt,name=yaw_deg,json=yawDeg,proto3" json:"yaw_deg,omitempty"`
}

func (x *EulerAngles) Reset() {
	*x = EulerAngles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EulerAngles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EulerAngles) ProtoMessage() {}

func (x *EulerAngles) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EulerAngles.ProtoReflect.Descriptor instead.
func (*EulerAngles) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{1}
}

func (x *EulerAngles) GetRollDeg() float64 {
	if x != nil {
		return x.RollDeg
	}
	return 0
}

func (x *EulerAngles) GetPitchDeg() float64 {
	if x != nil {
		return x.PitchDeg
	}
	return 0
}

func (x *EulerAngles) GetYawDeg() float64 {
	if x != nil {
		return x.YawDeg
	}
	return 0
}

// Acceleration contains linear acceleration in mm/s^2 across x/y/z axes.
type Acceleration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Acceleration in mm/s^2 across the x-axis
	XMmPerSecPerSec float64 `protobuf:"fixed64,1,opt,name=x_mm_per_sec_per_sec,json=xMmPerSecPerSec,proto3" json:"x_mm_per_sec_per_sec,omitempty"`
	// Acceleration in mm/s^2 across the y-axis
	YMmPerSecPerSec float64 `protobuf:"fixed64,2,opt,name=y_mm_per_sec_per_sec,json=yMmPerSecPerSec,proto3" json:"y_mm_per_sec_per_sec,omitempty"`
	// Acceleration in mm/s^2 across the z-axis
	ZMmPerSecPerSec float64 `protobuf:"fixed64,3,opt,name=z_mm_per_sec_per_sec,json=zMmPerSecPerSec,proto3" json:"z_mm_per_sec_per_sec,omitempty"`
}

func (x *Acceleration) Reset() {
	*x = Acceleration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Acceleration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Acceleration) ProtoMessage() {}

func (x *Acceleration) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Acceleration.ProtoReflect.Descriptor instead.
func (*Acceleration) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{2}
}

func (x *Acceleration) GetXMmPerSecPerSec() float64 {
	if x != nil {
		return x.XMmPerSecPerSec
	}
	return 0
}

func (x *Acceleration) GetYMmPerSecPerSec() float64 {
	if x != nil {
		return x.YMmPerSecPerSec
	}
	return 0
}

func (x *Acceleration) GetZMmPerSecPerSec() float64 {
	if x != nil {
		return x.ZMmPerSecPerSec
	}
	return 0
}

type IMUServiceReadAngularVelocityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an IMU
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IMUServiceReadAngularVelocityRequest) Reset() {
	*x = IMUServiceReadAngularVelocityRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceReadAngularVelocityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceReadAngularVelocityRequest) ProtoMessage() {}

func (x *IMUServiceReadAngularVelocityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceReadAngularVelocityRequest.ProtoReflect.Descriptor instead.
func (*IMUServiceReadAngularVelocityRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{3}
}

func (x *IMUServiceReadAngularVelocityRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IMUServiceReadAngularVelocityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// AngularVelocity contains angular velocity in deg/s across x/y/z axes.
	AngularVelocity *AngularVelocity `protobuf:"bytes,1,opt,name=angular_velocity,json=angularVelocity,proto3" json:"angular_velocity,omitempty"`
}

func (x *IMUServiceReadAngularVelocityResponse) Reset() {
	*x = IMUServiceReadAngularVelocityResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceReadAngularVelocityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceReadAngularVelocityResponse) ProtoMessage() {}

func (x *IMUServiceReadAngularVelocityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceReadAngularVelocityResponse.ProtoReflect.Descriptor instead.
func (*IMUServiceReadAngularVelocityResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{4}
}

func (x *IMUServiceReadAngularVelocityResponse) GetAngularVelocity() *AngularVelocity {
	if x != nil {
		return x.AngularVelocity
	}
	return nil
}

type IMUServiceReadOrientationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an IMU
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IMUServiceReadOrientationRequest) Reset() {
	*x = IMUServiceReadOrientationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceReadOrientationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceReadOrientationRequest) ProtoMessage() {}

func (x *IMUServiceReadOrientationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceReadOrientationRequest.ProtoReflect.Descriptor instead.
func (*IMUServiceReadOrientationRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{5}
}

func (x *IMUServiceReadOrientationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IMUServiceReadOrientationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// EulerAngles are three angles used to represent the rotation of an object in 3D Euclidean space
	// The Tait–Bryan angle formalism is used, with rotations around three distinct axes in the z-y′-x″ sequence.
	Orientation *EulerAngles `protobuf:"bytes,1,opt,name=orientation,proto3" json:"orientation,omitempty"`
}

func (x *IMUServiceReadOrientationResponse) Reset() {
	*x = IMUServiceReadOrientationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceReadOrientationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceReadOrientationResponse) ProtoMessage() {}

func (x *IMUServiceReadOrientationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceReadOrientationResponse.ProtoReflect.Descriptor instead.
func (*IMUServiceReadOrientationResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{6}
}

func (x *IMUServiceReadOrientationResponse) GetOrientation() *EulerAngles {
	if x != nil {
		return x.Orientation
	}
	return nil
}

type IMUServiceReadAccelerationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of an IMU
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *IMUServiceReadAccelerationRequest) Reset() {
	*x = IMUServiceReadAccelerationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceReadAccelerationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceReadAccelerationRequest) ProtoMessage() {}

func (x *IMUServiceReadAccelerationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceReadAccelerationRequest.ProtoReflect.Descriptor instead.
func (*IMUServiceReadAccelerationRequest) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{7}
}

func (x *IMUServiceReadAccelerationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type IMUServiceReadAccelerationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Acceleration contains acceleration in mm/s^2 across x/y/z axes.
	Acceleration *Acceleration `protobuf:"bytes,1,opt,name=acceleration,proto3" json:"acceleration,omitempty"`
}

func (x *IMUServiceReadAccelerationResponse) Reset() {
	*x = IMUServiceReadAccelerationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_api_component_v1_imu_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IMUServiceReadAccelerationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IMUServiceReadAccelerationResponse) ProtoMessage() {}

func (x *IMUServiceReadAccelerationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_api_component_v1_imu_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IMUServiceReadAccelerationResponse.ProtoReflect.Descriptor instead.
func (*IMUServiceReadAccelerationResponse) Descriptor() ([]byte, []int) {
	return file_proto_api_component_v1_imu_proto_rawDescGZIP(), []int{8}
}

func (x *IMUServiceReadAccelerationResponse) GetAcceleration() *Acceleration {
	if x != nil {
		return x.Acceleration
	}
	return nil
}

var File_proto_api_component_v1_imu_proto protoreflect.FileDescriptor

var file_proto_api_component_v1_imu_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x69, 0x6d, 0x75, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x80, 0x01, 0x0a, 0x0f, 0x41, 0x6e, 0x67,
	0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a, 0x0e,
	0x78, 0x5f, 0x64, 0x65, 0x67, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x78, 0x44, 0x65, 0x67, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x12, 0x23, 0x0a, 0x0e, 0x79, 0x5f, 0x64, 0x65, 0x67, 0x73, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x79, 0x44, 0x65, 0x67, 0x73,
	0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x12, 0x23, 0x0a, 0x0e, 0x7a, 0x5f, 0x64, 0x65, 0x67, 0x73,
	0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b,
	0x7a, 0x44, 0x65, 0x67, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x22, 0x5e, 0x0a, 0x0b, 0x45,
	0x75, 0x6c, 0x65, 0x72, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x72, 0x6f,
	0x6c, 0x6c, 0x5f, 0x64, 0x65, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x6f,
	0x6c, 0x6c, 0x44, 0x65, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x69, 0x74, 0x63, 0x68, 0x5f, 0x64,
	0x65, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x70, 0x69, 0x74, 0x63, 0x68, 0x44,
	0x65, 0x67, 0x12, 0x17, 0x0a, 0x07, 0x79, 0x61, 0x77, 0x5f, 0x64, 0x65, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x79, 0x61, 0x77, 0x44, 0x65, 0x67, 0x22, 0x9b, 0x01, 0x0a, 0x0c,
	0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x14,
	0x78, 0x5f, 0x6d, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x5f, 0x70, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x78, 0x4d, 0x6d, 0x50,
	0x65, 0x72, 0x53, 0x65, 0x63, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x12, 0x2d, 0x0a, 0x14, 0x79,
	0x5f, 0x6d, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x5f,
	0x73, 0x65, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x79, 0x4d, 0x6d, 0x50, 0x65,
	0x72, 0x53, 0x65, 0x63, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x12, 0x2d, 0x0a, 0x14, 0x7a, 0x5f,
	0x6d, 0x6d, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x63, 0x5f, 0x70, 0x65, 0x72, 0x5f, 0x73,
	0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0f, 0x7a, 0x4d, 0x6d, 0x50, 0x65, 0x72,
	0x53, 0x65, 0x63, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x22, 0x3a, 0x0a, 0x24, 0x49, 0x4d, 0x55,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6e, 0x67, 0x75, 0x6c,
	0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x7b, 0x0a, 0x25, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52,
	0x0a, 0x10, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69,
	0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74,
	0x79, 0x52, 0x0f, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69,
	0x74, 0x79, 0x22, 0x36, 0x0a, 0x20, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x52, 0x65, 0x61, 0x64, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x6a, 0x0a, 0x21, 0x49, 0x4d,
	0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x72, 0x69, 0x65,
	0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x45, 0x0a, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x75,
	0x6c, 0x65, 0x72, 0x41, 0x6e, 0x67, 0x6c, 0x65, 0x73, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x65, 0x6e,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x37, 0x0a, 0x21, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x6e, 0x0a, 0x22, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61,
	0x64, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32,
	0xd2, 0x04, 0x0a, 0x0a, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xc9,
	0x01, 0x0a, 0x13, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6e, 0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65,
	0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0x3c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6e,
	0x67, 0x75, 0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4d,
	0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x41, 0x6e, 0x67, 0x75,
	0x6c, 0x61, 0x72, 0x56, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x35, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2f, 0x12, 0x2d, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x69,
	0x6d, 0x75, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x6e, 0x67, 0x75, 0x6c, 0x61,
	0x72, 0x5f, 0x76, 0x65, 0x6c, 0x6f, 0x63, 0x69, 0x74, 0x79, 0x12, 0xb8, 0x01, 0x0a, 0x0f, 0x52,
	0x65, 0x61, 0x64, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x38,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f,
	0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x39, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64,
	0x4f, 0x72, 0x69, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x30, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2a, 0x12, 0x28, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x69,
	0x6d, 0x75, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x6f, 0x72, 0x69, 0x65, 0x6e, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0xbc, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x41, 0x63,
	0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2e, 0x76, 0x31, 0x2e, 0x49, 0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x61, 0x64, 0x41, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x3a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x49,
	0x4d, 0x55, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x61, 0x64, 0x41, 0x63, 0x63,
	0x65, 0x6c, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x31, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2f, 0x69, 0x6d, 0x75,
	0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x42, 0x4d, 0x0a, 0x23, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x69, 0x61, 0x6d,
	0x2e, 0x72, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x5a, 0x26, 0x67, 0x6f, 0x2e,
	0x76, 0x69, 0x61, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x64, 0x6b, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74,
	0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_api_component_v1_imu_proto_rawDescOnce sync.Once
	file_proto_api_component_v1_imu_proto_rawDescData = file_proto_api_component_v1_imu_proto_rawDesc
)

func file_proto_api_component_v1_imu_proto_rawDescGZIP() []byte {
	file_proto_api_component_v1_imu_proto_rawDescOnce.Do(func() {
		file_proto_api_component_v1_imu_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_api_component_v1_imu_proto_rawDescData)
	})
	return file_proto_api_component_v1_imu_proto_rawDescData
}

var file_proto_api_component_v1_imu_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_api_component_v1_imu_proto_goTypes = []interface{}{
	(*AngularVelocity)(nil),                       // 0: proto.api.component.v1.AngularVelocity
	(*EulerAngles)(nil),                           // 1: proto.api.component.v1.EulerAngles
	(*Acceleration)(nil),                          // 2: proto.api.component.v1.Acceleration
	(*IMUServiceReadAngularVelocityRequest)(nil),  // 3: proto.api.component.v1.IMUServiceReadAngularVelocityRequest
	(*IMUServiceReadAngularVelocityResponse)(nil), // 4: proto.api.component.v1.IMUServiceReadAngularVelocityResponse
	(*IMUServiceReadOrientationRequest)(nil),      // 5: proto.api.component.v1.IMUServiceReadOrientationRequest
	(*IMUServiceReadOrientationResponse)(nil),     // 6: proto.api.component.v1.IMUServiceReadOrientationResponse
	(*IMUServiceReadAccelerationRequest)(nil),     // 7: proto.api.component.v1.IMUServiceReadAccelerationRequest
	(*IMUServiceReadAccelerationResponse)(nil),    // 8: proto.api.component.v1.IMUServiceReadAccelerationResponse
}
var file_proto_api_component_v1_imu_proto_depIdxs = []int32{
	0, // 0: proto.api.component.v1.IMUServiceReadAngularVelocityResponse.angular_velocity:type_name -> proto.api.component.v1.AngularVelocity
	1, // 1: proto.api.component.v1.IMUServiceReadOrientationResponse.orientation:type_name -> proto.api.component.v1.EulerAngles
	2, // 2: proto.api.component.v1.IMUServiceReadAccelerationResponse.acceleration:type_name -> proto.api.component.v1.Acceleration
	3, // 3: proto.api.component.v1.IMUService.ReadAngularVelocity:input_type -> proto.api.component.v1.IMUServiceReadAngularVelocityRequest
	5, // 4: proto.api.component.v1.IMUService.ReadOrientation:input_type -> proto.api.component.v1.IMUServiceReadOrientationRequest
	7, // 5: proto.api.component.v1.IMUService.ReadAcceleration:input_type -> proto.api.component.v1.IMUServiceReadAccelerationRequest
	4, // 6: proto.api.component.v1.IMUService.ReadAngularVelocity:output_type -> proto.api.component.v1.IMUServiceReadAngularVelocityResponse
	6, // 7: proto.api.component.v1.IMUService.ReadOrientation:output_type -> proto.api.component.v1.IMUServiceReadOrientationResponse
	8, // 8: proto.api.component.v1.IMUService.ReadAcceleration:output_type -> proto.api.component.v1.IMUServiceReadAccelerationResponse
	6, // [6:9] is the sub-list for method output_type
	3, // [3:6] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_api_component_v1_imu_proto_init() }
func file_proto_api_component_v1_imu_proto_init() {
	if File_proto_api_component_v1_imu_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_api_component_v1_imu_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AngularVelocity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_component_v1_imu_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EulerAngles); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_component_v1_imu_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Acceleration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_component_v1_imu_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceReadAngularVelocityRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_component_v1_imu_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceReadAngularVelocityResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_component_v1_imu_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceReadOrientationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_component_v1_imu_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceReadOrientationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_component_v1_imu_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceReadAccelerationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_api_component_v1_imu_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IMUServiceReadAccelerationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_api_component_v1_imu_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_api_component_v1_imu_proto_goTypes,
		DependencyIndexes: file_proto_api_component_v1_imu_proto_depIdxs,
		MessageInfos:      file_proto_api_component_v1_imu_proto_msgTypes,
	}.Build()
	File_proto_api_component_v1_imu_proto = out.File
	file_proto_api_component_v1_imu_proto_rawDesc = nil
	file_proto_api_component_v1_imu_proto_goTypes = nil
	file_proto_api_component_v1_imu_proto_depIdxs = nil
}
