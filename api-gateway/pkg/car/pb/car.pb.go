// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.1
// source: car.proto

package pb

import (
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

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand        string  `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Model        string  `protobuf:"bytes,3,opt,name=model,proto3" json:"model,omitempty"`
	Year         int32   `protobuf:"varint,4,opt,name=year,proto3" json:"year,omitempty"`
	Displacement float32 `protobuf:"fixed32,6,opt,name=displacement,proto3" json:"displacement,omitempty"`
	Body         string  `protobuf:"bytes,7,opt,name=body,proto3" json:"body,omitempty"`
	Mpg          float32 `protobuf:"fixed32,8,opt,name=mpg,proto3" json:"mpg,omitempty"`
	Cylinders    int32   `protobuf:"varint,9,opt,name=cylinders,proto3" json:"cylinders,omitempty"`
	Horsepower   float32 `protobuf:"fixed32,10,opt,name=horsepower,proto3" json:"horsepower,omitempty"`
	BrakeSystem  string  `protobuf:"bytes,11,opt,name=brake_system,json=brakeSystem,proto3" json:"brake_system,omitempty"`
	Aspiration   string  `protobuf:"bytes,12,opt,name=aspiration,proto3" json:"aspiration,omitempty"`
	Acceleration float32 `protobuf:"fixed32,13,opt,name=acceleration,proto3" json:"acceleration,omitempty"`
	Country      string  `protobuf:"bytes,14,opt,name=country,proto3" json:"country,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

func (x *Car) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Car) GetBrand() string {
	if x != nil {
		return x.Brand
	}
	return ""
}

func (x *Car) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Car) GetYear() int32 {
	if x != nil {
		return x.Year
	}
	return 0
}

func (x *Car) GetDisplacement() float32 {
	if x != nil {
		return x.Displacement
	}
	return 0
}

func (x *Car) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Car) GetMpg() float32 {
	if x != nil {
		return x.Mpg
	}
	return 0
}

func (x *Car) GetCylinders() int32 {
	if x != nil {
		return x.Cylinders
	}
	return 0
}

func (x *Car) GetHorsepower() float32 {
	if x != nil {
		return x.Horsepower
	}
	return 0
}

func (x *Car) GetBrakeSystem() string {
	if x != nil {
		return x.BrakeSystem
	}
	return ""
}

func (x *Car) GetAspiration() string {
	if x != nil {
		return x.Aspiration
	}
	return ""
}

func (x *Car) GetAcceleration() float32 {
	if x != nil {
		return x.Acceleration
	}
	return 0
}

func (x *Car) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

type CreateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CreateCarRequest) Reset() {
	*x = CreateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarRequest) ProtoMessage() {}

