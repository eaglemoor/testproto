// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: location.proto

package testproto

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

type Locale int32

const (
	Locale_LOCALE_UNSPECIFIED Locale = 0
	Locale_LOCALE_SPORTS      Locale = 1
	Locale_LOCALE_CYBER       Locale = 2
	Locale_LOCALE_UA_TRIBUNA  Locale = 3
	Locale_LOCALE_BY_TRIBUNA  Locale = 4
)

// Enum value maps for Locale.
var (
	Locale_name = map[int32]string{
		0: "LOCALE_UNSPECIFIED",
		1: "LOCALE_SPORTS",
		2: "LOCALE_CYBER",
		3: "LOCALE_UA_TRIBUNA",
		4: "LOCALE_BY_TRIBUNA",
	}
	Locale_value = map[string]int32{
		"LOCALE_UNSPECIFIED": 0,
		"LOCALE_SPORTS":      1,
		"LOCALE_CYBER":       2,
		"LOCALE_UA_TRIBUNA":  3,
		"LOCALE_BY_TRIBUNA":  4,
	}
)

func (x Locale) Enum() *Locale {
	p := new(Locale)
	*p = x
	return p
}

func (x Locale) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Locale) Descriptor() protoreflect.EnumDescriptor {
	return file_location_proto_enumTypes[0].Descriptor()
}

func (Locale) Type() protoreflect.EnumType {
	return &file_location_proto_enumTypes[0]
}

func (x Locale) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Locale.Descriptor instead.
func (Locale) EnumDescriptor() ([]byte, []int) {
	return file_location_proto_rawDescGZIP(), []int{0}
}

var File_location_proto protoreflect.FileDescriptor

var file_location_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x13, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x6d, 0x6f, 0x6f, 0x72, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x73, 0x0a, 0x06, 0x4c, 0x6f, 0x63, 0x61, 0x6c, 0x65, 0x12,
	0x16, 0x0a, 0x12, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x4f, 0x43, 0x41, 0x4c,
	0x45, 0x5f, 0x53, 0x50, 0x4f, 0x52, 0x54, 0x53, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4c, 0x4f,
	0x43, 0x41, 0x4c, 0x45, 0x5f, 0x43, 0x59, 0x42, 0x45, 0x52, 0x10, 0x02, 0x12, 0x15, 0x0a, 0x11,
	0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x5f, 0x55, 0x41, 0x5f, 0x54, 0x52, 0x49, 0x42, 0x55, 0x4e,
	0x41, 0x10, 0x03, 0x12, 0x15, 0x0a, 0x11, 0x4c, 0x4f, 0x43, 0x41, 0x4c, 0x45, 0x5f, 0x42, 0x59,
	0x5f, 0x54, 0x52, 0x49, 0x42, 0x55, 0x4e, 0x41, 0x10, 0x04, 0x42, 0x20, 0x5a, 0x1e, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x61, 0x67, 0x6c, 0x65, 0x6d, 0x6f,
	0x6f, 0x72, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_location_proto_rawDescOnce sync.Once
	file_location_proto_rawDescData = file_location_proto_rawDesc
)

func file_location_proto_rawDescGZIP() []byte {
	file_location_proto_rawDescOnce.Do(func() {
		file_location_proto_rawDescData = protoimpl.X.CompressGZIP(file_location_proto_rawDescData)
	})
	return file_location_proto_rawDescData
}

var file_location_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_location_proto_goTypes = []interface{}{
	(Locale)(0), // 0: eaglemoor.testproto.Locale
}
var file_location_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_location_proto_init() }
func file_location_proto_init() {
	if File_location_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_location_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_location_proto_goTypes,
		DependencyIndexes: file_location_proto_depIdxs,
		EnumInfos:         file_location_proto_enumTypes,
	}.Build()
	File_location_proto = out.File
	file_location_proto_rawDesc = nil
	file_location_proto_goTypes = nil
	file_location_proto_depIdxs = nil
}
