// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.10.0
// source: contract.proto

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

//↓↓↓↓↓↓↓↓↓合约参数包装类↓↓↓↓↓↓↓↓↓
//人员信息
type MemberInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` //人员名称
}

func (x *MemberInfo) Reset() {
	*x = MemberInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberInfo) ProtoMessage() {}

func (x *MemberInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberInfo.ProtoReflect.Descriptor instead.
func (*MemberInfo) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{0}
}

func (x *MemberInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//人员信息列表
type MemberInfos struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name []*MemberInfo `protobuf:"bytes,1,rep,name=name,proto3" json:"name,omitempty"` //人员名称
}

func (x *MemberInfos) Reset() {
	*x = MemberInfos{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberInfos) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberInfos) ProtoMessage() {}

func (x *MemberInfos) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberInfos.ProtoReflect.Descriptor instead.
func (*MemberInfos) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{1}
}

func (x *MemberInfos) GetName() []*MemberInfo {
	if x != nil {
		return x.Name
	}
	return nil
}

//节点信息
type NodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`           //节点名称
	Cpu       string `protobuf:"bytes,2,opt,name=cpu,proto3" json:"cpu,omitempty"`             //cpu数量
	Memory    string `protobuf:"bytes,3,opt,name=memory,proto3" json:"memory,omitempty"`       //内存大小
	Disk      string `protobuf:"bytes,4,opt,name=disk,proto3" json:"disk,omitempty"`           //磁盘大小
	Bandwidth string `protobuf:"bytes,5,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"` //带宽
}

func (x *NodeInfo) Reset() {
	*x = NodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NodeInfo) ProtoMessage() {}

func (x *NodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NodeInfo.ProtoReflect.Descriptor instead.
func (*NodeInfo) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{2}
}

func (x *NodeInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NodeInfo) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *NodeInfo) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *NodeInfo) GetDisk() string {
	if x != nil {
		return x.Disk
	}
	return ""
}

func (x *NodeInfo) GetBandwidth() string {
	if x != nil {
		return x.Bandwidth
	}
	return ""
}

//组织信息
type OrgInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OrgName    string `protobuf:"bytes,1,opt,name=orgName,proto3" json:"orgName,omitempty"`       //
	OrgAddress string `protobuf:"bytes,2,opt,name=orgAddress,proto3" json:"orgAddress,omitempty"` //
}

func (x *OrgInfo) Reset() {
	*x = OrgInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrgInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrgInfo) ProtoMessage() {}

func (x *OrgInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrgInfo.ProtoReflect.Descriptor instead.
func (*OrgInfo) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{3}
}

func (x *OrgInfo) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *OrgInfo) GetOrgAddress() string {
	if x != nil {
		return x.OrgAddress
	}
	return ""
}

//业务域合约信息
type BizContractInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name   string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`     //合约名称
	Remark string `protobuf:"bytes,2,opt,name=remark,proto3" json:"remark,omitempty"` //备注
}

func (x *BizContractInfo) Reset() {
	*x = BizContractInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BizContractInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BizContractInfo) ProtoMessage() {}

func (x *BizContractInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BizContractInfo.ProtoReflect.Descriptor instead.
func (*BizContractInfo) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{4}
}

func (x *BizContractInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *BizContractInfo) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

//资源节点信息
type ResourceNodeInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeName   string `protobuf:"bytes,1,opt,name=nodeName,proto3" json:"nodeName,omitempty"`     //
	OrgName    string `protobuf:"bytes,2,opt,name=orgName,proto3" json:"orgName,omitempty"`       //
	OrgAddress string `protobuf:"bytes,3,opt,name=orgAddress,proto3" json:"orgAddress,omitempty"` //
	DSName     string `protobuf:"bytes,5,opt,name=dSName,proto3" json:"dSName,omitempty"`         //存管域名称
	Cpu        string `protobuf:"bytes,6,opt,name=cpu,proto3" json:"cpu,omitempty"`               //cpu数量
	Memory     string `protobuf:"bytes,7,opt,name=memory,proto3" json:"memory,omitempty"`         //内存大小
	Disk       string `protobuf:"bytes,8,opt,name=disk,proto3" json:"disk,omitempty"`             //磁盘大小
	Bandwidth  string `protobuf:"bytes,9,opt,name=bandwidth,proto3" json:"bandwidth,omitempty"`   //带宽
}

func (x *ResourceNodeInfo) Reset() {
	*x = ResourceNodeInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResourceNodeInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceNodeInfo) ProtoMessage() {}

func (x *ResourceNodeInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceNodeInfo.ProtoReflect.Descriptor instead.
func (*ResourceNodeInfo) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{5}
}

func (x *ResourceNodeInfo) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *ResourceNodeInfo) GetOrgName() string {
	if x != nil {
		return x.OrgName
	}
	return ""
}

func (x *ResourceNodeInfo) GetOrgAddress() string {
	if x != nil {
		return x.OrgAddress
	}
	return ""
}

func (x *ResourceNodeInfo) GetDSName() string {
	if x != nil {
		return x.DSName
	}
	return ""
}

func (x *ResourceNodeInfo) GetCpu() string {
	if x != nil {
		return x.Cpu
	}
	return ""
}

func (x *ResourceNodeInfo) GetMemory() string {
	if x != nil {
		return x.Memory
	}
	return ""
}

func (x *ResourceNodeInfo) GetDisk() string {
	if x != nil {
		return x.Disk
	}
	return ""
}

func (x *ResourceNodeInfo) GetBandwidth() string {
	if x != nil {
		return x.Bandwidth
	}
	return ""
}

//数据存管域信息
type DSDomainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` //数据存管域名称
}

