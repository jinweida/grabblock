// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: filestore.proto

package entity

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FileSegment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Index            int32            `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	Skip             int32            `protobuf:"varint,2,opt,name=skip,proto3" json:"skip,omitempty"`
	Size             int64            `protobuf:"varint,3,opt,name=size,proto3" json:"size,omitempty"`
	SegmentUuid      string           `protobuf:"bytes,4,opt,name=segmentUuid,proto3" json:"segmentUuid,omitempty"`
	SegmentSecretKey string           `protobuf:"bytes,5,opt,name=segmentSecretKey,proto3" json:"segmentSecretKey,omitempty"`
	StoreNodeInfo    []*StoreNodeInfo `protobuf:"bytes,6,rep,name=storeNodeInfo,proto3" json:"storeNodeInfo,omitempty"` //分片存储信息
	PartSign         string           `protobuf:"bytes,7,opt,name=partSign,proto3" json:"partSign,omitempty"`           //签名数据
	VolumeCount      int64            `protobuf:"varint,8,opt,name=volumeCount,proto3" json:"volumeCount,omitempty"`    //该分片有多少721V可以挖
	Compress         bool             `protobuf:"varint,9,opt,name=compress,proto3" json:"compress,omitempty"`          //是否开启压缩
}

func (x *FileSegment) Reset() {
	*x = FileSegment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filestore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileSegment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileSegment) ProtoMessage() {}

func (x *FileSegment) ProtoReflect() protoreflect.Message {
	mi := &file_filestore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileSegment.ProtoReflect.Descriptor instead.
func (*FileSegment) Descriptor() ([]byte, []int) {
	return file_filestore_proto_rawDescGZIP(), []int{0}
}

func (x *FileSegment) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *FileSegment) GetSkip() int32 {
	if x != nil {
		return x.Skip
	}
	return 0
}

func (x *FileSegment) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *FileSegment) GetSegmentUuid() string {
	if x != nil {
		return x.SegmentUuid
	}
	return ""
}

func (x *FileSegment) GetSegmentSecretKey() string {
	if x != nil {
		return x.SegmentSecretKey
	}
	return ""
}

func (x *FileSegment) GetStoreNodeInfo() []*StoreNodeInfo {
	if x != nil {
		return x.StoreNodeInfo
	}
	return nil
}

func (x *FileSegment) GetPartSign() string {
	if x != nil {
		return x.PartSign
	}
	return ""
}

func (x *FileSegment) GetVolumeCount() int64 {
	if x != nil {
		return x.VolumeCount
	}
	return 0
}

func (x *FileSegment) GetCompress() bool {
	if x != nil {
		return x.Compress
	}
	return false
}

type FileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From     string `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`          //从哪个地址来的
	FileSize int64  `protobuf:"varint,2,opt,name=fileSize,proto3" json:"fileSize,omitempty"` //文件大小
	Filename string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`  //文件名
	FileDesc string `protobuf:"bytes,4,opt,name=fileDesc,proto3" json:"fileDesc,omitempty"`  //文件描述
	Domain   string `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`      //文件所属域
	// string	symbol721c = 6;// 721token的名称
	Parts       []*FileSegment `protobuf:"bytes,7,rep,name=parts,proto3" json:"parts,omitempty"`              //
	RepeatCount int32          `protobuf:"varint,8,opt,name=repeatCount,proto3" json:"repeatCount,omitempty"` //备份个数
	SliceCount  int32          `protobuf:"varint,9,opt,name=sliceCount,proto3" json:"sliceCount,omitempty"`   //分片数量
	// int64	totalVolumeCount = 10;//该文件一共有多少个721V
	StoreDeadLine   int64  `protobuf:"varint,11,opt,name=storeDeadLine,proto3" json:"storeDeadLine,omitempty"`     //存储时间限制
	OnChainDeadLine int64  `protobuf:"varint,12,opt,name=onChainDeadLine,proto3" json:"onChainDeadLine,omitempty"` //上链时间限制
	FileSign        string `protobuf:"bytes,13,opt,name=fileSign,proto3" json:"fileSign,omitempty"`                //文件内容MD5
	ExData          []byte `protobuf:"bytes,14,opt,name=exData,proto3" json:"exData,omitempty"`                    //附加信息
}

