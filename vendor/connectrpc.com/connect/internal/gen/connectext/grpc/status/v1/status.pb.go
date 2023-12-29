// Copyright 2021-2023 The Connect Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: connectext/grpc/status/v1/status.proto

// This package is for internal use by Connect, and provides no backward
// compatibility guarantees whatsoever.

package statusv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// See https://cloud.google.com/apis/design/errors.
//
// This struct must remain binary-compatible with
// https://github.com/googleapis/googleapis/blob/master/google/rpc/status.proto.
type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    int32        `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`      // a google.rpc.Code
	Message string       `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"` // developer-facing, English (localize in details or client-side)
	Details []*anypb.Any `protobuf:"bytes,3,rep,name=details,proto3" json:"details,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_connectext_grpc_status_v1_status_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_connectext_grpc_status_v1_status_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_connectext_grpc_status_v1_status_proto_rawDescGZIP(), []int{0}
}

func (x *Status) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Status) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Status) GetDetails() []*anypb.Any {
	if x != nil {
		return x.Details
	}
	return nil
}

var File_connectext_grpc_status_v1_status_proto protoreflect.FileDescriptor

var file_connectext_grpc_status_v1_status_proto_rawDesc = []byte{
	0x0a, 0x26, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x66, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2e, 0x0a, 0x07, 0x64,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41,
	0x6e, 0x79, 0x52, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x42, 0xc3, 0x01, 0x0a, 0x12,
	0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e,
	0x76, 0x31, 0x42, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x46, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x72, 0x70, 0x63, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x65, 0x78,
	0x74, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2f, 0x76, 0x31,
	0x3b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x53, 0x58, 0xaa,
	0x02, 0x0e, 0x47, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x0e, 0x47, 0x72, 0x70, 0x63, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5c, 0x56,
	0x31, 0xe2, 0x02, 0x1a, 0x47, 0x72, 0x70, 0x63, 0x5c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x10, 0x47, 0x72, 0x70, 0x63, 0x3a, 0x3a, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_connectext_grpc_status_v1_status_proto_rawDescOnce sync.Once
	file_connectext_grpc_status_v1_status_proto_rawDescData = file_connectext_grpc_status_v1_status_proto_rawDesc
)

func file_connectext_grpc_status_v1_status_proto_rawDescGZIP() []byte {
	file_connectext_grpc_status_v1_status_proto_rawDescOnce.Do(func() {
		file_connectext_grpc_status_v1_status_proto_rawDescData = protoimpl.X.CompressGZIP(file_connectext_grpc_status_v1_status_proto_rawDescData)
	})
	return file_connectext_grpc_status_v1_status_proto_rawDescData
}

var file_connectext_grpc_status_v1_status_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_connectext_grpc_status_v1_status_proto_goTypes = []interface{}{
	(*Status)(nil),    // 0: grpc.status.v1.Status
	(*anypb.Any)(nil), // 1: google.protobuf.Any
}
var file_connectext_grpc_status_v1_status_proto_depIdxs = []int32{
	1, // 0: grpc.status.v1.Status.details:type_name -> google.protobuf.Any
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_connectext_grpc_status_v1_status_proto_init() }
func file_connectext_grpc_status_v1_status_proto_init() {
	if File_connectext_grpc_status_v1_status_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_connectext_grpc_status_v1_status_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
			RawDescriptor: file_connectext_grpc_status_v1_status_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_connectext_grpc_status_v1_status_proto_goTypes,
		DependencyIndexes: file_connectext_grpc_status_v1_status_proto_depIdxs,
		MessageInfos:      file_connectext_grpc_status_v1_status_proto_msgTypes,
	}.Build()
	File_connectext_grpc_status_v1_status_proto = out.File
	file_connectext_grpc_status_v1_status_proto_rawDesc = nil
	file_connectext_grpc_status_v1_status_proto_goTypes = nil
	file_connectext_grpc_status_v1_status_proto_depIdxs = nil
}