func (x *DSDomainInfo) Reset() {
	*x = DSDomainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DSDomainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DSDomainInfo) ProtoMessage() {}

func (x *DSDomainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DSDomainInfo.ProtoReflect.Descriptor instead.
func (*DSDomainInfo) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{6}
}

func (x *DSDomainInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//业务域信息
type BizDomainInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` //业务域名称
}

func (x *BizDomainInfo) Reset() {
	*x = BizDomainInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BizDomainInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BizDomainInfo) ProtoMessage() {}

func (x *BizDomainInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BizDomainInfo.ProtoReflect.Descriptor instead.
func (*BizDomainInfo) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{7}
}

func (x *BizDomainInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

//业务系统信息
type BizSystemInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SysAddress string `protobuf:"bytes,1,opt,name=sysAddress,proto3" json:"sysAddress,omitempty"` //业务系统地址
	SysName    string `protobuf:"bytes,2,opt,name=sysName,proto3" json:"sysName,omitempty"`       //业务系统名称
}

func (x *BizSystemInfo) Reset() {
	*x = BizSystemInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contract_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BizSystemInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BizSystemInfo) ProtoMessage() {}

func (x *BizSystemInfo) ProtoReflect() protoreflect.Message {
	mi := &file_contract_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BizSystemInfo.ProtoReflect.Descriptor instead.
func (*BizSystemInfo) Descriptor() ([]byte, []int) {
	return file_contract_proto_rawDescGZIP(), []int{8}
}

func (x *BizSystemInfo) GetSysAddress() string {
	if x != nil {
		return x.SysAddress
	}
	return ""
}

func (x *BizSystemInfo) GetSysName() string {
	if x != nil {
		return x.SysName
	}
	return ""
}

var File_contract_proto protoreflect.FileDescriptor

var file_contract_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x20, 0x0a, 0x0a, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x35, 0x0a, 0x0b, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x26, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x7a, 0x0a, 0x08, 0x4e, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x69, 0x73, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x12,
	0x1c, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x22, 0x43, 0x0a,
	0x07, 0x4f, 0x72, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x67, 0x4e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6f, 0x72, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6f, 0x72, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x22, 0x3d, 0x0a, 0x0f, 0x42, 0x69, 0x7a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x22, 0xdc, 0x01, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x6f,
	0x64, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a,
	0x6f, 0x72, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x6f, 0x72, 0x67, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x53, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x53,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x70, 0x75, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x63, 0x70, 0x75, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x12, 0x12,
	0x0a, 0x04, 0x64, 0x69, 0x73, 0x6b, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x69,
	0x73, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62, 0x61, 0x6e, 0x64, 0x77, 0x69, 0x64, 0x74, 0x68,
	0x22, 0x22, 0x0a, 0x0c, 0x44, 0x53, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x23, 0x0a, 0x0d, 0x42, 0x69, 0x7a, 0x44, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x49, 0x0a, 0x0d, 0x42, 0x69, 0x7a,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x79,
	0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x73, 0x79, 0x73, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x79,
	0x73, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x79, 0x73,
	0x4e, 0x61, 0x6d, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contract_proto_rawDescOnce sync.Once
	file_contract_proto_rawDescData = file_contract_proto_rawDesc
)

func file_contract_proto_rawDescGZIP() []byte {
	file_contract_proto_rawDescOnce.Do(func() {
		file_contract_proto_rawDescData = protoimpl.X.CompressGZIP(file_contract_proto_rawDescData)
	})
	return file_contract_proto_rawDescData
}

var file_contract_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_contract_proto_goTypes = []interface{}{
	(*MemberInfo)(nil),       // 0: entity.MemberInfo
	(*MemberInfos)(nil),      // 1: entity.MemberInfos
	(*NodeInfo)(nil),         // 2: entity.NodeInfo
	(*OrgInfo)(nil),          // 3: entity.OrgInfo
	(*BizContractInfo)(nil),  // 4: entity.BizContractInfo
	(*ResourceNodeInfo)(nil), // 5: entity.ResourceNodeInfo
	(*DSDomainInfo)(nil),     // 6: entity.DSDomainInfo
	(*BizDomainInfo)(nil),    // 7: entity.BizDomainInfo
	(*BizSystemInfo)(nil),    // 8: entity.BizSystemInfo
}
var file_contract_proto_depIdxs = []int32{
	0, // 0: entity.MemberInfos.name:type_name -> entity.MemberInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contract_proto_init() }
func file_contract_proto_init() {
	if File_contract_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contract_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberInfo); i {
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
		file_contract_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberInfos); i {
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
		file_contract_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NodeInfo); i {
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
		file_contract_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrgInfo); i {
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
		file_contract_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BizContractInfo); i {
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
		file_contract_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResourceNodeInfo); i {
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
		file_contract_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DSDomainInfo); i {
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
		file_contract_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BizDomainInfo); i {
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
		file_contract_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BizSystemInfo); i {
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
			RawDescriptor: file_contract_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_contract_proto_goTypes,
		DependencyIndexes: file_contract_proto_depIdxs,
		MessageInfos:      file_contract_proto_msgTypes,
	}.Build()
	File_contract_proto = out.File
	file_contract_proto_rawDesc = nil
	file_contract_proto_goTypes = nil
	file_contract_proto_depIdxs = nil
}
