// Code generated by protoc-gen-gogo.
// source: stn.proto
// DO NOT EDIT!

/*
Package stn is a generated protocol buffer package.

It is generated from these files:
	stn.proto

It has these top-level messages:
	StnRule
*/
package stn

import proto "github.com/gogo/protobuf/proto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal

type StnRule struct {
	IpAddress string `protobuf:"bytes,1,opt,name=ip_address,proto3" json:"ip_address,omitempty"`
	Interface string `protobuf:"bytes,2,opt,name=interface,proto3" json:"interface,omitempty"`
	RuleName  string `protobuf:"bytes,3,opt,name=rule_name,proto3" json:"rule_name,omitempty"`
}

func (m *StnRule) Reset()         { *m = StnRule{} }
func (m *StnRule) String() string { return proto.CompactTextString(m) }
func (*StnRule) ProtoMessage()    {}