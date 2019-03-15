// Code generated by protoc-gen-go. DO NOT EDIT.
// source: profile.proto

package profile

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GetInfoRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoRequest) Reset()         { *m = GetInfoRequest{} }
func (m *GetInfoRequest) String() string { return proto.CompactTextString(m) }
func (*GetInfoRequest) ProtoMessage()    {}
func (*GetInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{0}
}

func (m *GetInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoRequest.Unmarshal(m, b)
}
func (m *GetInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoRequest.Marshal(b, m, deterministic)
}
func (m *GetInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoRequest.Merge(m, src)
}
func (m *GetInfoRequest) XXX_Size() int {
	return xxx_messageInfo_GetInfoRequest.Size(m)
}
func (m *GetInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoRequest proto.InternalMessageInfo

func (m *GetInfoRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

type GetInfoResponse struct {
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Birth                int64    `protobuf:"varint,4,opt,name=Birth,proto3" json:"Birth,omitempty"`
	Status               string   `protobuf:"bytes,5,opt,name=Status,proto3" json:"Status,omitempty"`
	Email                string   `protobuf:"bytes,6,opt,name=Email,proto3" json:"Email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetInfoResponse) Reset()         { *m = GetInfoResponse{} }
func (m *GetInfoResponse) String() string { return proto.CompactTextString(m) }
func (*GetInfoResponse) ProtoMessage()    {}
func (*GetInfoResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{1}
}

func (m *GetInfoResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetInfoResponse.Unmarshal(m, b)
}
func (m *GetInfoResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetInfoResponse.Marshal(b, m, deterministic)
}
func (m *GetInfoResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetInfoResponse.Merge(m, src)
}
func (m *GetInfoResponse) XXX_Size() int {
	return xxx_messageInfo_GetInfoResponse.Size(m)
}
func (m *GetInfoResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetInfoResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetInfoResponse proto.InternalMessageInfo

func (m *GetInfoResponse) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *GetInfoResponse) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *GetInfoResponse) GetBirth() int64 {
	if m != nil {
		return m.Birth
	}
	return 0
}

func (m *GetInfoResponse) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *GetInfoResponse) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type UpdateStatusRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	NewStatus            string   `protobuf:"bytes,2,opt,name=newStatus,proto3" json:"newStatus,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateStatusRequest) Reset()         { *m = UpdateStatusRequest{} }
func (m *UpdateStatusRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateStatusRequest) ProtoMessage()    {}
func (*UpdateStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{2}
}

func (m *UpdateStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateStatusRequest.Unmarshal(m, b)
}
func (m *UpdateStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateStatusRequest.Marshal(b, m, deterministic)
}
func (m *UpdateStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateStatusRequest.Merge(m, src)
}
func (m *UpdateStatusRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateStatusRequest.Size(m)
}
func (m *UpdateStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateStatusRequest proto.InternalMessageInfo

func (m *UpdateStatusRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateStatusRequest) GetNewStatus() string {
	if m != nil {
		return m.NewStatus
	}
	return ""
}

type UpdateEmailRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	NewEmail             string   `protobuf:"bytes,2,opt,name=newEmail,proto3" json:"newEmail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateEmailRequest) Reset()         { *m = UpdateEmailRequest{} }
func (m *UpdateEmailRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateEmailRequest) ProtoMessage()    {}
func (*UpdateEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{3}
}

func (m *UpdateEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateEmailRequest.Unmarshal(m, b)
}
func (m *UpdateEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateEmailRequest.Marshal(b, m, deterministic)
}
func (m *UpdateEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateEmailRequest.Merge(m, src)
}
func (m *UpdateEmailRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateEmailRequest.Size(m)
}
func (m *UpdateEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateEmailRequest proto.InternalMessageInfo

func (m *UpdateEmailRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateEmailRequest) GetNewEmail() string {
	if m != nil {
		return m.NewEmail
	}
	return ""
}

type ConfirmEmailRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	EmailToken           string   `protobuf:"bytes,2,opt,name=emailToken,proto3" json:"emailToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfirmEmailRequest) Reset()         { *m = ConfirmEmailRequest{} }
func (m *ConfirmEmailRequest) String() string { return proto.CompactTextString(m) }
func (*ConfirmEmailRequest) ProtoMessage()    {}
func (*ConfirmEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{4}
}

func (m *ConfirmEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfirmEmailRequest.Unmarshal(m, b)
}
func (m *ConfirmEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfirmEmailRequest.Marshal(b, m, deterministic)
}
func (m *ConfirmEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfirmEmailRequest.Merge(m, src)
}
func (m *ConfirmEmailRequest) XXX_Size() int {
	return xxx_messageInfo_ConfirmEmailRequest.Size(m)
}
func (m *ConfirmEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfirmEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfirmEmailRequest proto.InternalMessageInfo

func (m *ConfirmEmailRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ConfirmEmailRequest) GetEmailToken() string {
	if m != nil {
		return m.EmailToken
	}
	return ""
}

type UpdateInfoRequest struct {
	UserID               string   `protobuf:"bytes,1,opt,name=userID,proto3" json:"userID,omitempty"`
	FirstName            string   `protobuf:"bytes,2,opt,name=FirstName,proto3" json:"FirstName,omitempty"`
	LastName             string   `protobuf:"bytes,3,opt,name=LastName,proto3" json:"LastName,omitempty"`
	Birth                int64    `protobuf:"varint,4,opt,name=Birth,proto3" json:"Birth,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateInfoRequest) Reset()         { *m = UpdateInfoRequest{} }
func (m *UpdateInfoRequest) String() string { return proto.CompactTextString(m) }
func (*UpdateInfoRequest) ProtoMessage()    {}
func (*UpdateInfoRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{5}
}

func (m *UpdateInfoRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateInfoRequest.Unmarshal(m, b)
}
func (m *UpdateInfoRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateInfoRequest.Marshal(b, m, deterministic)
}
func (m *UpdateInfoRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateInfoRequest.Merge(m, src)
}
func (m *UpdateInfoRequest) XXX_Size() int {
	return xxx_messageInfo_UpdateInfoRequest.Size(m)
}
func (m *UpdateInfoRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateInfoRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateInfoRequest proto.InternalMessageInfo

func (m *UpdateInfoRequest) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *UpdateInfoRequest) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *UpdateInfoRequest) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

func (m *UpdateInfoRequest) GetBirth() int64 {
	if m != nil {
		return m.Birth
	}
	return 0
}

type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_744bf7a47b381504, []int{6}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func init() {
	proto.RegisterType((*GetInfoRequest)(nil), "GetInfoRequest")
	proto.RegisterType((*GetInfoResponse)(nil), "GetInfoResponse")
	proto.RegisterType((*UpdateStatusRequest)(nil), "UpdateStatusRequest")
	proto.RegisterType((*UpdateEmailRequest)(nil), "UpdateEmailRequest")
	proto.RegisterType((*ConfirmEmailRequest)(nil), "ConfirmEmailRequest")
	proto.RegisterType((*UpdateInfoRequest)(nil), "UpdateInfoRequest")
	proto.RegisterType((*Response)(nil), "Response")
}

func init() { proto.RegisterFile("profile.proto", fileDescriptor_744bf7a47b381504) }

var fileDescriptor_744bf7a47b381504 = []byte{
	// 334 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x53, 0xcb, 0x4e, 0xeb, 0x30,
	0x14, 0x54, 0x6e, 0x6f, 0x5f, 0x87, 0x47, 0xe1, 0xb4, 0x42, 0x51, 0x84, 0x50, 0xe5, 0x55, 0x25,
	0xc0, 0x48, 0xf0, 0x07, 0xbc, 0x2b, 0x1e, 0x8b, 0x02, 0x1f, 0x10, 0xc4, 0xa9, 0x88, 0x68, 0xed,
	0x60, 0x3b, 0xea, 0x82, 0x8f, 0xe0, 0x6b, 0xd9, 0xa3, 0xc4, 0x26, 0x35, 0x6d, 0xa5, 0x76, 0xc1,
	0x72, 0x8e, 0xc7, 0xe3, 0xc9, 0x99, 0x09, 0x6c, 0xa4, 0x4a, 0x0e, 0x93, 0x11, 0xf1, 0x54, 0x49,
	0x23, 0x59, 0x0f, 0x36, 0xaf, 0xc8, 0xf4, 0xc5, 0x50, 0x0e, 0xe8, 0x3d, 0x23, 0x6d, 0x70, 0x07,
	0x6a, 0x99, 0x26, 0xd5, 0x3f, 0x0f, 0x83, 0x6e, 0xd0, 0x6b, 0x0e, 0x1c, 0x62, 0x9f, 0x01, 0xb4,
	0x4a, 0xaa, 0x4e, 0xa5, 0xd0, 0x84, 0xbb, 0xd0, 0xbc, 0x4c, 0x94, 0x36, 0xf7, 0xf1, 0x98, 0xc2,
	0x7f, 0x05, 0x7d, 0x3a, 0xc0, 0x08, 0x1a, 0xb7, 0xb1, 0x3b, 0xac, 0x14, 0x87, 0x25, 0xc6, 0x0e,
	0x54, 0x4f, 0x13, 0x65, 0x5e, 0xc3, 0xff, 0xdd, 0xa0, 0x57, 0x19, 0x58, 0x90, 0xbf, 0xfd, 0x60,
	0x62, 0x93, 0xe9, 0xb0, 0x6a, 0xdf, 0xb6, 0x28, 0x67, 0x5f, 0x8c, 0xe3, 0x64, 0x14, 0xd6, 0x8a,
	0xb1, 0x05, 0xec, 0x06, 0xda, 0x4f, 0xe9, 0x4b, 0x6c, 0xc8, 0xb2, 0x96, 0x7c, 0x40, 0x6e, 0x56,
	0xd0, 0xc4, 0xe9, 0x3b, 0xb3, 0xe5, 0x80, 0x5d, 0x03, 0x5a, 0xb1, 0x42, 0x7b, 0x99, 0x56, 0x04,
	0x0d, 0x41, 0x13, 0xeb, 0xc9, 0x4a, 0x95, 0x98, 0xdd, 0x41, 0xfb, 0x4c, 0x8a, 0x61, 0xa2, 0xc6,
	0x2b, 0x49, 0xed, 0x01, 0x50, 0xce, 0x7b, 0x94, 0x6f, 0x24, 0x9c, 0x98, 0x37, 0x61, 0x1f, 0xb0,
	0x6d, 0x8d, 0xad, 0x10, 0xd2, 0x5f, 0x07, 0xc2, 0x00, 0x1a, 0x3f, 0x61, 0x1f, 0x7f, 0x05, 0x50,
	0x77, 0xe5, 0xc1, 0x7d, 0x80, 0xa9, 0x29, 0x44, 0x3e, 0xe7, 0x30, 0x6a, 0xf2, 0xb2, 0x25, 0x47,
	0xb0, 0xee, 0xe7, 0x84, 0x1d, 0xbe, 0x20, 0x36, 0xff, 0xc2, 0x21, 0xac, 0x79, 0x59, 0x60, 0x9b,
	0xcf, 0x27, 0x33, 0xa3, 0xef, 0x2f, 0x1c, 0x3b, 0x7c, 0xc1, 0xfe, 0xfd, 0x0b, 0x07, 0x50, 0x77,
	0x4d, 0xc6, 0x16, 0xff, 0x5d, 0xff, 0x68, 0x8b, 0xcf, 0x94, 0xfc, 0xb9, 0x56, 0xfc, 0x29, 0x27,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xdc, 0xe3, 0xaa, 0x99, 0x3a, 0x03, 0x00, 0x00,
}