func (x *FileInfo) Reset() {
	*x = FileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filestore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileInfo) ProtoMessage() {}

func (x *FileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filestore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileInfo.ProtoReflect.Descriptor instead.
func (*FileInfo) Descriptor() ([]byte, []int) {
	return file_filestore_proto_rawDescGZIP(), []int{1}
}

func (x *FileInfo) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *FileInfo) GetFileSize() int64 {
	if x != nil {
		return x.FileSize
	}
	return 0
}

func (x *FileInfo) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

func (x *FileInfo) GetFileDesc() string {
	if x != nil {
		return x.FileDesc
	}
	return ""
}

func (x *FileInfo) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *FileInfo) GetParts() []*FileSegment {
	if x != nil {
		return x.Parts
	}
	return nil
}

func (x *FileInfo) GetRepeatCount() int32 {
	if x != nil {
		return x.RepeatCount
	}
	return 0
}

func (x *FileInfo) GetSliceCount() int32 {
	if x != nil {
		return x.SliceCount
	}
	return 0
}

func (x *FileInfo) GetStoreDeadLine() int64 {
	if x != nil {
		return x.StoreDeadLine
	}
	return 0
}

func (x *FileInfo) GetOnChainDeadLine() int64 {
	if x != nil {
		return x.OnChainDeadLine
	}
	return 0
}

func (x *FileInfo) GetFileSign() string {
	if x != nil {
		return x.FileSign
	}
	return ""
}

func (x *FileInfo) GetExData() []byte {
	if x != nil {
		return x.ExData
	}
	return nil
}

type StoreNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileStoreName  string `protobuf:"bytes,1,opt,name=fileStoreName,proto3" json:"fileStoreName,omitempty"`   //文件存储名称
	NodesCoAddress string `protobuf:"bytes,2,opt,name=nodesCoAddress,proto3" json:"nodesCoAddress,omitempty"` //在哪些节点
	NodesUrl       string `protobuf:"bytes,3,opt,name=nodesUrl,proto3" json:"nodesUrl,omitempty"`             //在哪些节点
}

func (x *StoreNodeInfo) Reset() {
	*x = StoreNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filestore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreNodeInfo) ProtoMessage() {}

func (x *StoreNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filestore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreNodeInfo.ProtoReflect.Descriptor instead.
func (*StoreNodeInfo) Descriptor() ([]byte, []int) {
	return file_filestore_proto_rawDescGZIP(), []int{2}
}

func (x *StoreNodeInfo) GetFileStoreName() string {
	if x != nil {
		return x.FileStoreName
	}
	return ""
}

func (x *StoreNodeInfo) GetNodesCoAddress() string {
	if x != nil {
		return x.NodesCoAddress
	}
	return ""
}

func (x *StoreNodeInfo) GetNodesUrl() string {
	if x != nil {
		return x.NodesUrl
	}
	return ""
}

type FileStorageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthoriseAddrs  []string          `protobuf:"bytes,2,rep,name=authoriseAddrs,proto3" json:"authoriseAddrs,omitempty"`                                                                                      //授权账户地址列表
	Freeze          bool              `protobuf:"varint,4,opt,name=freeze,proto3" json:"freeze,omitempty"`                                                                                                     //冻结状态
	Owner           string            `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`                                                                                                        //文件上传者地址
	IsDel           bool              `protobuf:"varint,7,opt,name=isDel,proto3" json:"isDel,omitempty"`                                                                                                       //文件是否被删除
	DelReviewer     [][]byte          `protobuf:"bytes,8,rep,name=delReviewer,proto3" json:"delReviewer,omitempty"`                                                                                            //文件删除审核人
	AgreeRecord     map[string]string `protobuf:"bytes,9,rep,name=agreeRecord,proto3" json:"agreeRecord,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`    //同意删除记录 key:address;value:txHash
	RefuseRecord    map[string]string `protobuf:"bytes,10,rep,name=refuseRecord,proto3" json:"refuseRecord,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"` //拒绝删除记录 key:address;value:txHash
	Domain          string            `protobuf:"bytes,11,opt,name=domain,proto3" json:"domain,omitempty"`
	LastVersionHash []byte            `protobuf:"bytes,12,opt,name=lastVersionHash,proto3" json:"lastVersionHash,omitempty"`
	LastUpdatedAt   int64             `protobuf:"varint,13,opt,name=lastUpdatedAt,proto3" json:"lastUpdatedAt,omitempty"`
}

