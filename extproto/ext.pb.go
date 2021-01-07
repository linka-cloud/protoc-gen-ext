// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: extproto/ext.proto

package extproto

import (
	reflect "reflect"

	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

var file_extproto_ext_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         10071,
		Name:          "extproto.hash_all",
		Tag:           "varint,10071,opt,name=hash_all",
		Filename:      "extproto/ext.proto",
	},
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         10072,
		Name:          "extproto.equal_all",
		Tag:           "varint,10072,opt,name=equal_all",
		Filename:      "extproto/ext.proto",
	},
	{
		ExtendedType:  (*descriptor.FileOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         10073,
		Name:          "extproto.merge_all",
		Tag:           "varint,10073,opt,name=merge_all",
		Filename:      "extproto/ext.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         10071,
		Name:          "extproto.skip_hashing",
		Tag:           "varint,10071,opt,name=skip_hashing",
		Filename:      "extproto/ext.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         10072,
		Name:          "extproto.skip_merging",
		Tag:           "varint,10072,opt,name=skip_merging",
		Filename:      "extproto/ext.proto",
	},
	{
		ExtendedType:  (*descriptor.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         10073,
		Name:          "extproto.sensitive",
		Tag:           "varint,10073,opt,name=sensitive",
		Filename:      "extproto/ext.proto",
	},
	{
		ExtendedType:  (*descriptor.OneofOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         10072,
		Name:          "extproto.skip_merging_oneof",
		Tag:           "varint,10072,opt,name=skip_merging_oneof",
		Filename:      "extproto/ext.proto",
	},
}

// Extension fields to descriptor.FileOptions.
var (
	// Disabled nullifies any validation rules for this message, including any
	// message fields associated with it that do support validation.
	//
	// optional bool hash_all = 10071;
	E_HashAll = &file_extproto_ext_proto_extTypes[0]
	// optional bool equal_all = 10072;
	E_EqualAll = &file_extproto_ext_proto_extTypes[1]
	// enabling merge_all for a file will generate a Merge(m) method for all Messages in the file.
	// Merge(m) will replace non-nil fields from a Proto passed as an override.
	//
	// optional bool merge_all = 10073;
	E_MergeAll = &file_extproto_ext_proto_extTypes[2]
)

// Extension fields to descriptor.FieldOptions.
var (
	// Rules specify the validations to be performed on this field. By default,
	// no validation is performed against a field.
	//
	// optional bool skip_hashing = 10071;
	E_SkipHashing = &file_extproto_ext_proto_extTypes[3]
	// This field will not be merged when a message's Merge() method is called.
	//
	// optional bool skip_merging = 10072;
	E_SkipMerging = &file_extproto_ext_proto_extTypes[4]
	// This field will not be merged when a message's Merge() method is called.
	//
	// optional bool sensitive = 10073;
	E_Sensitive = &file_extproto_ext_proto_extTypes[5]
)

// Extension fields to descriptor.OneofOptions.
var (
	// The fields in this oneof will not be merged when a message's Merge() method is called.
	//
	// optional bool skip_merging_oneof = 10072;
	E_SkipMergingOneof = &file_extproto_ext_proto_extTypes[6]
)

var File_extproto_ext_proto protoreflect.FileDescriptor

var file_extproto_ext_proto_rawDesc = []byte{
	0x0a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3a, 0x38, 0x0a, 0x08, 0x68, 0x61, 0x73, 0x68, 0x5f, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46,
	0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd7, 0x4e, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x68, 0x61, 0x73, 0x68, 0x41, 0x6c, 0x6c, 0x3a, 0x3a, 0x0a, 0x09, 0x65, 0x71,
	0x75, 0x61, 0x6c, 0x5f, 0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd8, 0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x65, 0x71,
	0x75, 0x61, 0x6c, 0x41, 0x6c, 0x6c, 0x3a, 0x3a, 0x0a, 0x09, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f,
	0x61, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd9, 0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x41,
	0x6c, 0x6c, 0x3a, 0x41, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x69,
	0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0xd7, 0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6b, 0x69, 0x70, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x6e, 0x67, 0x3a, 0x41, 0x0a, 0x0c, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x6d, 0x65,
	0x72, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd8, 0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0b, 0x73, 0x6b, 0x69,
	0x70, 0x4d, 0x65, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x3a, 0x3c, 0x0a, 0x09, 0x73, 0x65, 0x6e, 0x73,
	0x69, 0x74, 0x69, 0x76, 0x65, 0x12, 0x1d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x4f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd9, 0x4e, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x73, 0x65, 0x6e,
	0x73, 0x69, 0x74, 0x69, 0x76, 0x65, 0x3a, 0x4c, 0x0a, 0x12, 0x73, 0x6b, 0x69, 0x70, 0x5f, 0x6d,
	0x65, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x1d, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4f,
	0x6e, 0x65, 0x6f, 0x66, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xd8, 0x4e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x10, 0x73, 0x6b, 0x69, 0x70, 0x4d, 0x65, 0x72, 0x67, 0x69, 0x6e, 0x67, 0x4f,
	0x6e, 0x65, 0x6f, 0x66, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x65, 0x78, 0x74, 0x2f, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f,
	0x74, 0x6f,
}

var file_extproto_ext_proto_goTypes = []interface{}{
	(*descriptor.FileOptions)(nil),  // 0: google.protobuf.FileOptions
	(*descriptor.FieldOptions)(nil), // 1: google.protobuf.FieldOptions
	(*descriptor.OneofOptions)(nil), // 2: google.protobuf.OneofOptions
}
var file_extproto_ext_proto_depIdxs = []int32{
	0, // 0: extproto.hash_all:extendee -> google.protobuf.FileOptions
	0, // 1: extproto.equal_all:extendee -> google.protobuf.FileOptions
	0, // 2: extproto.merge_all:extendee -> google.protobuf.FileOptions
	1, // 3: extproto.skip_hashing:extendee -> google.protobuf.FieldOptions
	1, // 4: extproto.skip_merging:extendee -> google.protobuf.FieldOptions
	1, // 5: extproto.sensitive:extendee -> google.protobuf.FieldOptions
	2, // 6: extproto.skip_merging_oneof:extendee -> google.protobuf.OneofOptions
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	0, // [0:7] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_extproto_ext_proto_init() }
func file_extproto_ext_proto_init() {
	if File_extproto_ext_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_extproto_ext_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 7,
			NumServices:   0,
		},
		GoTypes:           file_extproto_ext_proto_goTypes,
		DependencyIndexes: file_extproto_ext_proto_depIdxs,
		ExtensionInfos:    file_extproto_ext_proto_extTypes,
	}.Build()
	File_extproto_ext_proto = out.File
	file_extproto_ext_proto_rawDesc = nil
	file_extproto_ext_proto_goTypes = nil
	file_extproto_ext_proto_depIdxs = nil
}
