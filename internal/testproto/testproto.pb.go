// Code generated by protoc-gen-go. DO NOT EDIT.
// source: testproto.proto

/*
Package testproto is a generated protocol buffer package.

It is generated from these files:
	testproto.proto

It has these top-level messages:
	OAuthState
	Session
*/
package testproto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// OAuthState contains data associated with a single oauth flow (currently just the url to redirect the user to after
// authentication completes)
type OAuthState struct {
	RedirectUrl string `protobuf:"bytes,1,opt,name=redirect_url,json=redirectUrl" json:"redirect_url,omitempty"`
}

func (m *OAuthState) Reset()                    { *m = OAuthState{} }
func (m *OAuthState) String() string            { return proto.CompactTextString(m) }
func (*OAuthState) ProtoMessage()               {}
func (*OAuthState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *OAuthState) GetRedirectUrl() string {
	if m != nil {
		return m.RedirectUrl
	}
	return ""
}

// Session contains data associated with a single user: who that user is and whether they're authenticated & authorized
type Session struct {
	User       string                     `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	ExpiresAt  *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=expires_at,json=expiresAt" json:"expires_at,omitempty"`
	Authorized bool                       `protobuf:"varint,3,opt,name=authorized" json:"authorized,omitempty"`
}

func (m *Session) Reset()                    { *m = Session{} }
func (m *Session) String() string            { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()               {}
func (*Session) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Session) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *Session) GetExpiresAt() *google_protobuf.Timestamp {
	if m != nil {
		return m.ExpiresAt
	}
	return nil
}

func (m *Session) GetAuthorized() bool {
	if m != nil {
		return m.Authorized
	}
	return false
}

func init() {
	proto.RegisterType((*OAuthState)(nil), "testproto.OAuthState")
	proto.RegisterType((*Session)(nil), "testproto.Session")
}

func init() { proto.RegisterFile("testproto.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 195 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8e, 0x41, 0x4b, 0x85, 0x40,
	0x10, 0x80, 0xd9, 0x8a, 0xca, 0x31, 0x08, 0xf6, 0x24, 0x1e, 0xca, 0x3c, 0x79, 0x5a, 0xa1, 0x4e,
	0x1d, 0xfd, 0x05, 0x81, 0xd6, 0x59, 0xd6, 0x9c, 0x74, 0x41, 0x5b, 0x99, 0x9d, 0x05, 0x79, 0xbf,
	0xfe, 0xc1, 0xaa, 0x8f, 0x77, 0x19, 0x66, 0x3e, 0xbe, 0x81, 0x0f, 0x9e, 0x19, 0x1d, 0x2f, 0x64,
	0xd9, 0xaa, 0x30, 0x65, 0x74, 0x01, 0xe9, 0xeb, 0x60, 0xed, 0x30, 0x61, 0x19, 0xae, 0xce, 0xff,
	0x95, 0x6c, 0x66, 0x74, 0xac, 0xe7, 0x65, 0x73, 0xf3, 0x12, 0xe0, 0xab, 0xf2, 0x3c, 0x36, 0xac,
	0x19, 0xe5, 0x1b, 0x3c, 0x11, 0xf6, 0x86, 0xf0, 0x97, 0x5b, 0x4f, 0x53, 0x22, 0x32, 0x51, 0x44,
	0x75, 0x7c, 0xb0, 0x1f, 0x9a, 0xf2, 0x15, 0x1e, 0x1a, 0x74, 0xce, 0xd8, 0x7f, 0x29, 0xe1, 0xce,
	0x3b, 0xa4, 0xdd, 0x0a, 0xbb, 0xfc, 0x04, 0xc0, 0x75, 0x31, 0x84, 0xae, 0xd5, 0x9c, 0xdc, 0x64,
	0xa2, 0x88, 0xdf, 0x53, 0xb5, 0x55, 0xa8, 0xa3, 0x42, 0x7d, 0x1f, 0x15, 0x75, 0xb4, 0xdb, 0x15,
	0xcb, 0x17, 0x00, 0xed, 0x79, 0xb4, 0x64, 0x4e, 0xd8, 0x27, 0xb7, 0x99, 0x28, 0x1e, 0xeb, 0x2b,
	0xd2, 0xdd, 0x87, 0xf7, 0x8f, 0x73, 0x00, 0x00, 0x00, 0xff, 0xff, 0x31, 0x51, 0x31, 0x72, 0xf0,
	0x00, 0x00, 0x00,
}
