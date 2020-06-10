/*
Copyright 2020 The KubeCarrier Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by protoc-gen-go. DO NOT EDIT.
// source: types.proto

package v1

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
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

type ObjectReference struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ObjectReference) Reset()         { *m = ObjectReference{} }
func (m *ObjectReference) String() string { return proto.CompactTextString(m) }
func (*ObjectReference) ProtoMessage()    {}
func (*ObjectReference) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{0}
}

func (m *ObjectReference) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ObjectReference.Unmarshal(m, b)
}
func (m *ObjectReference) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ObjectReference.Marshal(b, m, deterministic)
}
func (m *ObjectReference) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ObjectReference.Merge(m, src)
}
func (m *ObjectReference) XXX_Size() int {
	return xxx_messageInfo_ObjectReference.Size(m)
}
func (m *ObjectReference) XXX_DiscardUnknown() {
	xxx_messageInfo_ObjectReference.DiscardUnknown(m)
}

var xxx_messageInfo_ObjectReference proto.InternalMessageInfo

func (m *ObjectReference) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type CRDInformation struct {
	Name                 string           `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ApiGroup             string           `protobuf:"bytes,2,opt,name=apiGroup,proto3" json:"apiGroup,omitempty"`
	Kind                 string           `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	Plural               string           `protobuf:"bytes,4,opt,name=plural,proto3" json:"plural,omitempty"`
	Versions             []*CRDVersion    `protobuf:"bytes,5,rep,name=versions,proto3" json:"versions,omitempty"`
	Region               *ObjectReference `protobuf:"bytes,6,opt,name=region,proto3" json:"region,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *CRDInformation) Reset()         { *m = CRDInformation{} }
func (m *CRDInformation) String() string { return proto.CompactTextString(m) }
func (*CRDInformation) ProtoMessage()    {}
func (*CRDInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{1}
}

func (m *CRDInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CRDInformation.Unmarshal(m, b)
}
func (m *CRDInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CRDInformation.Marshal(b, m, deterministic)
}
func (m *CRDInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CRDInformation.Merge(m, src)
}
func (m *CRDInformation) XXX_Size() int {
	return xxx_messageInfo_CRDInformation.Size(m)
}
func (m *CRDInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_CRDInformation.DiscardUnknown(m)
}

var xxx_messageInfo_CRDInformation proto.InternalMessageInfo

func (m *CRDInformation) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CRDInformation) GetApiGroup() string {
	if m != nil {
		return m.ApiGroup
	}
	return ""
}

func (m *CRDInformation) GetKind() string {
	if m != nil {
		return m.Kind
	}
	return ""
}

func (m *CRDInformation) GetPlural() string {
	if m != nil {
		return m.Plural
	}
	return ""
}

func (m *CRDInformation) GetVersions() []*CRDVersion {
	if m != nil {
		return m.Versions
	}
	return nil
}

func (m *CRDInformation) GetRegion() *ObjectReference {
	if m != nil {
		return m.Region
	}
	return nil
}

type CRDVersion struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Schema               string   `protobuf:"bytes,2,opt,name=schema,proto3" json:"schema,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CRDVersion) Reset()         { *m = CRDVersion{} }
func (m *CRDVersion) String() string { return proto.CompactTextString(m) }
func (*CRDVersion) ProtoMessage()    {}
func (*CRDVersion) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{2}
}

func (m *CRDVersion) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CRDVersion.Unmarshal(m, b)
}
func (m *CRDVersion) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CRDVersion.Marshal(b, m, deterministic)
}
func (m *CRDVersion) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CRDVersion.Merge(m, src)
}
func (m *CRDVersion) XXX_Size() int {
	return xxx_messageInfo_CRDVersion.Size(m)
}
func (m *CRDVersion) XXX_DiscardUnknown() {
	xxx_messageInfo_CRDVersion.DiscardUnknown(m)
}

var xxx_messageInfo_CRDVersion proto.InternalMessageInfo

func (m *CRDVersion) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CRDVersion) GetSchema() string {
	if m != nil {
		return m.Schema
	}
	return ""
}

type Image struct {
	MediaType            string   `protobuf:"bytes,1,opt,name=mediaType,proto3" json:"mediaType,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{3}
}

func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (m *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(m, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetMediaType() string {
	if m != nil {
		return m.MediaType
	}
	return ""
}

func (m *Image) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type ConditionStatus struct {
	Status               string   `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConditionStatus) Reset()         { *m = ConditionStatus{} }
func (m *ConditionStatus) String() string { return proto.CompactTextString(m) }
func (*ConditionStatus) ProtoMessage()    {}
func (*ConditionStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_d938547f84707355, []int{4}
}

func (m *ConditionStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConditionStatus.Unmarshal(m, b)
}
func (m *ConditionStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConditionStatus.Marshal(b, m, deterministic)
}
func (m *ConditionStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConditionStatus.Merge(m, src)
}
func (m *ConditionStatus) XXX_Size() int {
	return xxx_messageInfo_ConditionStatus.Size(m)
}
func (m *ConditionStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_ConditionStatus.DiscardUnknown(m)
}

var xxx_messageInfo_ConditionStatus proto.InternalMessageInfo

func (m *ConditionStatus) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func init() {
	proto.RegisterType((*ObjectReference)(nil), "kubecarrier.api.v1.ObjectReference")
	proto.RegisterType((*CRDInformation)(nil), "kubecarrier.api.v1.CRDInformation")
	proto.RegisterType((*CRDVersion)(nil), "kubecarrier.api.v1.CRDVersion")
	proto.RegisterType((*Image)(nil), "kubecarrier.api.v1.Image")
	proto.RegisterType((*ConditionStatus)(nil), "kubecarrier.api.v1.ConditionStatus")
}

func init() {
	proto.RegisterFile("types.proto", fileDescriptor_d938547f84707355)
}

var fileDescriptor_d938547f84707355 = []byte{
	// 298 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x51, 0x4d, 0x4b, 0x33, 0x31,
	0x18, 0x64, 0xfb, 0xb1, 0xb4, 0x4f, 0x5f, 0xde, 0x42, 0x0e, 0x25, 0x88, 0x48, 0x59, 0x11, 0xea,
	0x65, 0xa1, 0xf5, 0xe2, 0xc7, 0xcd, 0x16, 0xa4, 0x27, 0x21, 0x8a, 0x07, 0x6f, 0xd9, 0xdd, 0xa7,
	0x35, 0xb6, 0x9b, 0x84, 0x6c, 0x76, 0xa1, 0x3f, 0xd8, 0xff, 0x21, 0x9b, 0x8d, 0x2d, 0xe8, 0xde,
	0x66, 0x86, 0x99, 0x49, 0x26, 0x81, 0x91, 0x3d, 0x68, 0x2c, 0x62, 0x6d, 0x94, 0x55, 0x84, 0xec,
	0xca, 0x04, 0x53, 0x6e, 0x8c, 0x40, 0x13, 0x73, 0x2d, 0xe2, 0x6a, 0x1e, 0x5d, 0xc1, 0xf8, 0x39,
	0xf9, 0xc4, 0xd4, 0x32, 0xdc, 0xa0, 0x41, 0x99, 0x22, 0x21, 0xd0, 0x93, 0x3c, 0x47, 0x1a, 0x4c,
	0x83, 0xd9, 0x90, 0x39, 0x1c, 0x7d, 0x05, 0xf0, 0x7f, 0xc9, 0x56, 0x6b, 0xb9, 0x51, 0x26, 0xe7,
	0x56, 0x28, 0xd9, 0x66, 0x23, 0x67, 0x30, 0xe0, 0x5a, 0x3c, 0x19, 0x55, 0x6a, 0xda, 0x71, 0xfa,
	0x91, 0xd7, 0xfe, 0x9d, 0x90, 0x19, 0xed, 0x36, 0xfe, 0x1a, 0x93, 0x09, 0x84, 0x7a, 0x5f, 0x1a,
	0xbe, 0xa7, 0x3d, 0xa7, 0x7a, 0x46, 0xee, 0x61, 0x50, 0xa1, 0x29, 0x84, 0x92, 0x05, 0xed, 0x4f,
	0xbb, 0xb3, 0xd1, 0xe2, 0x22, 0xfe, 0x7b, 0xf9, 0x78, 0xc9, 0x56, 0x6f, 0x8d, 0x8d, 0x1d, 0xfd,
	0xe4, 0x01, 0x42, 0x83, 0x5b, 0xa1, 0x24, 0x0d, 0xa7, 0xc1, 0x6c, 0xb4, 0xb8, 0x6c, 0x4b, 0xfe,
	0xda, 0xcc, 0x7c, 0x24, 0xba, 0x05, 0x38, 0x95, 0xb6, 0x4e, 0x9c, 0x40, 0x58, 0xa4, 0x1f, 0x98,
	0x73, 0x3f, 0xd0, 0xb3, 0xe8, 0x0e, 0xfa, 0xeb, 0x9c, 0x6f, 0x91, 0x9c, 0xc3, 0x30, 0xc7, 0x4c,
	0xf0, 0xd7, 0x83, 0xfe, 0x49, 0x9e, 0x84, 0xba, 0x32, 0xe3, 0xb6, 0x09, 0xff, 0x63, 0x0e, 0x47,
	0xd7, 0x30, 0x5e, 0x2a, 0x99, 0x89, 0xfa, 0x59, 0x5f, 0x2c, 0xb7, 0x65, 0xe1, 0x4e, 0x71, 0xc8,
	0x37, 0x78, 0xf6, 0xd8, 0x7b, 0xef, 0x54, 0xf3, 0x24, 0x74, 0xff, 0x79, 0xf3, 0x1d, 0x00, 0x00,
	0xff, 0xff, 0x78, 0xdc, 0x03, 0x7e, 0xde, 0x01, 0x00, 0x00,
}