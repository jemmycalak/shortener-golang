// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.2
// source: protos/urlshortener.proto

package pbuild

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UrlShortenerProto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	BaseUrl  string `protobuf:"bytes,2,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	Redirect string `protobuf:"bytes,3,opt,name=redirect,proto3" json:"redirect,omitempty"`
	Clicked  int64  `protobuf:"varint,4,opt,name=clicked,proto3" json:"clicked,omitempty"`
	UserId   int64  `protobuf:"varint,5,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *UrlShortenerProto) Reset() {
	*x = UrlShortenerProto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_urlshortener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UrlShortenerProto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UrlShortenerProto) ProtoMessage() {}

func (x *UrlShortenerProto) ProtoReflect() protoreflect.Message {
	mi := &file_protos_urlshortener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UrlShortenerProto.ProtoReflect.Descriptor instead.
func (*UrlShortenerProto) Descriptor() ([]byte, []int) {
	return file_protos_urlshortener_proto_rawDescGZIP(), []int{0}
}

func (x *UrlShortenerProto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UrlShortenerProto) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *UrlShortenerProto) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

func (x *UrlShortenerProto) GetClicked() int64 {
	if x != nil {
		return x.Clicked
	}
	return 0
}

func (x *UrlShortenerProto) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetStatisticByUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseUrl string `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	UserId  int64  `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *GetStatisticByUrlRequest) Reset() {
	*x = GetStatisticByUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_urlshortener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatisticByUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatisticByUrlRequest) ProtoMessage() {}

func (x *GetStatisticByUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_urlshortener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatisticByUrlRequest.ProtoReflect.Descriptor instead.
func (*GetStatisticByUrlRequest) Descriptor() ([]byte, []int) {
	return file_protos_urlshortener_proto_rawDescGZIP(), []int{1}
}

func (x *GetStatisticByUrlRequest) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *GetStatisticByUrlRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type GetStatisticByUrlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllTime int64 `protobuf:"varint,1,opt,name=all_time,json=allTime,proto3" json:"all_time,omitempty"`
	Day     int64 `protobuf:"varint,2,opt,name=day,proto3" json:"day,omitempty"`
	Week    int64 `protobuf:"varint,3,opt,name=week,proto3" json:"week,omitempty"`
	Month   int64 `protobuf:"varint,4,opt,name=month,proto3" json:"month,omitempty"`
}

func (x *GetStatisticByUrlResponse) Reset() {
	*x = GetStatisticByUrlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_urlshortener_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatisticByUrlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatisticByUrlResponse) ProtoMessage() {}

func (x *GetStatisticByUrlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protos_urlshortener_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatisticByUrlResponse.ProtoReflect.Descriptor instead.
func (*GetStatisticByUrlResponse) Descriptor() ([]byte, []int) {
	return file_protos_urlshortener_proto_rawDescGZIP(), []int{2}
}

func (x *GetStatisticByUrlResponse) GetAllTime() int64 {
	if x != nil {
		return x.AllTime
	}
	return 0
}

func (x *GetStatisticByUrlResponse) GetDay() int64 {
	if x != nil {
		return x.Day
	}
	return 0
}

func (x *GetStatisticByUrlResponse) GetWeek() int64 {
	if x != nil {
		return x.Week
	}
	return 0
}

func (x *GetStatisticByUrlResponse) GetMonth() int64 {
	if x != nil {
		return x.Month
	}
	return 0
}

type SearchByShortUrlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseUrl string `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
}

func (x *SearchByShortUrlRequest) Reset() {
	*x = SearchByShortUrlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_urlshortener_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchByShortUrlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchByShortUrlRequest) ProtoMessage() {}

func (x *SearchByShortUrlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_urlshortener_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchByShortUrlRequest.ProtoReflect.Descriptor instead.
func (*SearchByShortUrlRequest) Descriptor() ([]byte, []int) {
	return file_protos_urlshortener_proto_rawDescGZIP(), []int{3}
}

func (x *SearchByShortUrlRequest) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

type IncreaseAccessedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseUrl   string `protobuf:"bytes,1,opt,name=base_url,json=baseUrl,proto3" json:"base_url,omitempty"`
	IpAddress string `protobuf:"bytes,2,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
}

func (x *IncreaseAccessedRequest) Reset() {
	*x = IncreaseAccessedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_urlshortener_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncreaseAccessedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncreaseAccessedRequest) ProtoMessage() {}

