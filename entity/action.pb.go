// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: action.proto

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

type RetAccountMessage_AccountType int32

const (
	RetAccountMessage_ACCOUNT      RetAccountMessage_AccountType = 0
	RetAccountMessage_UNIONACCOUNT RetAccountMessage_AccountType = 1
	RetAccountMessage_CONTRACT     RetAccountMessage_AccountType = 2
)

// Enum value maps for RetAccountMessage_AccountType.
var (
	RetAccountMessage_AccountType_name = map[int32]string{
		0: "ACCOUNT",
		1: "UNIONACCOUNT",
		2: "CONTRACT",
	}
	RetAccountMessage_AccountType_value = map[string]int32{
		"ACCOUNT":      0,
		"UNIONACCOUNT": 1,
		"CONTRACT":     2,
	}
)

func (x RetAccountMessage_AccountType) Enum() *RetAccountMessage_AccountType {
	p := new(RetAccountMessage_AccountType)
	*p = x
	return p
}

func (x RetAccountMessage_AccountType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RetAccountMessage_AccountType) Descriptor() protoreflect.EnumDescriptor {
	return file_action_proto_enumTypes[0].Descriptor()
}

func (RetAccountMessage_AccountType) Type() protoreflect.EnumType {
	return &file_action_proto_enumTypes[0]
}

func (x RetAccountMessage_AccountType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RetAccountMessage_AccountType.Descriptor instead.
func (RetAccountMessage_AccountType) EnumDescriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{2, 0}
}

type RetBlockMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode int32        `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg  string       `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	Block   []*BlockInfo `protobuf:"bytes,3,rep,name=block,proto3" json:"block,omitempty"`
}

func (x *RetBlockMessage) Reset() {
	*x = RetBlockMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetBlockMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetBlockMessage) ProtoMessage() {}

func (x *RetBlockMessage) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetBlockMessage.ProtoReflect.Descriptor instead.
func (*RetBlockMessage) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{0}
}

func (x *RetBlockMessage) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *RetBlockMessage) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *RetBlockMessage) GetBlock() []*BlockInfo {
	if x != nil {
		return x.Block
	}
	return nil
}

type RetTransactionMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode     int32              `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg      string             `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	Transaction *TransactionInfo   `protobuf:"bytes,3,opt,name=transaction,proto3" json:"transaction,omitempty"`
	Status      *TransactionStatus `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"` // v2021版本
	BlockHash   string             `protobuf:"bytes,5,opt,name=blockHash,proto3" json:"blockHash,omitempty"`
}

func (x *RetTransactionMessage) Reset() {
	*x = RetTransactionMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetTransactionMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetTransactionMessage) ProtoMessage() {}

func (x *RetTransactionMessage) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetTransactionMessage.ProtoReflect.Descriptor instead.
func (*RetTransactionMessage) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{1}
}

func (x *RetTransactionMessage) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *RetTransactionMessage) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *RetTransactionMessage) GetTransaction() *TransactionInfo {
	if x != nil {
		return x.Transaction
	}
	return nil
}

func (x *RetTransactionMessage) GetStatus() *TransactionStatus {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *RetTransactionMessage) GetBlockHash() string {
	if x != nil {
		return x.BlockHash
	}
	return ""
}

type RetAccountMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RetCode         int32                         `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RetMsg          string                        `protobuf:"bytes,2,opt,name=retMsg,proto3" json:"retMsg,omitempty"`
	Type            RetAccountMessage_AccountType `protobuf:"varint,3,opt,name=type,proto3,enum=entity.RetAccountMessage_AccountType" json:"type,omitempty"`
	Address         []byte                        `protobuf:"bytes,10,opt,name=address,proto3" json:"address,omitempty"`
	Nonce           int32                         `protobuf:"varint,13,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Balance         string                        `protobuf:"bytes,14,opt,name=balance,proto3" json:"balance,omitempty"`
	Status          int32                         `protobuf:"varint,5,opt,name=status,proto3" json:"status,omitempty"`                                           //0：正常，-1：异常锁定黑名单
	StorageTrieRoot []byte                        `protobuf:"bytes,6,opt,name=storage_trie_root,json=storageTrieRoot,proto3" json:"storage_trie_root,omitempty"` //trie_sub_addressess(token,crypto,code,0001(union_subaddrs),0000(storage))
	ExtData         []byte                        `protobuf:"bytes,7,opt,name=ext_data,json=extData,proto3" json:"ext_data,omitempty"`                           //扩展信息
}

func (x *RetAccountMessage) Reset() {
	*x = RetAccountMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_action_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RetAccountMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RetAccountMessage) ProtoMessage() {}

