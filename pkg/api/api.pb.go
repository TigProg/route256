// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: api.proto

package api

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

type BusBooking struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Route string `protobuf:"bytes,2,opt,name=route,proto3" json:"route,omitempty"`
	Date  string `protobuf:"bytes,3,opt,name=date,proto3" json:"date,omitempty"`
	Seat  uint32 `protobuf:"varint,4,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *BusBooking) Reset() {
	*x = BusBooking{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBooking) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBooking) ProtoMessage() {}

func (x *BusBooking) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBooking.ProtoReflect.Descriptor instead.
func (*BusBooking) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *BusBooking) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BusBooking) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *BusBooking) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BusBooking) GetSeat() uint32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

type BusBookingListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BusBookingListRequest) Reset() {
	*x = BusBookingListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingListRequest) ProtoMessage() {}

func (x *BusBookingListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingListRequest.ProtoReflect.Descriptor instead.
func (*BusBookingListRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

type BusBookingListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusBookings []*BusBooking `protobuf:"bytes,1,rep,name=bus_bookings,json=busBookings,proto3" json:"bus_bookings,omitempty"`
}

func (x *BusBookingListResponse) Reset() {
	*x = BusBookingListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingListResponse) ProtoMessage() {}

func (x *BusBookingListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingListResponse.ProtoReflect.Descriptor instead.
func (*BusBookingListResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *BusBookingListResponse) GetBusBookings() []*BusBooking {
	if x != nil {
		return x.BusBookings
	}
	return nil
}

type BusBookingAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Route string `protobuf:"bytes,1,opt,name=route,proto3" json:"route,omitempty"`
	Date  string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Seat  uint32 `protobuf:"varint,3,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *BusBookingAddRequest) Reset() {
	*x = BusBookingAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingAddRequest) ProtoMessage() {}

func (x *BusBookingAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingAddRequest.ProtoReflect.Descriptor instead.
func (*BusBookingAddRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *BusBookingAddRequest) GetRoute() string {
	if x != nil {
		return x.Route
	}
	return ""
}

func (x *BusBookingAddRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BusBookingAddRequest) GetSeat() uint32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

type BusBookingAddResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BusBookingAddResponse) Reset() {
	*x = BusBookingAddResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingAddResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingAddResponse) ProtoMessage() {}

func (x *BusBookingAddResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingAddResponse.ProtoReflect.Descriptor instead.
func (*BusBookingAddResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *BusBookingAddResponse) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BusBookingGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BusBookingGetRequest) Reset() {
	*x = BusBookingGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingGetRequest) ProtoMessage() {}