func (x *FileStorageInfo) Reset() {
	*x = FileStorageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filestore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStorageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStorageInfo) ProtoMessage() {}

func (x *FileStorageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filestore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStorageInfo.ProtoReflect.Descriptor instead.
func (*FileStorageInfo) Descriptor() ([]byte, []int) {
	return file_filestore_proto_rawDescGZIP(), []int{3}
}

func (x *FileStorageInfo) GetAuthoriseAddrs() []string {
	if x != nil {
		return x.AuthoriseAddrs
	}
	return nil
}

func (x *FileStorageInfo) GetFreeze() bool {
	if x != nil {
		return x.Freeze
	}
	return false
}

func (x *FileStorageInfo) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *FileStorageInfo) GetIsDel() bool {
	if x != nil {
		return x.IsDel
	}
	return false
}

func (x *FileStorageInfo) GetDelReviewer() [][]byte {
	if x != nil {
		return x.DelReviewer
	}
	return nil
}

func (x *FileStorageInfo) GetAgreeRecord() map[string]string {
	if x != nil {
		return x.AgreeRecord
	}
	return nil
}

func (x *FileStorageInfo) GetRefuseRecord() map[string]string {
	if x != nil {
		return x.RefuseRecord
	}
	return nil
}

func (x *FileStorageInfo) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *FileStorageInfo) GetLastVersionHash() []byte {
	if x != nil {
		return x.LastVersionHash
	}
	return nil
}

func (x *FileStorageInfo) GetLastUpdatedAt() int64 {
	if x != nil {
		return x.LastUpdatedAt
	}
	return 0
}

type FileStorageItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncFileInfo     []byte `protobuf:"bytes,1,opt,name=encFileInfo,proto3" json:"encFileInfo,omitempty"` //加密的文件信息
	UpdatedAt       int64  `protobuf:"varint,2,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	DesKey          []byte `protobuf:"bytes,3,opt,name=desKey,proto3" json:"desKey,omitempty"`       //des加密密钥。chainkey公钥加密密文形式存储
	CKVersion       string `protobuf:"bytes,6,opt,name=CKVersion,proto3" json:"CKVersion,omitempty"` //chainkey version
	PrevVersionHash []byte `protobuf:"bytes,7,opt,name=prevVersionHash,proto3" json:"prevVersionHash,omitempty"`
}

func (x *FileStorageItem) Reset() {
	*x = FileStorageItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filestore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileStorageItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileStorageItem) ProtoMessage() {}

func (x *FileStorageItem) ProtoReflect() protoreflect.Message {
	mi := &file_filestore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileStorageItem.ProtoReflect.Descriptor instead.
func (*FileStorageItem) Descriptor() ([]byte, []int) {
	return file_filestore_proto_rawDescGZIP(), []int{4}
}

func (x *FileStorageItem) GetEncFileInfo() []byte {
	if x != nil {
		return x.EncFileInfo
	}
	return nil
}

func (x *FileStorageItem) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *FileStorageItem) GetDesKey() []byte {
	if x != nil {
		return x.DesKey
	}
	return nil
}

func (x *FileStorageItem) GetCKVersion() string {
	if x != nil {
		return x.CKVersion
	}
	return ""
}

func (x *FileStorageItem) GetPrevVersionHash() []byte {
	if x != nil {
		return x.PrevVersionHash
	}
	return nil
}

type DataStorageInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AuthoriseAddrs  [][]byte `protobuf:"bytes,2,rep,name=authoriseAddrs,proto3" json:"authoriseAddrs,omitempty"` //授权账户地址列表
	Owner           []byte   `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`                   //文件上传者地址
	IsDel           bool     `protobuf:"varint,7,opt,name=isDel,proto3" json:"isDel,omitempty"`                  //文件是否被删除
	Domain          string   `protobuf:"bytes,11,opt,name=domain,proto3" json:"domain,omitempty"`
	LastVersionHash []byte   `protobuf:"bytes,12,opt,name=lastVersionHash,proto3" json:"lastVersionHash,omitempty"`
	LastUpdatedAt   int64    `protobuf:"varint,13,opt,name=lastUpdatedAt,proto3" json:"lastUpdatedAt,omitempty"`
}

func (x *DataStorageInfo) Reset() {
	*x = DataStorageInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filestore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStorageInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStorageInfo) ProtoMessage() {}

func (x *DataStorageInfo) ProtoReflect() protoreflect.Message {
	mi := &file_filestore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStorageInfo.ProtoReflect.Descriptor instead.
func (*DataStorageInfo) Descriptor() ([]byte, []int) {
	return file_filestore_proto_rawDescGZIP(), []int{5}
}

func (x *DataStorageInfo) GetAuthoriseAddrs() [][]byte {
	if x != nil {
		return x.AuthoriseAddrs
	}
	return nil
}

func (x *DataStorageInfo) GetOwner() []byte {
	if x != nil {
		return x.Owner
	}
	return nil
}

func (x *DataStorageInfo) GetIsDel() bool {
	if x != nil {
		return x.IsDel
	}
	return false
}

func (x *DataStorageInfo) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *DataStorageInfo) GetLastVersionHash() []byte {
	if x != nil {
		return x.LastVersionHash
	}
	return nil
}

func (x *DataStorageInfo) GetLastUpdatedAt() int64 {
	if x != nil {
		return x.LastUpdatedAt
	}
	return 0
}

type DataStorageItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EncData         []byte `protobuf:"bytes,1,opt,name=encData,proto3" json:"encData,omitempty"`
	UpdatedAt       int64  `protobuf:"varint,2,opt,name=updatedAt,proto3" json:"updatedAt,omitempty"`
	PrevVersionHash []byte `protobuf:"bytes,3,opt,name=prevVersionHash,proto3" json:"prevVersionHash,omitempty"`
	DesKey          []byte `protobuf:"bytes,4,opt,name=desKey,proto3" json:"desKey,omitempty"`
	CKVersion       string `protobuf:"bytes,5,opt,name=CKVersion,proto3" json:"CKVersion,omitempty"`
}

func (x *DataStorageItem) Reset() {
	*x = DataStorageItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filestore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataStorageItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataStorageItem) ProtoMessage() {}

func (x *DataStorageItem) ProtoReflect() protoreflect.Message {
	mi := &file_filestore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataStorageItem.ProtoReflect.Descriptor instead.
func (*DataStorageItem) Descriptor() ([]byte, []int) {
	return file_filestore_proto_rawDescGZIP(), []int{6}
}

func (x *DataStorageItem) GetEncData() []byte {
	if x != nil {
		return x.EncData
	}
	return nil
}

func (x *DataStorageItem) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *DataStorageItem) GetPrevVersionHash() []byte {
	if x != nil {
		return x.PrevVersionHash
	}
	return nil
}

func (x *DataStorageItem) GetDesKey() []byte {
	if x != nil {
		return x.DesKey
	}
	return nil
}

func (x *DataStorageItem) GetCKVersion() string {
	if x != nil {
		return x.CKVersion
	}
	return ""
}

type FileTrunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FileIdx [][]byte `protobuf:"bytes,1,rep,name=fileIdx,proto3" json:"fileIdx,omitempty"` //文件索引
}

func (x *FileTrunk) Reset() {
	*x = FileTrunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_filestore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileTrunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileTrunk) ProtoMessage() {}

func (x *FileTrunk) ProtoReflect() protoreflect.Message {
	mi := &file_filestore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileTrunk.ProtoReflect.Descriptor instead.
func (*FileTrunk) Descriptor() ([]byte, []int) {
	return file_filestore_proto_rawDescGZIP(), []int{7}
}

func (x *FileTrunk) GetFileIdx() [][]byte {
	if x != nil {
		return x.FileIdx
	}
	return nil
}

var File_filestore_proto protoreflect.FileDescriptor

var file_filestore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0xb0, 0x02, 0x0a, 0x0b, 0x46, 0x69,
	0x6c, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64,
	0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12,
	0x12, 0x0a, 0x04, 0x73, 0x6b, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x73,
	0x6b, 0x69, 0x70, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x65, 0x67, 0x6d, 0x65,
	0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x65,
	0x67, 0x6d, 0x65, 0x6e, 0x74, 0x55, 0x75, 0x69, 0x64, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x65, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72,
	0x65, 0x74, 0x4b, 0x65, 0x79, 0x12, 0x3b, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x20,
	0x0a, 0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0b, 0x76, 0x6f, 0x6c, 0x75, 0x6d, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x22, 0xfb, 0x02, 0x0a,
	0x08, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f,
	0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x1a, 0x0a,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c,
	0x65, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x73,
	0x63, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x70, 0x61, 0x72,
	0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x70,
	0x61, 0x72, 0x74, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x72, 0x65, 0x70, 0x65, 0x61, 0x74, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x72, 0x65, 0x70, 0x65, 0x61,
	0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x6c, 0x69, 0x63, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x6c, 0x69, 0x63,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x44,
	0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x28, 0x0a, 0x0f,
	0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x18,
	0x0c, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x6f, 0x6e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x44, 0x65,
	0x61, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x69,
	0x67, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x06, 0x65, 0x78, 0x44, 0x61, 0x74, 0x61, 0x22, 0x79, 0x0a, 0x0d, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x24, 0x0a, 0x0d, 0x66,
	0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x26, 0x0a, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x43, 0x6f, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x6e, 0x6f, 0x64, 0x65, 0x73,
	0x43, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64,
	0x65, 0x73, 0x55, 0x72, 0x6c, 0x22, 0xa3, 0x04, 0x0a, 0x0f, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72,
	0x73, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x06, 0x66, 0x72, 0x65, 0x65, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x14, 0x0a, 0x05, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05,
	0x69, 0x73, 0x44, 0x65, 0x6c, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x64, 0x65, 0x6c, 0x52,
	0x65, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x12, 0x4a, 0x0a, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x18, 0x09, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67,
	0x65, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x41, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72,
	0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x67, 0x72, 0x65, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x12, 0x4d, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x73, 0x65, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x6e,
	0x66, 0x6f, 0x2e, 0x52, 0x65, 0x66, 0x75, 0x73, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x75, 0x73, 0x65, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x6c, 0x61,
	0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x67,
	0x72, 0x65, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3f, 0x0a, 0x11, 0x52, 0x65,
	0x66, 0x75, 0x73, 0x65, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xb1, 0x01, 0x0a, 0x0f,
	0x46, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x74, 0x65, 0x6d, 0x12,
	0x20, 0x0a, 0x0b, 0x65, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b, 0x65, 0x6e, 0x63, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x1c, 0x0a, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x64, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x06, 0x64, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x43, 0x4b, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x43, 0x4b, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x28, 0x0a, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f,
	0x70, 0x72, 0x65, 0x76, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x22,
	0xcd, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49,
	0x6e, 0x66, 0x6f, 0x12, 0x26, 0x0a, 0x0e, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x73, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0e, 0x61, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x69, 0x73, 0x65, 0x41, 0x64, 0x64, 0x72, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65,
	0x72, 0x12, 0x14, 0x0a, 0x05, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x05, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12,
	0x28, 0x0a, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61,
	0x73, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x6c, 0x61, 0x73, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x24, 0x0a, 0x0d, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22,
	0xa9, 0x01, 0x0a, 0x0f, 0x44, 0x61, 0x74, 0x61, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49,
	0x74, 0x65, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x63, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x65, 0x6e, 0x63, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x0a,
	0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x28, 0x0a, 0x0f, 0x70,
	0x72, 0x65, 0x76, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x70, 0x72, 0x65, 0x76, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x64, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a,
	0x09, 0x43, 0x4b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x43, 0x4b, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x09, 0x46,
	0x69, 0x6c, 0x65, 0x54, 0x72, 0x75, 0x6e, 0x6b, 0x12, 0x18, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x65,
	0x49, 0x64, 0x78, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x65, 0x49,
	0x64, 0x78, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_filestore_proto_rawDescOnce sync.Once
	file_filestore_proto_rawDescData = file_filestore_proto_rawDesc
)

func file_filestore_proto_rawDescGZIP() []byte {
	file_filestore_proto_rawDescOnce.Do(func() {
		file_filestore_proto_rawDescData = protoimpl.X.CompressGZIP(file_filestore_proto_rawDescData)
	})
	return file_filestore_proto_rawDescData
}

var file_filestore_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_filestore_proto_goTypes = []interface{}{
	(*FileSegment)(nil),     // 0: entity.FileSegment
	(*FileInfo)(nil),        // 1: entity.FileInfo
	(*StoreNodeInfo)(nil),   // 2: entity.StoreNodeInfo
	(*FileStorageInfo)(nil), // 3: entity.FileStorageInfo
	(*FileStorageItem)(nil), // 4: entity.FileStorageItem
	(*DataStorageInfo)(nil), // 5: entity.DataStorageInfo
	(*DataStorageItem)(nil), // 6: entity.DataStorageItem
	(*FileTrunk)(nil),       // 7: entity.FileTrunk
	nil,                     // 8: entity.FileStorageInfo.AgreeRecordEntry
	nil,                     // 9: entity.FileStorageInfo.RefuseRecordEntry
}
var file_filestore_proto_depIdxs = []int32{
	2, // 0: entity.FileSegment.storeNodeInfo:type_name -> entity.StoreNodeInfo
	0, // 1: entity.FileInfo.parts:type_name -> entity.FileSegment
	8, // 2: entity.FileStorageInfo.agreeRecord:type_name -> entity.FileStorageInfo.AgreeRecordEntry
	9, // 3: entity.FileStorageInfo.refuseRecord:type_name -> entity.FileStorageInfo.RefuseRecordEntry
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_filestore_proto_init() }
func file_filestore_proto_init() {
	if File_filestore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_filestore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileSegment); i {
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
		file_filestore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileInfo); i {
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
		file_filestore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreNodeInfo); i {
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
		file_filestore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileStorageInfo); i {
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
		file_filestore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileStorageItem); i {
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
		file_filestore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStorageInfo); i {
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
		file_filestore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataStorageItem); i {
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
		file_filestore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileTrunk); i {
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
			RawDescriptor: file_filestore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_filestore_proto_goTypes,
		DependencyIndexes: file_filestore_proto_depIdxs,
		MessageInfos:      file_filestore_proto_msgTypes,
	}.Build()
	File_filestore_proto = out.File
	file_filestore_proto_rawDesc = nil
	file_filestore_proto_goTypes = nil
	file_filestore_proto_depIdxs = nil
}
