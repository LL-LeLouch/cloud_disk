// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: disk.proto

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

type FileDetail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Identity  string `protobuf:"bytes,2,opt,name=identity,proto3" json:"identity,omitempty"`
	Uid       string `protobuf:"bytes,3,opt,name=uid,proto3" json:"uid,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Pid       int64  `protobuf:"varint,5,opt,name=Pid,proto3" json:"Pid,omitempty"`
	DelState  int64  `protobuf:"varint,6,opt,name=DelState,proto3" json:"DelState,omitempty"`
	UpdatedAt int64  `protobuf:"varint,7,opt,name=UpdatedAt,proto3" json:"UpdatedAt,omitempty"`
}

func (x *FileDetail) Reset() {
	*x = FileDetail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileDetail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileDetail) ProtoMessage() {}

func (x *FileDetail) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileDetail.ProtoReflect.Descriptor instead.
func (*FileDetail) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{0}
}

func (x *FileDetail) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *FileDetail) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *FileDetail) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *FileDetail) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileDetail) GetPid() int64 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *FileDetail) GetDelState() int64 {
	if x != nil {
		return x.DelState
	}
	return 0
}

func (x *FileDetail) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

type ListFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
}

func (x *ListFileReq) Reset() {
	*x = ListFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileReq) ProtoMessage() {}

func (x *ListFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileReq.ProtoReflect.Descriptor instead.
func (*ListFileReq) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{1}
}

func (x *ListFileReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

type ListFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDetail *FileDetail `protobuf:"bytes,1,opt,name=fileDetail,proto3" json:"fileDetail,omitempty"`
}

func (x *ListFileResp) Reset() {
	*x = ListFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListFileResp) ProtoMessage() {}

func (x *ListFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListFileResp.ProtoReflect.Descriptor instead.
func (*ListFileResp) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{2}
}

func (x *ListFileResp) GetFileDetail() *FileDetail {
	if x != nil {
		return x.FileDetail
	}
	return nil
}

type FileUploadPrepareRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Md5  string `protobuf:"bytes,1,opt,name=md5,proto3" json:"md5,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Ext  string `protobuf:"bytes,3,opt,name=ext,proto3" json:"ext,omitempty"`
}

func (x *FileUploadPrepareRep) Reset() {
	*x = FileUploadPrepareRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadPrepareRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadPrepareRep) ProtoMessage() {}

func (x *FileUploadPrepareRep) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadPrepareRep.ProtoReflect.Descriptor instead.
func (*FileUploadPrepareRep) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{3}
}

func (x *FileUploadPrepareRep) GetMd5() string {
	if x != nil {
		return x.Md5
	}
	return ""
}

func (x *FileUploadPrepareRep) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *FileUploadPrepareRep) GetExt() string {
	if x != nil {
		return x.Ext
	}
	return ""
}

type FileUploadPrepareResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	UploadId string `protobuf:"bytes,2,opt,name=uploadId,proto3" json:"uploadId,omitempty"`
	Key      string `protobuf:"bytes,3,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *FileUploadPrepareResp) Reset() {
	*x = FileUploadPrepareResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileUploadPrepareResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileUploadPrepareResp) ProtoMessage() {}

func (x *FileUploadPrepareResp) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileUploadPrepareResp.ProtoReflect.Descriptor instead.
func (*FileUploadPrepareResp) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{4}
}

func (x *FileUploadPrepareResp) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *FileUploadPrepareResp) GetUploadId() string {
	if x != nil {
		return x.UploadId
	}
	return ""
}

func (x *FileUploadPrepareResp) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type UpdateFileReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDetail *FileDetail `protobuf:"bytes,1,opt,name=fileDetail,proto3" json:"fileDetail,omitempty"`
}

func (x *UpdateFileReq) Reset() {
	*x = UpdateFileReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileReq) ProtoMessage() {}

func (x *UpdateFileReq) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileReq.ProtoReflect.Descriptor instead.
func (*UpdateFileReq) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateFileReq) GetFileDetail() *FileDetail {
	if x != nil {
		return x.FileDetail
	}
	return nil
}

type UpdateFileResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileDetail *FileDetail `protobuf:"bytes,1,opt,name=fileDetail,proto3" json:"fileDetail,omitempty"`
}

func (x *UpdateFileResp) Reset() {
	*x = UpdateFileResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateFileResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateFileResp) ProtoMessage() {}

func (x *UpdateFileResp) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateFileResp.ProtoReflect.Descriptor instead.
func (*UpdateFileResp) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateFileResp) GetFileDetail() *FileDetail {
	if x != nil {
		return x.FileDetail
	}
	return nil
}

type StatisticsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uid      string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	DelState int64  `protobuf:"varint,2,opt,name=DelState,proto3" json:"DelState,omitempty"`
	IsUser   int64  `protobuf:"varint,3,opt,name=isUser,proto3" json:"isUser,omitempty"`
}

func (x *StatisticsReq) Reset() {
	*x = StatisticsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticsReq) ProtoMessage() {}

func (x *StatisticsReq) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticsReq.ProtoReflect.Descriptor instead.
func (*StatisticsReq) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{7}
}

func (x *StatisticsReq) GetUid() string {
	if x != nil {
		return x.Uid
	}
	return ""
}

func (x *StatisticsReq) GetDelState() int64 {
	if x != nil {
		return x.DelState
	}
	return 0
}

func (x *StatisticsReq) GetIsUser() int64 {
	if x != nil {
		return x.IsUser
	}
	return 0
}

type StatisticsResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalSize int64 `protobuf:"varint,1,opt,name=totalSize,proto3" json:"totalSize,omitempty"`
	Count     int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *StatisticsResp) Reset() {
	*x = StatisticsResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatisticsResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatisticsResp) ProtoMessage() {}

func (x *StatisticsResp) ProtoReflect() protoreflect.Message {
	mi := &file_disk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatisticsResp.ProtoReflect.Descriptor instead.
func (*StatisticsResp) Descriptor() ([]byte, []int) {
	return file_disk_proto_rawDescGZIP(), []int{8}
}

func (x *StatisticsResp) GetTotalSize() int64 {
	if x != nil {
		return x.TotalSize
	}
	return 0
}

func (x *StatisticsResp) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

var File_disk_proto protoreflect.FileDescriptor

var file_disk_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x64, 0x69, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0xaa, 0x01, 0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x50, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x50, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x1f, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x22, 0x3e,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2e,
	0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x4e,
	0x0a, 0x14, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x64, 0x35, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x64, 0x35, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03,
	0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x78, 0x74, 0x22, 0x61,
	0x0a, 0x15, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x65, 0x70,
	0x61, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x22, 0x3f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52,
	0x65, 0x71, 0x12, 0x2e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x22, 0x40, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x2e, 0x0a, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69,
	0x6c, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x0a, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65,
	0x74, 0x61, 0x69, 0x6c, 0x22, 0x55, 0x0a, 0x0d, 0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69,
	0x63, 0x73, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x69, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x44, 0x65, 0x6c, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x73, 0x55, 0x73, 0x65, 0x72, 0x22, 0x44, 0x0a, 0x0e, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1c, 0x0a,
	0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x32, 0xe9, 0x01, 0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12, 0x48, 0x0a, 0x11, 0x46, 0x69,
	0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x12,
	0x18, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50,
	0x72, 0x65, 0x70, 0x61, 0x72, 0x65, 0x52, 0x65, 0x70, 0x1a, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x0a, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2d,
	0x0a, 0x08, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x10, 0x2e, 0x70, 0x62,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_disk_proto_rawDescOnce sync.Once
	file_disk_proto_rawDescData = file_disk_proto_rawDesc
)

func file_disk_proto_rawDescGZIP() []byte {
	file_disk_proto_rawDescOnce.Do(func() {
		file_disk_proto_rawDescData = protoimpl.X.CompressGZIP(file_disk_proto_rawDescData)
	})
	return file_disk_proto_rawDescData
}

var file_disk_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_disk_proto_goTypes = []interface{}{
	(*FileDetail)(nil),            // 0: pb.FileDetail
	(*ListFileReq)(nil),           // 1: pb.ListFileReq
	(*ListFileResp)(nil),          // 2: pb.ListFileResp
	(*FileUploadPrepareRep)(nil),  // 3: pb.FileUploadPrepareRep
	(*FileUploadPrepareResp)(nil), // 4: pb.FileUploadPrepareResp
	(*UpdateFileReq)(nil),         // 5: pb.UpdateFileReq
	(*UpdateFileResp)(nil),        // 6: pb.UpdateFileResp
	(*StatisticsReq)(nil),         // 7: pb.StatisticsReq
	(*StatisticsResp)(nil),        // 8: pb.StatisticsResp
}
var file_disk_proto_depIdxs = []int32{
	0, // 0: pb.ListFileResp.fileDetail:type_name -> pb.FileDetail
	0, // 1: pb.UpdateFileReq.fileDetail:type_name -> pb.FileDetail
	0, // 2: pb.UpdateFileResp.fileDetail:type_name -> pb.FileDetail
	3, // 3: pb.disk.FileUploadPrepare:input_type -> pb.FileUploadPrepareRep
	5, // 4: pb.disk.UpdateFile:input_type -> pb.UpdateFileReq
	7, // 5: pb.disk.Statistics:input_type -> pb.StatisticsReq
	1, // 6: pb.disk.ListFile:input_type -> pb.ListFileReq
	4, // 7: pb.disk.FileUploadPrepare:output_type -> pb.FileUploadPrepareResp
	6, // 8: pb.disk.UpdateFile:output_type -> pb.UpdateFileResp
	8, // 9: pb.disk.Statistics:output_type -> pb.StatisticsResp
	2, // 10: pb.disk.ListFile:output_type -> pb.ListFileResp
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_disk_proto_init() }
func file_disk_proto_init() {
	if File_disk_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_disk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileDetail); i {
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
		file_disk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFileReq); i {
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
		file_disk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListFileResp); i {
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
		file_disk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadPrepareRep); i {
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
		file_disk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileUploadPrepareResp); i {
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
		file_disk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileReq); i {
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
		file_disk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateFileResp); i {
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
		file_disk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticsReq); i {
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
		file_disk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatisticsResp); i {
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
			RawDescriptor: file_disk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_disk_proto_goTypes,
		DependencyIndexes: file_disk_proto_depIdxs,
		MessageInfos:      file_disk_proto_msgTypes,
	}.Build()
	File_disk_proto = out.File
	file_disk_proto_rawDesc = nil
	file_disk_proto_goTypes = nil
	file_disk_proto_depIdxs = nil
}