func (x *RetAccountMessage) ProtoReflect() protoreflect.Message {
	mi := &file_action_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RetAccountMessage.ProtoReflect.Descriptor instead.
func (*RetAccountMessage) Descriptor() ([]byte, []int) {
	return file_action_proto_rawDescGZIP(), []int{2}
}

func (x *RetAccountMessage) GetRetCode() int32 {
	if x != nil {
		return x.RetCode
	}
	return 0
}

func (x *RetAccountMessage) GetRetMsg() string {
	if x != nil {
		return x.RetMsg
	}
	return ""
}

func (x *RetAccountMessage) GetType() RetAccountMessage_AccountType {
	if x != nil {
		return x.Type
	}
	return RetAccountMessage_ACCOUNT
}

func (x *RetAccountMessage) GetAddress() []byte {
	if x != nil {
		return x.Address
	}
	return nil
}

func (x *RetAccountMessage) GetNonce() int32 {
	if x != nil {
		return x.Nonce
	}
	return 0
}

func (x *RetAccountMessage) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

func (x *RetAccountMessage) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *RetAccountMessage) GetStorageTrieRoot() []byte {
	if x != nil {
		return x.StorageTrieRoot
	}
	return nil
}

func (x *RetAccountMessage) GetExtData() []byte {
	if x != nil {
		return x.ExtData
	}
	return nil
}

var File_action_proto protoreflect.FileDescriptor

var file_action_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x1a, 0x0b, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6c, 0x0a, 0x0f, 0x52, 0x65, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x12, 0x27, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x22, 0xd5, 0x01, 0x0a, 0x15, 0x52, 0x65, 0x74, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x74, 0x4d,
	0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x74, 0x4d, 0x73, 0x67,
	0x12, 0x39, 0x0a, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x31, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x48, 0x61, 0x73, 0x68, 0x22, 0xe5, 0x02, 0x0a,
	0x11, 0x52, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x72, 0x65, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x72, 0x65, 0x74, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65,
	0x74, 0x4d, 0x73, 0x67, 0x12, 0x39, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x25, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x52, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x6e,
	0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x72, 0x69,
	0x65, 0x5f, 0x72, 0x6f, 0x6f, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x73, 0x74,
	0x6f, 0x72, 0x61, 0x67, 0x65, 0x54, 0x72, 0x69, 0x65, 0x52, 0x6f, 0x6f, 0x74, 0x12, 0x19, 0x0a,
	0x08, 0x65, 0x78, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x07, 0x65, 0x78, 0x74, 0x44, 0x61, 0x74, 0x61, 0x22, 0x3a, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x41, 0x43, 0x43, 0x4f, 0x55,
	0x4e, 0x54, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x4e, 0x49, 0x4f, 0x4e, 0x41, 0x43, 0x43,
	0x4f, 0x55, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41,
	0x43, 0x54, 0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_action_proto_rawDescOnce sync.Once
	file_action_proto_rawDescData = file_action_proto_rawDesc
)

func file_action_proto_rawDescGZIP() []byte {
	file_action_proto_rawDescOnce.Do(func() {
		file_action_proto_rawDescData = protoimpl.X.CompressGZIP(file_action_proto_rawDescData)
	})
	return file_action_proto_rawDescData
}

var file_action_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_action_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_action_proto_goTypes = []interface{}{
	(RetAccountMessage_AccountType)(0), // 0: entity.RetAccountMessage.AccountType
	(*RetBlockMessage)(nil),            // 1: entity.RetBlockMessage
	(*RetTransactionMessage)(nil),      // 2: entity.RetTransactionMessage
	(*RetAccountMessage)(nil),          // 3: entity.RetAccountMessage
	(*BlockInfo)(nil),                  // 4: entity.BlockInfo
	(*TransactionInfo)(nil),            // 5: entity.TransactionInfo
	(*TransactionStatus)(nil),          // 6: entity.TransactionStatus
}
var file_action_proto_depIdxs = []int32{
	4, // 0: entity.RetBlockMessage.block:type_name -> entity.BlockInfo
	5, // 1: entity.RetTransactionMessage.transaction:type_name -> entity.TransactionInfo
	6, // 2: entity.RetTransactionMessage.status:type_name -> entity.TransactionStatus
	0, // 3: entity.RetAccountMessage.type:type_name -> entity.RetAccountMessage.AccountType
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_action_proto_init() }
func file_action_proto_init() {
	if File_action_proto != nil {
		return
	}
	file_block_proto_init()
	file_transaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_action_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetBlockMessage); i {
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
		file_action_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetTransactionMessage); i {
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
		file_action_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RetAccountMessage); i {
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
			RawDescriptor: file_action_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_action_proto_goTypes,
		DependencyIndexes: file_action_proto_depIdxs,
		EnumInfos:         file_action_proto_enumTypes,
		MessageInfos:      file_action_proto_msgTypes,
	}.Build()
	File_action_proto = out.File
	file_action_proto_rawDesc = nil
	file_action_proto_goTypes = nil
	file_action_proto_depIdxs = nil
}
