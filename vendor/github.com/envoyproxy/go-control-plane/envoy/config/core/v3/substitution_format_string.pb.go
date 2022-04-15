// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.19.1
// source: envoy/config/core/v3/substitution_format_string.proto

package envoy_config_core_v3

import (
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/go-control-plane/envoy/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
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

// Configuration to use multiple :ref:`command operators <config_access_log_command_operators>`
// to generate a new string in either plain text or JSON format.
// [#next-free-field: 7]
type SubstitutionFormatString struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Format:
	//	*SubstitutionFormatString_TextFormat
	//	*SubstitutionFormatString_JsonFormat
	//	*SubstitutionFormatString_TextFormatSource
	Format isSubstitutionFormatString_Format `protobuf_oneof:"format"`
	// If set to true, when command operators are evaluated to null,
	//
	// * for ``text_format``, the output of the empty operator is changed from ``-`` to an
	//   empty string, so that empty values are omitted entirely.
	// * for ``json_format`` the keys with null values are omitted in the output structure.
	OmitEmptyValues bool `protobuf:"varint,3,opt,name=omit_empty_values,json=omitEmptyValues,proto3" json:"omit_empty_values,omitempty"`
	// Specify a *content_type* field.
	// If this field is not set then ``text/plain`` is used for *text_format* and
	// ``application/json`` is used for *json_format*.
	//
	// .. validated-code-block:: yaml
	//   :type-name: envoy.config.core.v3.SubstitutionFormatString
	//
	//   content_type: "text/html; charset=UTF-8"
	//
	ContentType string `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	// Specifies a collection of Formatter plugins that can be called from the access log configuration.
	// See the formatters extensions documentation for details.
	// [#extension-category: envoy.formatter]
	Formatters []*TypedExtensionConfig `protobuf:"bytes,6,rep,name=formatters,proto3" json:"formatters,omitempty"`
}

func (x *SubstitutionFormatString) Reset() {
	*x = SubstitutionFormatString{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_config_core_v3_substitution_format_string_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubstitutionFormatString) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubstitutionFormatString) ProtoMessage() {}

func (x *SubstitutionFormatString) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_config_core_v3_substitution_format_string_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubstitutionFormatString.ProtoReflect.Descriptor instead.
func (*SubstitutionFormatString) Descriptor() ([]byte, []int) {
	return file_envoy_config_core_v3_substitution_format_string_proto_rawDescGZIP(), []int{0}
}

func (m *SubstitutionFormatString) GetFormat() isSubstitutionFormatString_Format {
	if m != nil {
		return m.Format
	}
	return nil
}

// Deprecated: Do not use.
func (x *SubstitutionFormatString) GetTextFormat() string {
	if x, ok := x.GetFormat().(*SubstitutionFormatString_TextFormat); ok {
		return x.TextFormat
	}
	return ""
}

func (x *SubstitutionFormatString) GetJsonFormat() *_struct.Struct {
	if x, ok := x.GetFormat().(*SubstitutionFormatString_JsonFormat); ok {
		return x.JsonFormat
	}
	return nil
}

func (x *SubstitutionFormatString) GetTextFormatSource() *DataSource {
	if x, ok := x.GetFormat().(*SubstitutionFormatString_TextFormatSource); ok {
		return x.TextFormatSource
	}
	return nil
}

func (x *SubstitutionFormatString) GetOmitEmptyValues() bool {
	if x != nil {
		return x.OmitEmptyValues
	}
	return false
}

func (x *SubstitutionFormatString) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

func (x *SubstitutionFormatString) GetFormatters() []*TypedExtensionConfig {
	if x != nil {
		return x.Formatters
	}
	return nil
}

type isSubstitutionFormatString_Format interface {
	isSubstitutionFormatString_Format()
}

type SubstitutionFormatString_TextFormat struct {
	// Specify a format with command operators to form a text string.
	// Its details is described in :ref:`format string<config_access_log_format_strings>`.
	//
	// For example, setting ``text_format`` like below,
	//
	// .. validated-code-block:: yaml
	//   :type-name: envoy.config.core.v3.SubstitutionFormatString
	//
	//   text_format: "%LOCAL_REPLY_BODY%:%RESPONSE_CODE%:path=%REQ(:path)%\n"
	//
	// generates plain text similar to:
	//
	// .. code-block:: text
	//
	//   upstream connect error:503:path=/foo
	//
	// Deprecated in favor of :ref:`text_format_source <envoy_v3_api_field_config.core.v3.SubstitutionFormatString.text_format_source>`. To migrate text format strings, use the :ref:`inline_string <envoy_v3_api_field_config.core.v3.DataSource.inline_string>` field.
	//
	// Deprecated: Do not use.
	TextFormat string `protobuf:"bytes,1,opt,name=text_format,json=textFormat,proto3,oneof"`
}

type SubstitutionFormatString_JsonFormat struct {
	// Specify a format with command operators to form a JSON string.
	// Its details is described in :ref:`format dictionary<config_access_log_format_dictionaries>`.
	// Values are rendered as strings, numbers, or boolean values as appropriate.
	// Nested JSON objects may be produced by some command operators (e.g. FILTER_STATE or DYNAMIC_METADATA).
	// See the documentation for a specific command operator for details.
	//
	// .. validated-code-block:: yaml
	//   :type-name: envoy.config.core.v3.SubstitutionFormatString
	//
	//   json_format:
	//     status: "%RESPONSE_CODE%"
	//     message: "%LOCAL_REPLY_BODY%"
	//
	// The following JSON object would be created:
	//
	// .. code-block:: json
	//
	//  {
	//    "status": 500,
	//    "message": "My error message"
	//  }
	//
	JsonFormat *_struct.Struct `protobuf:"bytes,2,opt,name=json_format,json=jsonFormat,proto3,oneof"`
}

type SubstitutionFormatString_TextFormatSource struct {
	// Specify a format with command operators to form a text string.
	// Its details is described in :ref:`format string<config_access_log_format_strings>`.
	//
	// For example, setting ``text_format`` like below,
	//
	// .. validated-code-block:: yaml
	//   :type-name: envoy.config.core.v3.SubstitutionFormatString
	//
	//   text_format_source:
	//     inline_string: "%LOCAL_REPLY_BODY%:%RESPONSE_CODE%:path=%REQ(:path)%\n"
	//
	// generates plain text similar to:
	//
	// .. code-block:: text
	//
	//   upstream connect error:503:path=/foo
	//
	TextFormatSource *DataSource `protobuf:"bytes,5,opt,name=text_format_source,json=textFormatSource,proto3,oneof"`
}

func (*SubstitutionFormatString_TextFormat) isSubstitutionFormatString_Format() {}

func (*SubstitutionFormatString_JsonFormat) isSubstitutionFormatString_Format() {}

func (*SubstitutionFormatString_TextFormatSource) isSubstitutionFormatString_Format() {}

var File_envoy_config_core_v3_substitution_format_string_proto protoreflect.FileDescriptor

var file_envoy_config_core_v3_substitution_format_string_proto_rawDesc = []byte{
	0x0a, 0x35, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x2f, 0x76, 0x33, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x14, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x1a, 0x1f, 0x65,
	0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65,
	0x2f, 0x76, 0x33, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x24,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72,
	0x65, 0x2f, 0x76, 0x33, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x23, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x65, 0x70, 0x72, 0x65, 0x63, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x8c, 0x03, 0x0a, 0x18, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69, 0x6f, 0x6e,
	0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2e, 0x0a, 0x0b,
	0x74, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x0b, 0x18, 0x01, 0x92, 0xc7, 0x86, 0xd8, 0x04, 0x03, 0x33, 0x2e, 0x30, 0x48, 0x00,
	0x52, 0x0a, 0x74, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x44, 0x0a, 0x0b,
	0x6a, 0x73, 0x6f, 0x6e, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x75, 0x63, 0x74, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a,
	0x01, 0x02, 0x10, 0x01, 0x48, 0x00, 0x52, 0x0a, 0x6a, 0x73, 0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d,
	0x61, 0x74, 0x12, 0x50, 0x0a, 0x12, 0x74, 0x65, 0x78, 0x74, 0x5f, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x5f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f,
	0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x48, 0x00, 0x52, 0x10, 0x74, 0x65, 0x78, 0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x6f, 0x6d, 0x69, 0x74, 0x5f, 0x65, 0x6d, 0x70,
	0x74, 0x79, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0f, 0x6f, 0x6d, 0x69, 0x74, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x4a, 0x0a, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72,
	0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x54,
	0x79, 0x70, 0x65, 0x64, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x0a, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x74, 0x65, 0x72, 0x73, 0x42,
	0x0d, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x4d,
	0x0a, 0x22, 0x69, 0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x2e, 0x76, 0x33, 0x42, 0x1d, 0x53, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x69,
	0x6f, 0x6e, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02, 0x10, 0x02, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_config_core_v3_substitution_format_string_proto_rawDescOnce sync.Once
	file_envoy_config_core_v3_substitution_format_string_proto_rawDescData = file_envoy_config_core_v3_substitution_format_string_proto_rawDesc
)

func file_envoy_config_core_v3_substitution_format_string_proto_rawDescGZIP() []byte {
	file_envoy_config_core_v3_substitution_format_string_proto_rawDescOnce.Do(func() {
		file_envoy_config_core_v3_substitution_format_string_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_config_core_v3_substitution_format_string_proto_rawDescData)
	})
	return file_envoy_config_core_v3_substitution_format_string_proto_rawDescData
}

var file_envoy_config_core_v3_substitution_format_string_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_envoy_config_core_v3_substitution_format_string_proto_goTypes = []interface{}{
	(*SubstitutionFormatString)(nil), // 0: envoy.config.core.v3.SubstitutionFormatString
	(*_struct.Struct)(nil),           // 1: google.protobuf.Struct
	(*DataSource)(nil),               // 2: envoy.config.core.v3.DataSource
	(*TypedExtensionConfig)(nil),     // 3: envoy.config.core.v3.TypedExtensionConfig
}
var file_envoy_config_core_v3_substitution_format_string_proto_depIdxs = []int32{
	1, // 0: envoy.config.core.v3.SubstitutionFormatString.json_format:type_name -> google.protobuf.Struct
	2, // 1: envoy.config.core.v3.SubstitutionFormatString.text_format_source:type_name -> envoy.config.core.v3.DataSource
	3, // 2: envoy.config.core.v3.SubstitutionFormatString.formatters:type_name -> envoy.config.core.v3.TypedExtensionConfig
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_config_core_v3_substitution_format_string_proto_init() }
func file_envoy_config_core_v3_substitution_format_string_proto_init() {
	if File_envoy_config_core_v3_substitution_format_string_proto != nil {
		return
	}
	file_envoy_config_core_v3_base_proto_init()
	file_envoy_config_core_v3_extension_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_envoy_config_core_v3_substitution_format_string_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubstitutionFormatString); i {
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
	file_envoy_config_core_v3_substitution_format_string_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SubstitutionFormatString_TextFormat)(nil),
		(*SubstitutionFormatString_JsonFormat)(nil),
		(*SubstitutionFormatString_TextFormatSource)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_envoy_config_core_v3_substitution_format_string_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_config_core_v3_substitution_format_string_proto_goTypes,
		DependencyIndexes: file_envoy_config_core_v3_substitution_format_string_proto_depIdxs,
		MessageInfos:      file_envoy_config_core_v3_substitution_format_string_proto_msgTypes,
	}.Build()
	File_envoy_config_core_v3_substitution_format_string_proto = out.File
	file_envoy_config_core_v3_substitution_format_string_proto_rawDesc = nil
	file_envoy_config_core_v3_substitution_format_string_proto_goTypes = nil
	file_envoy_config_core_v3_substitution_format_string_proto_depIdxs = nil
}