func (x *BusBookingGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingGetRequest.ProtoReflect.Descriptor instead.
func (*BusBookingGetRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *BusBookingGetRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BusBookingGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusBooking *BusBooking `protobuf:"bytes,1,opt,name=bus_booking,json=busBooking,proto3" json:"bus_booking,omitempty"`
}

func (x *BusBookingGetResponse) Reset() {
	*x = BusBookingGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingGetResponse) ProtoMessage() {}

func (x *BusBookingGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingGetResponse.ProtoReflect.Descriptor instead.
func (*BusBookingGetResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *BusBookingGetResponse) GetBusBooking() *BusBooking {
	if x != nil {
		return x.BusBooking
	}
	return nil
}

type BusBookingChangeSeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Seat uint32 `protobuf:"varint,2,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *BusBookingChangeSeatRequest) Reset() {
	*x = BusBookingChangeSeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingChangeSeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingChangeSeatRequest) ProtoMessage() {}

func (x *BusBookingChangeSeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingChangeSeatRequest.ProtoReflect.Descriptor instead.
func (*BusBookingChangeSeatRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *BusBookingChangeSeatRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BusBookingChangeSeatRequest) GetSeat() uint32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

type BusBookingChangeSeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BusBookingChangeSeatResponse) Reset() {
	*x = BusBookingChangeSeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingChangeSeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingChangeSeatResponse) ProtoMessage() {}

func (x *BusBookingChangeSeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingChangeSeatResponse.ProtoReflect.Descriptor instead.
func (*BusBookingChangeSeatResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{8}
}

type BusBookingChangeDateSeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Date string `protobuf:"bytes,2,opt,name=date,proto3" json:"date,omitempty"`
	Seat uint32 `protobuf:"varint,3,opt,name=seat,proto3" json:"seat,omitempty"`
}

func (x *BusBookingChangeDateSeatRequest) Reset() {
	*x = BusBookingChangeDateSeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingChangeDateSeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingChangeDateSeatRequest) ProtoMessage() {}

func (x *BusBookingChangeDateSeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingChangeDateSeatRequest.ProtoReflect.Descriptor instead.
func (*BusBookingChangeDateSeatRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{9}
}

func (x *BusBookingChangeDateSeatRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *BusBookingChangeDateSeatRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *BusBookingChangeDateSeatRequest) GetSeat() uint32 {
	if x != nil {
		return x.Seat
	}
	return 0
}

type BusBookingChangeDateSeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BusBookingChangeDateSeatResponse) Reset() {
	*x = BusBookingChangeDateSeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingChangeDateSeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingChangeDateSeatResponse) ProtoMessage() {}

func (x *BusBookingChangeDateSeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingChangeDateSeatResponse.ProtoReflect.Descriptor instead.
func (*BusBookingChangeDateSeatResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{10}
}

type BusBookingDeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *BusBookingDeleteRequest) Reset() {
	*x = BusBookingDeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingDeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingDeleteRequest) ProtoMessage() {}

func (x *BusBookingDeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingDeleteRequest.ProtoReflect.Descriptor instead.
func (*BusBookingDeleteRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{11}
}

func (x *BusBookingDeleteRequest) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BusBookingDeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BusBookingDeleteResponse) Reset() {
	*x = BusBookingDeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusBookingDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusBookingDeleteResponse) ProtoMessage() {}

func (x *BusBookingDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusBookingDeleteResponse.ProtoReflect.Descriptor instead.
func (*BusBookingDeleteResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{12}
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x6f, 0x7a, 0x6f,
	0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5a, 0x0a, 0x0a, 0x42,
	0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12,
	0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x17, 0x0a, 0x15, 0x42, 0x75, 0x73, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x59, 0x0a, 0x16, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3f, 0x0a, 0x0c, 0x62, 0x75,
	0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x0b,
	0x62, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x22, 0x54, 0x0a, 0x14, 0x42,
	0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x65, 0x61,
	0x74, 0x22, 0x27, 0x0a, 0x15, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x26, 0x0a, 0x14, 0x42, 0x75,
	0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x56, 0x0a, 0x15, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0b, 0x62,
	0x75, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x52, 0x0a,
	0x62, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x22, 0x41, 0x0a, 0x1b, 0x42, 0x75,
	0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x1e, 0x0a,
	0x1c, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x59, 0x0a,
	0x1f, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x44, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x65, 0x61, 0x74, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x04, 0x73, 0x65, 0x61, 0x74, 0x22, 0x22, 0x0a, 0x20, 0x42, 0x75, 0x73, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x29, 0x0a, 0x17,
	0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1a, 0x0a, 0x18, 0x42, 0x75, 0x73, 0x42, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x32, 0xc5, 0x06, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x7d, 0x0a,
	0x0e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x27, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x75, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x7c, 0x0a, 0x0d,
	0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x12, 0x26, 0x2e,
	0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76,
	0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b,
	0x69, 0x6e, 0x67, 0x41, 0x64, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x5f,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x12, 0x7e, 0x0a, 0x0d, 0x42, 0x75,
	0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x12, 0x26, 0x2e, 0x6f, 0x7a,
	0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42,
	0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d,
	0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x5f, 0x62, 0x6f,
	0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x96, 0x01, 0x0a, 0x14, 0x42,
	0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53,
	0x65, 0x61, 0x74, 0x12, 0x2d, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d,
	0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63,
	0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x1a, 0x14, 0x2f, 0x76, 0x31, 0x2f,
	0x62, 0x75, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x73, 0x65, 0x61, 0x74,
	0x3a, 0x01, 0x2a, 0x12, 0x9d, 0x01, 0x0a, 0x18, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74,
	0x12, 0x31, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x43, 0x68,
	0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x32, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d,
	0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65, 0x53, 0x65, 0x61, 0x74, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x1a,
	0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67,
	0x3a, 0x01, 0x2a, 0x12, 0x85, 0x01, 0x0a, 0x10, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x29, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e,
	0x64, 0x65, 0x76, 0x2e, 0x6d, 0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2a, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2e, 0x6d,
	0x63, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x42, 0x75, 0x73, 0x42, 0x6f, 0x6f, 0x6b, 0x69, 0x6e,
	0x67, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x2a, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x73,
	0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x67, 0x3a, 0x01, 0x2a, 0x42, 0x31, 0x5a, 0x2f, 0x67,
	0x69, 0x74, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x7a, 0x6f, 0x6e, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x74,
	0x69, 0x67, 0x70, 0x72, 0x6f, 0x67, 0x2f, 0x62, 0x75, 0x73, 0x5f, 0x62, 0x6f, 0x6f, 0x6b, 0x69,
	0x6e, 0x67, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x3b, 0x61, 0x70, 0x69, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_proto_goTypes = []interface{}{
	(*BusBooking)(nil),                       // 0: ozon.dev.mc2.api.BusBooking
	(*BusBookingListRequest)(nil),            // 1: ozon.dev.mc2.api.BusBookingListRequest
	(*BusBookingListResponse)(nil),           // 2: ozon.dev.mc2.api.BusBookingListResponse
	(*BusBookingAddRequest)(nil),             // 3: ozon.dev.mc2.api.BusBookingAddRequest
	(*BusBookingAddResponse)(nil),            // 4: ozon.dev.mc2.api.BusBookingAddResponse
	(*BusBookingGetRequest)(nil),             // 5: ozon.dev.mc2.api.BusBookingGetRequest
	(*BusBookingGetResponse)(nil),            // 6: ozon.dev.mc2.api.BusBookingGetResponse
	(*BusBookingChangeSeatRequest)(nil),      // 7: ozon.dev.mc2.api.BusBookingChangeSeatRequest
	(*BusBookingChangeSeatResponse)(nil),     // 8: ozon.dev.mc2.api.BusBookingChangeSeatResponse
	(*BusBookingChangeDateSeatRequest)(nil),  // 9: ozon.dev.mc2.api.BusBookingChangeDateSeatRequest
	(*BusBookingChangeDateSeatResponse)(nil), // 10: ozon.dev.mc2.api.BusBookingChangeDateSeatResponse
	(*BusBookingDeleteRequest)(nil),          // 11: ozon.dev.mc2.api.BusBookingDeleteRequest
	(*BusBookingDeleteResponse)(nil),         // 12: ozon.dev.mc2.api.BusBookingDeleteResponse
}
var file_api_proto_depIdxs = []int32{
	0,  // 0: ozon.dev.mc2.api.BusBookingListResponse.bus_bookings:type_name -> ozon.dev.mc2.api.BusBooking
	0,  // 1: ozon.dev.mc2.api.BusBookingGetResponse.bus_booking:type_name -> ozon.dev.mc2.api.BusBooking
	1,  // 2: ozon.dev.mc2.api.Admin.BusBookingList:input_type -> ozon.dev.mc2.api.BusBookingListRequest
	3,  // 3: ozon.dev.mc2.api.Admin.BusBookingAdd:input_type -> ozon.dev.mc2.api.BusBookingAddRequest
	5,  // 4: ozon.dev.mc2.api.Admin.BusBookingGet:input_type -> ozon.dev.mc2.api.BusBookingGetRequest
	7,  // 5: ozon.dev.mc2.api.Admin.BusBookingChangeSeat:input_type -> ozon.dev.mc2.api.BusBookingChangeSeatRequest
	9,  // 6: ozon.dev.mc2.api.Admin.BusBookingChangeDateSeat:input_type -> ozon.dev.mc2.api.BusBookingChangeDateSeatRequest
	11, // 7: ozon.dev.mc2.api.Admin.BusBookingDelete:input_type -> ozon.dev.mc2.api.BusBookingDeleteRequest
	2,  // 8: ozon.dev.mc2.api.Admin.BusBookingList:output_type -> ozon.dev.mc2.api.BusBookingListResponse
	4,  // 9: ozon.dev.mc2.api.Admin.BusBookingAdd:output_type -> ozon.dev.mc2.api.BusBookingAddResponse
	6,  // 10: ozon.dev.mc2.api.Admin.BusBookingGet:output_type -> ozon.dev.mc2.api.BusBookingGetResponse
	8,  // 11: ozon.dev.mc2.api.Admin.BusBookingChangeSeat:output_type -> ozon.dev.mc2.api.BusBookingChangeSeatResponse
	10, // 12: ozon.dev.mc2.api.Admin.BusBookingChangeDateSeat:output_type -> ozon.dev.mc2.api.BusBookingChangeDateSeatResponse
	12, // 13: ozon.dev.mc2.api.Admin.BusBookingDelete:output_type -> ozon.dev.mc2.api.BusBookingDeleteResponse
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBooking); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingListRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingListResponse); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingAddRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingAddResponse); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingGetRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingGetResponse); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingChangeSeatRequest); i {
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
		file_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingChangeSeatResponse); i {
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
		file_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingChangeDateSeatRequest); i {
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
		file_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingChangeDateSeatResponse); i {
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
		file_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingDeleteRequest); i {
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
		file_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusBookingDeleteResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
