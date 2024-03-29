// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.4
// source: pb/office.proto

package rpc

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

type OfficeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	FirmID   string `protobuf:"bytes,3,opt,name=FirmID,proto3" json:"FirmID,omitempty"`
	FirmName string `protobuf:"bytes,4,opt,name=FirmName,proto3" json:"FirmName,omitempty"`
	BlocID   string `protobuf:"bytes,5,opt,name=BlocID,proto3" json:"BlocID,omitempty"`
	BlocName string `protobuf:"bytes,6,opt,name=BlocName,proto3" json:"BlocName,omitempty"`
	Province string `protobuf:"bytes,7,opt,name=Province,proto3" json:"Province,omitempty"`
	City     string `protobuf:"bytes,8,opt,name=City,proto3" json:"City,omitempty"`
}

func (x *OfficeResponse) Reset() {
	*x = OfficeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_office_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeResponse) ProtoMessage() {}

func (x *OfficeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_office_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeResponse.ProtoReflect.Descriptor instead.
func (*OfficeResponse) Descriptor() ([]byte, []int) {
	return file_pb_office_proto_rawDescGZIP(), []int{0}
}

func (x *OfficeResponse) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *OfficeResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfficeResponse) GetFirmID() string {
	if x != nil {
		return x.FirmID
	}
	return ""
}

func (x *OfficeResponse) GetFirmName() string {
	if x != nil {
		return x.FirmName
	}
	return ""
}

func (x *OfficeResponse) GetBlocID() string {
	if x != nil {
		return x.BlocID
	}
	return ""
}

func (x *OfficeResponse) GetBlocName() string {
	if x != nil {
		return x.BlocName
	}
	return ""
}

func (x *OfficeResponse) GetProvince() string {
	if x != nil {
		return x.Province
	}
	return ""
}

func (x *OfficeResponse) GetCity() string {
	if x != nil {
		return x.City
	}
	return ""
}

type OfficeResponseList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offices []*OfficeResponse `protobuf:"bytes,1,rep,name=Offices,proto3" json:"Offices,omitempty"`
}

func (x *OfficeResponseList) Reset() {
	*x = OfficeResponseList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_office_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeResponseList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeResponseList) ProtoMessage() {}