func (x *CreateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarRequest.ProtoReflect.Descriptor instead.
func (*CreateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{1}
}

func (x *CreateCarRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type CreateCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CreateCarResponse) Reset() {
	*x = CreateCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarResponse) ProtoMessage() {}

func (x *CreateCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarResponse.ProtoReflect.Descriptor instead.
func (*CreateCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{2}
}

func (x *CreateCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type ReadCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *ReadCarRequest) Reset() {
	*x = ReadCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCarRequest) ProtoMessage() {}

func (x *ReadCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCarRequest.ProtoReflect.Descriptor instead.
func (*ReadCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{3}
}

func (x *ReadCarRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type ReadCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *ReadCarResponse) Reset() {
	*x = ReadCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadCarResponse) ProtoMessage() {}

func (x *ReadCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadCarResponse.ProtoReflect.Descriptor instead.
func (*ReadCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{4}
}

func (x *ReadCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type UpdateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *UpdateCarRequest) Reset() {
	*x = UpdateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarRequest) ProtoMessage() {}

func (x *UpdateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarRequest.ProtoReflect.Descriptor instead.
func (*UpdateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateCarRequest) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type UpdateCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *UpdateCarResponse) Reset() {
	*x = UpdateCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarResponse) ProtoMessage() {}

func (x *UpdateCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarResponse.ProtoReflect.Descriptor instead.
func (*UpdateCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type DeleteCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *DeleteCarRequest) Reset() {
	*x = DeleteCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarRequest) ProtoMessage() {}

func (x *DeleteCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarRequest.ProtoReflect.Descriptor instead.
func (*DeleteCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteCarRequest) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type DeleteCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
}

func (x *DeleteCarResponse) Reset() {
	*x = DeleteCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarResponse) ProtoMessage() {}

func (x *DeleteCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarResponse.ProtoReflect.Descriptor instead.
func (*DeleteCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCarResponse) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

type ListCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListCarRequest) Reset() {
	*x = ListCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarRequest) ProtoMessage() {}

func (x *ListCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarRequest.ProtoReflect.Descriptor instead.
func (*ListCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{9}
}

type ListCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Car *Car `protobuf:"bytes,1,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *ListCarResponse) Reset() {
	*x = ListCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListCarResponse) ProtoMessage() {}

func (x *ListCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListCarResponse.ProtoReflect.Descriptor instead.
func (*ListCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{10}
}

func (x *ListCarResponse) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

var File_car_proto protoreflect.FileDescriptor

var file_car_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x63, 0x61, 0x72,
	0x22, 0xde, 0x02, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x22, 0x0a, 0x0c, 0x64, 0x69, 0x73, 0x70,
	0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x70, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6d,
	0x70, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x79, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x79, 0x6c, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x73,
	0x12, 0x1e, 0x0a, 0x0a, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x70, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x0a, 0x68, 0x6f, 0x72, 0x73, 0x65, 0x70, 0x6f, 0x77, 0x65, 0x72,
	0x12, 0x21, 0x0a, 0x0c, 0x62, 0x72, 0x61, 0x6b, 0x65, 0x5f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x62, 0x72, 0x61, 0x6b, 0x65, 0x53, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x73, 0x70, 0x69, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x73, 0x70, 0x69, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x02, 0x52, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x6c,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72,
	0x79, 0x22, 0x2e, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61,
	0x72, 0x22, 0x2f, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63,
	0x61, 0x72, 0x22, 0x27, 0x0a, 0x0e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x0f, 0x52,
	0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a,
	0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61,
	0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x2e, 0x0a, 0x10, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a,
	0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61,
	0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x2f, 0x0a, 0x11, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1a, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63,
	0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22, 0x29, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x63,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x49, 0x64, 0x22, 0x10, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03,
	0x63, 0x61, 0x72, 0x32, 0xae, 0x02, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12,
	0x15, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34,
	0x0a, 0x07, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x72, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14,
	0x2e, 0x63, 0x61, 0x72, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x12, 0x15, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x15, 0x2e,
	0x63, 0x61, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x07,
	0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x12, 0x13, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63,
	0x61, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x30, 0x01, 0x42, 0x0e, 0x5a, 0x0c, 0x2e, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x61,
	0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_car_proto_rawDescOnce sync.Once
	file_car_proto_rawDescData = file_car_proto_rawDesc
)

func file_car_proto_rawDescGZIP() []byte {
	file_car_proto_rawDescOnce.Do(func() {
		file_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_proto_rawDescData)
	})
	return file_car_proto_rawDescData
}

var file_car_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_car_proto_goTypes = []interface{}{
	(*Car)(nil),               // 0: car.Car
	(*CreateCarRequest)(nil),  // 1: car.CreateCarRequest
	(*CreateCarResponse)(nil), // 2: car.CreateCarResponse
	(*ReadCarRequest)(nil),    // 3: car.ReadCarRequest
	(*ReadCarResponse)(nil),   // 4: car.ReadCarResponse
	(*UpdateCarRequest)(nil),  // 5: car.UpdateCarRequest
	(*UpdateCarResponse)(nil), // 6: car.UpdateCarResponse
	(*DeleteCarRequest)(nil),  // 7: car.DeleteCarRequest
	(*DeleteCarResponse)(nil), // 8: car.DeleteCarResponse
	(*ListCarRequest)(nil),    // 9: car.ListCarRequest
	(*ListCarResponse)(nil),   // 10: car.ListCarResponse
}
var file_car_proto_depIdxs = []int32{
	0,  // 0: car.CreateCarRequest.car:type_name -> car.Car
	0,  // 1: car.CreateCarResponse.car:type_name -> car.Car
	0,  // 2: car.ReadCarResponse.car:type_name -> car.Car
	0,  // 3: car.UpdateCarRequest.car:type_name -> car.Car
	0,  // 4: car.UpdateCarResponse.car:type_name -> car.Car
	0,  // 5: car.ListCarResponse.car:type_name -> car.Car
	1,  // 6: car.CarService.CreateCar:input_type -> car.CreateCarRequest
	3,  // 7: car.CarService.ReadCar:input_type -> car.ReadCarRequest
	5,  // 8: car.CarService.UpdateCar:input_type -> car.UpdateCarRequest
	7,  // 9: car.CarService.DeleteCar:input_type -> car.DeleteCarRequest
	9,  // 10: car.CarService.ListCar:input_type -> car.ListCarRequest
	2,  // 11: car.CarService.CreateCar:output_type -> car.CreateCarResponse
	4,  // 12: car.CarService.ReadCar:output_type -> car.ReadCarResponse
	6,  // 13: car.CarService.UpdateCar:output_type -> car.UpdateCarResponse
	8,  // 14: car.CarService.DeleteCar:output_type -> car.DeleteCarResponse
	10, // 15: car.CarService.ListCar:output_type -> car.ListCarResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_car_proto_init() }
func file_car_proto_init() {
	if File_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarRequest); i {
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
		file_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarResponse); i {
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
		file_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCarRequest); i {
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
		file_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadCarResponse); i {
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
		file_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarRequest); i {
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
		file_car_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarResponse); i {
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
		file_car_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarRequest); i {
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
		file_car_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarResponse); i {
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
		file_car_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarRequest); i {
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
		file_car_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListCarResponse); i {
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
			RawDescriptor: file_car_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_proto_goTypes,
		DependencyIndexes: file_car_proto_depIdxs,
		MessageInfos:      file_car_proto_msgTypes,
	}.Build()
	File_car_proto = out.File
	file_car_proto_rawDesc = nil
	file_car_proto_goTypes = nil
	file_car_proto_depIdxs = nil
}
