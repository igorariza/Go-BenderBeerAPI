// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: beer.proto

package v1alpha1

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

type Beer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       uint32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Brewery  string `protobuf:"bytes,3,opt,name=brewery,proto3" json:"brewery,omitempty"`
	Country  string `protobuf:"bytes,4,opt,name=country,proto3" json:"country,omitempty"`
	Price    int32  `protobuf:"varint,5,opt,name=price,proto3" json:"price,omitempty"`
	Currency int32  `protobuf:"varint,6,opt,name=currency,proto3" json:"currency,omitempty"`
}

func (x *Beer) Reset() {
	*x = Beer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Beer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Beer) ProtoMessage() {}

func (x *Beer) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Beer.ProtoReflect.Descriptor instead.
func (*Beer) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{0}
}

func (x *Beer) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Beer) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Beer) GetBrewery() string {
	if x != nil {
		return x.Brewery
	}
	return ""
}

func (x *Beer) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Beer) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Beer) GetCurrency() int32 {
	if x != nil {
		return x.Currency
	}
	return 0
}

type CreateBeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beer *Beer `protobuf:"bytes,1,opt,name=beer,proto3" json:"beer,omitempty"`
}

func (x *CreateBeerRequest) Reset() {
	*x = CreateBeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBeerRequest) ProtoMessage() {}

func (x *CreateBeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBeerRequest.ProtoReflect.Descriptor instead.
func (*CreateBeerRequest) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{1}
}

func (x *CreateBeerRequest) GetBeer() *Beer {
	if x != nil {
		return x.Beer
	}
	return nil
}

type CreateBeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beer *Beer  `protobuf:"bytes,1,opt,name=beer,proto3" json:"beer,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CreateBeerResponse) Reset() {
	*x = CreateBeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBeerResponse) ProtoMessage() {}

func (x *CreateBeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBeerResponse.ProtoReflect.Descriptor instead.
func (*CreateBeerResponse) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{2}
}

func (x *CreateBeerResponse) GetBeer() *Beer {
	if x != nil {
		return x.Beer
	}
	return nil
}

func (x *CreateBeerResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListBeersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beer *Beer `protobuf:"bytes,1,opt,name=beer,proto3" json:"beer,omitempty"`
}

func (x *ListBeersRequest) Reset() {
	*x = ListBeersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBeersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBeersRequest) ProtoMessage() {}

func (x *ListBeersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBeersRequest.ProtoReflect.Descriptor instead.
func (*ListBeersRequest) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{3}
}

func (x *ListBeersRequest) GetBeer() *Beer {
	if x != nil {
		return x.Beer
	}
	return nil
}

type ListBeersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beer []*Beer `protobuf:"bytes,1,rep,name=beer,proto3" json:"beer,omitempty"`
	Msg  string  `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ListBeersResponse) Reset() {
	*x = ListBeersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListBeersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListBeersResponse) ProtoMessage() {}

func (x *ListBeersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListBeersResponse.ProtoReflect.Descriptor instead.
func (*ListBeersResponse) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{4}
}

func (x *ListBeersResponse) GetBeer() []*Beer {
	if x != nil {
		return x.Beer
	}
	return nil
}

func (x *ListBeersResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DeleteBeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeerId uint32 `protobuf:"varint,1,opt,name=beer_id,json=beerId,proto3" json:"beer_id,omitempty"`
}

func (x *DeleteBeerRequest) Reset() {
	*x = DeleteBeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBeerRequest) ProtoMessage() {}

func (x *DeleteBeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBeerRequest.ProtoReflect.Descriptor instead.
func (*DeleteBeerRequest) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteBeerRequest) GetBeerId() uint32 {
	if x != nil {
		return x.BeerId
	}
	return 0
}

type DeleteBeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteBeerResponse) Reset() {
	*x = DeleteBeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteBeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteBeerResponse) ProtoMessage() {}

func (x *DeleteBeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteBeerResponse.ProtoReflect.Descriptor instead.
func (*DeleteBeerResponse) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteBeerResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateBeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beer *Beer `protobuf:"bytes,1,opt,name=beer,proto3" json:"beer,omitempty"`
}

func (x *UpdateBeerRequest) Reset() {
	*x = UpdateBeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBeerRequest) ProtoMessage() {}

func (x *UpdateBeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBeerRequest.ProtoReflect.Descriptor instead.
func (*UpdateBeerRequest) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{7}
}

func (x *UpdateBeerRequest) GetBeer() *Beer {
	if x != nil {
		return x.Beer
	}
	return nil
}

type UpdateBeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateBeerResponse) Reset() {
	*x = UpdateBeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateBeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateBeerResponse) ProtoMessage() {}

func (x *UpdateBeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateBeerResponse.ProtoReflect.Descriptor instead.
func (*UpdateBeerResponse) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{8}
}

func (x *UpdateBeerResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type GetOneBeerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BeerId uint32 `protobuf:"varint,1,opt,name=beer_id,json=beerId,proto3" json:"beer_id,omitempty"`
}

func (x *GetOneBeerRequest) Reset() {
	*x = GetOneBeerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneBeerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneBeerRequest) ProtoMessage() {}

func (x *GetOneBeerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneBeerRequest.ProtoReflect.Descriptor instead.
func (*GetOneBeerRequest) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{9}
}

func (x *GetOneBeerRequest) GetBeerId() uint32 {
	if x != nil {
		return x.BeerId
	}
	return 0
}

type GetOneBeerResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Beer *Beer  `protobuf:"bytes,1,opt,name=beer,proto3" json:"beer,omitempty"`
	Msg  string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetOneBeerResponse) Reset() {
	*x = GetOneBeerResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_beer_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneBeerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneBeerResponse) ProtoMessage() {}

func (x *GetOneBeerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_beer_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneBeerResponse.ProtoReflect.Descriptor instead.
func (*GetOneBeerResponse) Descriptor() ([]byte, []int) {
	return file_beer_proto_rawDescGZIP(), []int{10}
}

func (x *GetOneBeerResponse) GetBeer() *Beer {
	if x != nil {
		return x.Beer
	}
	return nil
}

func (x *GetOneBeerResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_beer_proto protoreflect.FileDescriptor

var file_beer_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x65,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x22, 0x90, 0x01, 0x0a, 0x04,
	0x42, 0x65, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x72, 0x65, 0x77,
	0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x65, 0x77, 0x65,
	0x72, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x22, 0x3c,
	0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x42, 0x65, 0x65, 0x72, 0x52, 0x04, 0x62, 0x65, 0x65, 0x72, 0x22, 0x4f, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x42, 0x65, 0x65, 0x72, 0x52, 0x04, 0x62, 0x65, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3b, 0x0a,
	0x10, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x13, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x42, 0x65, 0x65, 0x72, 0x52, 0x04, 0x62, 0x65, 0x65, 0x72, 0x22, 0x4e, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x42, 0x65, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x27, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x65,
	0x65, 0x72, 0x52, 0x04, 0x62, 0x65, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2c, 0x0a, 0x11, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x62, 0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x06, 0x62, 0x65, 0x65, 0x72, 0x49, 0x64, 0x22, 0x26, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67,
	0x22, 0x3c, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x65, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x42, 0x65, 0x65, 0x72, 0x52, 0x04, 0x62, 0x65, 0x65, 0x72, 0x22, 0x26,
	0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x2c, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65,
	0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x62,
	0x65, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x06, 0x62, 0x65,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x4f, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x42, 0x65,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x04, 0x62, 0x65,
	0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x42, 0x65, 0x65, 0x72, 0x52, 0x04, 0x62,
	0x65, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xac, 0x03, 0x0a, 0x0e, 0x42, 0x65, 0x65, 0x72, 0x41, 0x50,
	0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x4c,
	0x69, 0x73, 0x74, 0x42, 0x65, 0x65, 0x72, 0x73, 0x12, 0x1f, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x65, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x62, 0x65, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x42, 0x65,
	0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x62, 0x65, 0x65, 0x72,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51,
	0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x12, 0x20, 0x2e, 0x62,
	0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x51, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x42, 0x65, 0x65, 0x72, 0x12,
	0x20, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x21, 0x2e, 0x62, 0x65, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x42, 0x65, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3a, 0x5a, 0x38, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x69, 0x67, 0x6f, 0x72, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x2f, 0x67, 0x6f, 0x67,
	0x72, 0x70, 0x63, 0x2d, 0x61, 0x70, 0x69, 0x62, 0x65, 0x65, 0x72, 0x2f, 0x67, 0x65, 0x6e, 0x2f,
	0x67, 0x6f, 0x2f, 0x62, 0x65, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_beer_proto_rawDescOnce sync.Once
	file_beer_proto_rawDescData = file_beer_proto_rawDesc
)

func file_beer_proto_rawDescGZIP() []byte {
	file_beer_proto_rawDescOnce.Do(func() {
		file_beer_proto_rawDescData = protoimpl.X.CompressGZIP(file_beer_proto_rawDescData)
	})
	return file_beer_proto_rawDescData
}

var file_beer_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_beer_proto_goTypes = []interface{}{
	(*Beer)(nil),               // 0: beer.v1alpha1.Beer
	(*CreateBeerRequest)(nil),  // 1: beer.v1alpha1.CreateBeerRequest
	(*CreateBeerResponse)(nil), // 2: beer.v1alpha1.CreateBeerResponse
	(*ListBeersRequest)(nil),   // 3: beer.v1alpha1.ListBeersRequest
	(*ListBeersResponse)(nil),  // 4: beer.v1alpha1.ListBeersResponse
	(*DeleteBeerRequest)(nil),  // 5: beer.v1alpha1.DeleteBeerRequest
	(*DeleteBeerResponse)(nil), // 6: beer.v1alpha1.DeleteBeerResponse
	(*UpdateBeerRequest)(nil),  // 7: beer.v1alpha1.UpdateBeerRequest
	(*UpdateBeerResponse)(nil), // 8: beer.v1alpha1.UpdateBeerResponse
	(*GetOneBeerRequest)(nil),  // 9: beer.v1alpha1.GetOneBeerRequest
	(*GetOneBeerResponse)(nil), // 10: beer.v1alpha1.GetOneBeerResponse
}
var file_beer_proto_depIdxs = []int32{
	0,  // 0: beer.v1alpha1.CreateBeerRequest.beer:type_name -> beer.v1alpha1.Beer
	0,  // 1: beer.v1alpha1.CreateBeerResponse.beer:type_name -> beer.v1alpha1.Beer
	0,  // 2: beer.v1alpha1.ListBeersRequest.beer:type_name -> beer.v1alpha1.Beer
	0,  // 3: beer.v1alpha1.ListBeersResponse.beer:type_name -> beer.v1alpha1.Beer
	0,  // 4: beer.v1alpha1.UpdateBeerRequest.beer:type_name -> beer.v1alpha1.Beer
	0,  // 5: beer.v1alpha1.GetOneBeerResponse.beer:type_name -> beer.v1alpha1.Beer
	1,  // 6: beer.v1alpha1.BeerAPIService.CreateBeer:input_type -> beer.v1alpha1.CreateBeerRequest
	3,  // 7: beer.v1alpha1.BeerAPIService.ListBeers:input_type -> beer.v1alpha1.ListBeersRequest
	5,  // 8: beer.v1alpha1.BeerAPIService.DeleteBeer:input_type -> beer.v1alpha1.DeleteBeerRequest
	7,  // 9: beer.v1alpha1.BeerAPIService.UpdateBeer:input_type -> beer.v1alpha1.UpdateBeerRequest
	9,  // 10: beer.v1alpha1.BeerAPIService.GetOneBeer:input_type -> beer.v1alpha1.GetOneBeerRequest
	2,  // 11: beer.v1alpha1.BeerAPIService.CreateBeer:output_type -> beer.v1alpha1.CreateBeerResponse
	4,  // 12: beer.v1alpha1.BeerAPIService.ListBeers:output_type -> beer.v1alpha1.ListBeersResponse
	6,  // 13: beer.v1alpha1.BeerAPIService.DeleteBeer:output_type -> beer.v1alpha1.DeleteBeerResponse
	8,  // 14: beer.v1alpha1.BeerAPIService.UpdateBeer:output_type -> beer.v1alpha1.UpdateBeerResponse
	10, // 15: beer.v1alpha1.BeerAPIService.GetOneBeer:output_type -> beer.v1alpha1.GetOneBeerResponse
	11, // [11:16] is the sub-list for method output_type
	6,  // [6:11] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_beer_proto_init() }
func file_beer_proto_init() {
	if File_beer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_beer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Beer); i {
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
		file_beer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBeerRequest); i {
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
		file_beer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBeerResponse); i {
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
		file_beer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBeersRequest); i {
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
		file_beer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListBeersResponse); i {
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
		file_beer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBeerRequest); i {
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
		file_beer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteBeerResponse); i {
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
		file_beer_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBeerRequest); i {
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
		file_beer_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateBeerResponse); i {
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
		file_beer_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneBeerRequest); i {
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
		file_beer_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneBeerResponse); i {
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
			RawDescriptor: file_beer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_beer_proto_goTypes,
		DependencyIndexes: file_beer_proto_depIdxs,
		MessageInfos:      file_beer_proto_msgTypes,
	}.Build()
	File_beer_proto = out.File
	file_beer_proto_rawDesc = nil
	file_beer_proto_goTypes = nil
	file_beer_proto_depIdxs = nil
}