func (x *IncreaseAccessedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protos_urlshortener_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncreaseAccessedRequest.ProtoReflect.Descriptor instead.
func (*IncreaseAccessedRequest) Descriptor() ([]byte, []int) {
	return file_protos_urlshortener_proto_rawDescGZIP(), []int{4}
}

func (x *IncreaseAccessedRequest) GetBaseUrl() string {
	if x != nil {
		return x.BaseUrl
	}
	return ""
}

func (x *IncreaseAccessedRequest) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

var File_protos_urlshortener_proto protoreflect.FileDescriptor

var file_protos_urlshortener_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x75, 0x72, 0x6c, 0x73, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8d, 0x01, 0x0a, 0x11, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65,
	0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72,
	0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x6c, 0x69, 0x63, 0x6b, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x4e, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x42, 0x79, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x22, 0x72, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63,
	0x42, 0x79, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a,
	0x08, 0x61, 0x6c, 0x6c, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x61, 0x6c, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x64, 0x61, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x64, 0x61, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x77, 0x65,
	0x65, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x77, 0x65, 0x65, 0x6b, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x6e, 0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x6d,
	0x6f, 0x6e, 0x74, 0x68, 0x22, 0x34, 0x0a, 0x17, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79,
	0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c, 0x22, 0x53, 0x0a, 0x17, 0x49, 0x6e,
	0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61, 0x73, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x32,
	0x87, 0x03, 0x0a, 0x18, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3e, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x72, 0x6c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x3c, 0x0a, 0x04,
	0x45, 0x64, 0x69, 0x74, 0x12, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x72,
	0x6c, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x72, 0x6c, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x50, 0x0a, 0x09, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x12, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x42, 0x79, 0x55,
	0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x42,
	0x79, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x10,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x42, 0x79, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c,
	0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68,
	0x42, 0x79, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x55, 0x72, 0x6c, 0x53, 0x68,
	0x6f, 0x72, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x4b, 0x0a, 0x10,
	0x49, 0x6e, 0x63, 0x72, 0x65, 0x61, 0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64,
	0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x63, 0x72, 0x65, 0x61,
	0x73, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x70, 0x62,
	0x75, 0x69, 0x6c, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_urlshortener_proto_rawDescOnce sync.Once
	file_protos_urlshortener_proto_rawDescData = file_protos_urlshortener_proto_rawDesc
)

func file_protos_urlshortener_proto_rawDescGZIP() []byte {
	file_protos_urlshortener_proto_rawDescOnce.Do(func() {
		file_protos_urlshortener_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_urlshortener_proto_rawDescData)
	})
	return file_protos_urlshortener_proto_rawDescData
}

var file_protos_urlshortener_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_protos_urlshortener_proto_goTypes = []interface{}{
	(*UrlShortenerProto)(nil),         // 0: protos.UrlShortenerProto
	(*GetStatisticByUrlRequest)(nil),  // 1: protos.GetStatisticByUrlRequest
	(*GetStatisticByUrlResponse)(nil), // 2: protos.GetStatisticByUrlResponse
	(*SearchByShortUrlRequest)(nil),   // 3: protos.SearchByShortUrlRequest
	(*IncreaseAccessedRequest)(nil),   // 4: protos.IncreaseAccessedRequest
	(*emptypb.Empty)(nil),             // 5: google.protobuf.Empty
}
var file_protos_urlshortener_proto_depIdxs = []int32{
	0, // 0: protos.UrlShortenerServiceProto.Create:input_type -> protos.UrlShortenerProto
	0, // 1: protos.UrlShortenerServiceProto.Edit:input_type -> protos.UrlShortenerProto
	1, // 2: protos.UrlShortenerServiceProto.Statistic:input_type -> protos.GetStatisticByUrlRequest
	3, // 3: protos.UrlShortenerServiceProto.SearchByShortUrl:input_type -> protos.SearchByShortUrlRequest
	4, // 4: protos.UrlShortenerServiceProto.IncreaseAccessed:input_type -> protos.IncreaseAccessedRequest
	0, // 5: protos.UrlShortenerServiceProto.Create:output_type -> protos.UrlShortenerProto
	0, // 6: protos.UrlShortenerServiceProto.Edit:output_type -> protos.UrlShortenerProto
	2, // 7: protos.UrlShortenerServiceProto.Statistic:output_type -> protos.GetStatisticByUrlResponse
	0, // 8: protos.UrlShortenerServiceProto.SearchByShortUrl:output_type -> protos.UrlShortenerProto
	5, // 9: protos.UrlShortenerServiceProto.IncreaseAccessed:output_type -> google.protobuf.Empty
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_urlshortener_proto_init() }
func file_protos_urlshortener_proto_init() {
	if File_protos_urlshortener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_urlshortener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UrlShortenerProto); i {
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
		file_protos_urlshortener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatisticByUrlRequest); i {
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
		file_protos_urlshortener_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatisticByUrlResponse); i {
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
		file_protos_urlshortener_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchByShortUrlRequest); i {
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
		file_protos_urlshortener_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncreaseAccessedRequest); i {
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
			RawDescriptor: file_protos_urlshortener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_urlshortener_proto_goTypes,
		DependencyIndexes: file_protos_urlshortener_proto_depIdxs,
		MessageInfos:      file_protos_urlshortener_proto_msgTypes,
	}.Build()
	File_protos_urlshortener_proto = out.File
	file_protos_urlshortener_proto_rawDesc = nil
	file_protos_urlshortener_proto_goTypes = nil
	file_protos_urlshortener_proto_depIdxs = nil
}