func (x *OfficeResponseList) ProtoReflect() protoreflect.Message {
	mi := &file_pb_office_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeResponseList.ProtoReflect.Descriptor instead.
func (*OfficeResponseList) Descriptor() ([]byte, []int) {
	return file_pb_office_proto_rawDescGZIP(), []int{1}
}

func (x *OfficeResponseList) GetOffices() []*OfficeResponse {
	if x != nil {
		return x.Offices
	}
	return nil
}

type OfficeBaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name         string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`                  //名称
	Logo         string   `protobuf:"bytes,2,opt,name=Logo,proto3" json:"Logo,omitempty"`                  //机构logo
	Master       string   `protobuf:"bytes,3,opt,name=Master,proto3" json:"Master,omitempty"`              //负责人
	Phone        string   `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`                //电话
	Email        string   `protobuf:"bytes,5,opt,name=Email,proto3" json:"Email,omitempty"`                //邮箱
	ProvinceID   string   `protobuf:"bytes,6,opt,name=ProvinceID,proto3" json:"ProvinceID,omitempty"`      //所属省id
	CityID       string   `protobuf:"bytes,7,opt,name=CityID,proto3" json:"CityID,omitempty"`              //所属市id
	Address      string   `protobuf:"bytes,8,opt,name=Address,proto3" json:"Address,omitempty"`            //地址
	Remark       string   `protobuf:"bytes,9,opt,name=Remark,proto3" json:"Remark,omitempty"`              //描述信息
	Applications []string `protobuf:"bytes,10,rep,name=Applications,proto3" json:"Applications,omitempty"` //应用IDs
	ID           string   `protobuf:"bytes,11,opt,name=ID,proto3" json:"ID,omitempty"`                     //id
	BlockID      string   `protobuf:"bytes,12,opt,name=BlockID,proto3" json:"BlockID,omitempty"`           // 集团ID
}

func (x *OfficeBaseRequest) Reset() {
	*x = OfficeBaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_office_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeBaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeBaseRequest) ProtoMessage() {}

func (x *OfficeBaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_office_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeBaseRequest.ProtoReflect.Descriptor instead.
func (*OfficeBaseRequest) Descriptor() ([]byte, []int) {
	return file_pb_office_proto_rawDescGZIP(), []int{2}
}

func (x *OfficeBaseRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OfficeBaseRequest) GetLogo() string {
	if x != nil {
		return x.Logo
	}
	return ""
}

func (x *OfficeBaseRequest) GetMaster() string {
	if x != nil {
		return x.Master
	}
	return ""
}

func (x *OfficeBaseRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *OfficeBaseRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *OfficeBaseRequest) GetProvinceID() string {
	if x != nil {
		return x.ProvinceID
	}
	return ""
}

func (x *OfficeBaseRequest) GetCityID() string {
	if x != nil {
		return x.CityID
	}
	return ""
}

func (x *OfficeBaseRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *OfficeBaseRequest) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *OfficeBaseRequest) GetApplications() []string {
	if x != nil {
		return x.Applications
	}
	return nil
}

func (x *OfficeBaseRequest) GetID() string {
	if x != nil {
		return x.ID
	}
	return ""
}

func (x *OfficeBaseRequest) GetBlockID() string {
	if x != nil {
		return x.BlockID
	}
	return ""
}

type OfficeBaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows int64 `protobuf:"varint,1,opt,name=Rows,proto3" json:"Rows,omitempty"`
}

func (x *OfficeBaseResponse) Reset() {
	*x = OfficeBaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_office_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OfficeBaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OfficeBaseResponse) ProtoMessage() {}

func (x *OfficeBaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pb_office_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OfficeBaseResponse.ProtoReflect.Descriptor instead.
func (*OfficeBaseResponse) Descriptor() ([]byte, []int) {
	return file_pb_office_proto_rawDescGZIP(), []int{3}
}

func (x *OfficeBaseResponse) GetRows() int64 {
	if x != nil {
		return x.Rows
	}
	return 0
}

var File_pb_office_proto protoreflect.FileDescriptor

var file_pb_office_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x62, 0x2f, 0x6f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x0f, 0x70, 0x62, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x69, 0x72,
	0x6d, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x46, 0x69, 0x72, 0x6d, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x72, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x46, 0x69, 0x72, 0x6d, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x42, 0x6c, 0x6f, 0x63, 0x49, 0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x42,
	0x6c, 0x6f, 0x63, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x6c, 0x6f, 0x63, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x43, 0x69, 0x74, 0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x43, 0x69, 0x74,
	0x79, 0x22, 0x3f, 0x0a, 0x12, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29, 0x0a, 0x07, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x07, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x73, 0x22, 0xb7, 0x02, 0x0a, 0x11, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x42, 0x61, 0x73,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x4c, 0x6f, 0x67, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4c, 0x6f, 0x67, 0x6f,
	0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x6e,
	0x63, 0x65, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x69, 0x74, 0x79, 0x49, 0x44, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x69, 0x74, 0x79, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b,
	0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x52, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x22,
	0x0a, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x0c, 0x41, 0x70, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x44, 0x22, 0x28, 0x0a, 0x12,
	0x4f, 0x66, 0x66, 0x69, 0x63, 0x65, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x52, 0x6f, 0x77, 0x73, 0x32, 0xc7, 0x01, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x12, 0x26, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x44, 0x12, 0x0a, 0x2e, 0x49,
	0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0d, 0x46, 0x69, 0x6e,
	0x64, 0x42, 0x79, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x0f, 0x2e, 0x43, 0x75, 0x72,
	0x72, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x2d, 0x0a, 0x09, 0x46, 0x69, 0x6e, 0x64, 0x49, 0x6e, 0x49, 0x44, 0x73, 0x12, 0x0b, 0x2e,
	0x49, 0x44, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4f, 0x66, 0x66,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x2f, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x12, 0x2e, 0x4f, 0x66, 0x66, 0x69, 0x63, 0x65,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x4f, 0x66,
	0x66, 0x69, 0x63, 0x65, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x06, 0x5a, 0x04, 0x2f, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_office_proto_rawDescOnce sync.Once
	file_pb_office_proto_rawDescData = file_pb_office_proto_rawDesc
)

func file_pb_office_proto_rawDescGZIP() []byte {
	file_pb_office_proto_rawDescOnce.Do(func() {
		file_pb_office_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_office_proto_rawDescData)
	})
	return file_pb_office_proto_rawDescData
}

var file_pb_office_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pb_office_proto_goTypes = []interface{}{
	(*OfficeResponse)(nil),     // 0: OfficeResponse
	(*OfficeResponseList)(nil), // 1: OfficeResponseList
	(*OfficeBaseRequest)(nil),  // 2: OfficeBaseRequest
	(*OfficeBaseResponse)(nil), // 3: OfficeBaseResponse
	(*IDRequest)(nil),          // 4: IDRequest
	(*CurrentRequest)(nil),     // 5: CurrentRequest
	(*IDsRequest)(nil),         // 6: IDsRequest
}
var file_pb_office_proto_depIdxs = []int32{
	0, // 0: OfficeResponseList.Offices:type_name -> OfficeResponse
	4, // 1: Office.GetByID:input_type -> IDRequest
	5, // 2: Office.FindByCurrent:input_type -> CurrentRequest
	6, // 3: Office.FindInIDs:input_type -> IDsRequest
	2, // 4: Office.Save:input_type -> OfficeBaseRequest
	0, // 5: Office.GetByID:output_type -> OfficeResponse
	1, // 6: Office.FindByCurrent:output_type -> OfficeResponseList
	1, // 7: Office.FindInIDs:output_type -> OfficeResponseList
	3, // 8: Office.Save:output_type -> OfficeBaseResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pb_office_proto_init() }
func file_pb_office_proto_init() {
	if File_pb_office_proto != nil {
		return
	}
	file_pb_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_pb_office_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeResponse); i {
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
		file_pb_office_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeResponseList); i {
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
		file_pb_office_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeBaseRequest); i {
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
		file_pb_office_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OfficeBaseResponse); i {
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
			RawDescriptor: file_pb_office_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_office_proto_goTypes,
		DependencyIndexes: file_pb_office_proto_depIdxs,
		MessageInfos:      file_pb_office_proto_msgTypes,
	}.Build()
	File_pb_office_proto = out.File
	file_pb_office_proto_rawDesc = nil
	file_pb_office_proto_goTypes = nil
	file_pb_office_proto_depIdxs = nil
}
