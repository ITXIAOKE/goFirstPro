// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.2
// source: google/cloud/integrations/v1alpha/json_validation.proto

package integrations

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Options for how to validate json schemas.
type JsonValidationOption int32

const (
	// As per the default behavior, no validation will be run. Will not override
	// any option set in a Task.
	JsonValidationOption_JSON_VALIDATION_OPTION_UNSPECIFIED JsonValidationOption = 0
	// Do not run any validation against JSON schemas.
	JsonValidationOption_SKIP JsonValidationOption = 1
	// Validate all potential input JSON parameters against schemas specified in
	// IntegrationParameter.
	JsonValidationOption_PRE_EXECUTION JsonValidationOption = 2
	// Validate all potential output JSON parameters against schemas specified in
	// IntegrationParameter.
	JsonValidationOption_POST_EXECUTION JsonValidationOption = 3
	// Perform both PRE_EXECUTION and POST_EXECUTION validations.
	JsonValidationOption_PRE_POST_EXECUTION JsonValidationOption = 4
)

// Enum value maps for JsonValidationOption.
var (
	JsonValidationOption_name = map[int32]string{
		0: "JSON_VALIDATION_OPTION_UNSPECIFIED",
		1: "SKIP",
		2: "PRE_EXECUTION",
		3: "POST_EXECUTION",
		4: "PRE_POST_EXECUTION",
	}
	JsonValidationOption_value = map[string]int32{
		"JSON_VALIDATION_OPTION_UNSPECIFIED": 0,
		"SKIP":                               1,
		"PRE_EXECUTION":                      2,
		"POST_EXECUTION":                     3,
		"PRE_POST_EXECUTION":                 4,
	}
)

func (x JsonValidationOption) Enum() *JsonValidationOption {
	p := new(JsonValidationOption)
	*p = x
	return p
}

func (x JsonValidationOption) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (JsonValidationOption) Descriptor() protoreflect.EnumDescriptor {
	return file_google_cloud_integrations_v1alpha_json_validation_proto_enumTypes[0].Descriptor()
}

func (JsonValidationOption) Type() protoreflect.EnumType {
	return &file_google_cloud_integrations_v1alpha_json_validation_proto_enumTypes[0]
}

func (x JsonValidationOption) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use JsonValidationOption.Descriptor instead.
func (JsonValidationOption) EnumDescriptor() ([]byte, []int) {
	return file_google_cloud_integrations_v1alpha_json_validation_proto_rawDescGZIP(), []int{0}
}

var File_google_cloud_integrations_v1alpha_json_validation_proto protoreflect.FileDescriptor

var file_google_cloud_integrations_v1alpha_json_validation_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x21, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2a, 0x87, 0x01, 0x0a,
	0x14, 0x4a, 0x73, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4f,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x26, 0x0a, 0x22, 0x4a, 0x53, 0x4f, 0x4e, 0x5f, 0x56, 0x41,
	0x4c, 0x49, 0x44, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x4f, 0x50, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x53, 0x4b, 0x49, 0x50, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x50, 0x52, 0x45, 0x5f, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x12, 0x0a, 0x0e, 0x50, 0x4f,
	0x53, 0x54, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x03, 0x12, 0x16,
	0x0a, 0x12, 0x50, 0x52, 0x45, 0x5f, 0x50, 0x4f, 0x53, 0x54, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x10, 0x04, 0x42, 0x8d, 0x01, 0x0a, 0x25, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x42, 0x13, 0x4a, 0x73, 0x6f, 0x6e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2f, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x3b, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_google_cloud_integrations_v1alpha_json_validation_proto_rawDescOnce sync.Once
	file_google_cloud_integrations_v1alpha_json_validation_proto_rawDescData = file_google_cloud_integrations_v1alpha_json_validation_proto_rawDesc
)

func file_google_cloud_integrations_v1alpha_json_validation_proto_rawDescGZIP() []byte {
	file_google_cloud_integrations_v1alpha_json_validation_proto_rawDescOnce.Do(func() {
		file_google_cloud_integrations_v1alpha_json_validation_proto_rawDescData = protoimpl.X.CompressGZIP(file_google_cloud_integrations_v1alpha_json_validation_proto_rawDescData)
	})
	return file_google_cloud_integrations_v1alpha_json_validation_proto_rawDescData
}

var file_google_cloud_integrations_v1alpha_json_validation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_google_cloud_integrations_v1alpha_json_validation_proto_goTypes = []interface{}{
	(JsonValidationOption)(0), // 0: google.cloud.integrations.v1alpha.JsonValidationOption
}
var file_google_cloud_integrations_v1alpha_json_validation_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_cloud_integrations_v1alpha_json_validation_proto_init() }
func file_google_cloud_integrations_v1alpha_json_validation_proto_init() {
	if File_google_cloud_integrations_v1alpha_json_validation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_cloud_integrations_v1alpha_json_validation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_google_cloud_integrations_v1alpha_json_validation_proto_goTypes,
		DependencyIndexes: file_google_cloud_integrations_v1alpha_json_validation_proto_depIdxs,
		EnumInfos:         file_google_cloud_integrations_v1alpha_json_validation_proto_enumTypes,
	}.Build()
	File_google_cloud_integrations_v1alpha_json_validation_proto = out.File
	file_google_cloud_integrations_v1alpha_json_validation_proto_rawDesc = nil
	file_google_cloud_integrations_v1alpha_json_validation_proto_goTypes = nil
	file_google_cloud_integrations_v1alpha_json_validation_proto_depIdxs = nil
}
