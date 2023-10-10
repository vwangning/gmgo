// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.16.0
// source: envoy/extensions/filters/http/ip_tagging/v3/ip_tagging.proto

package envoy_extensions_filters_http_ip_tagging_v3

import (
	v3 "gitee.com/zhaochuninhefei/gmgo/go-control-plane/envoy/config/core/v3"
	_ "github.com/cncf/xds/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// The type of requests the filter should apply to. The supported types
// are internal, external or both. The
// :ref:`x-forwarded-for<config_http_conn_man_headers_x-forwarded-for_internal_origin>` header is
// used to determine if a request is internal and will result in
// :ref:`x-envoy-internal<config_http_conn_man_headers_x-envoy-internal>`
// being set. The filter defaults to both, and it will apply to all request types.
type IPTagging_RequestType int32

const (
	// Both external and internal requests will be tagged. This is the default value.
	IPTagging_BOTH IPTagging_RequestType = 0
	// Only internal requests will be tagged.
	IPTagging_INTERNAL IPTagging_RequestType = 1
	// Only external requests will be tagged.
	IPTagging_EXTERNAL IPTagging_RequestType = 2
)

// Enum value maps for IPTagging_RequestType.
var (
	IPTagging_RequestType_name = map[int32]string{
		0: "BOTH",
		1: "INTERNAL",
		2: "EXTERNAL",
	}
	IPTagging_RequestType_value = map[string]int32{
		"BOTH":     0,
		"INTERNAL": 1,
		"EXTERNAL": 2,
	}
)

func (x IPTagging_RequestType) Enum() *IPTagging_RequestType {
	p := new(IPTagging_RequestType)
	*p = x
	return p
}

func (x IPTagging_RequestType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IPTagging_RequestType) Descriptor() protoreflect.EnumDescriptor {
	return file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_enumTypes[0].Descriptor()
}

func (IPTagging_RequestType) Type() protoreflect.EnumType {
	return &file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_enumTypes[0]
}

func (x IPTagging_RequestType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IPTagging_RequestType.Descriptor instead.
func (IPTagging_RequestType) EnumDescriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescGZIP(), []int{0, 0}
}

type IPTagging struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The type of request the filter should apply to.
	RequestType IPTagging_RequestType `protobuf:"varint,1,opt,name=request_type,json=requestType,proto3,enum=envoy.extensions.filters.http.ip_tagging.v3.IPTagging_RequestType" json:"request_type,omitempty"`
	// [#comment:TODO(ccaraman): Extend functionality to load IP tags from file system.
	// Tracked by issue https://github.com/envoyproxy/envoy/issues/2695]
	// The set of IP tags for the filter.
	IpTags []*IPTagging_IPTag `protobuf:"bytes,4,rep,name=ip_tags,json=ipTags,proto3" json:"ip_tags,omitempty"`
}

func (x *IPTagging) Reset() {
	*x = IPTagging{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPTagging) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPTagging) ProtoMessage() {}

func (x *IPTagging) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPTagging.ProtoReflect.Descriptor instead.
func (*IPTagging) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescGZIP(), []int{0}
}

func (x *IPTagging) GetRequestType() IPTagging_RequestType {
	if x != nil {
		return x.RequestType
	}
	return IPTagging_BOTH
}

func (x *IPTagging) GetIpTags() []*IPTagging_IPTag {
	if x != nil {
		return x.IpTags
	}
	return nil
}

// Supplies the IP tag name and the IP address subnets.
type IPTagging_IPTag struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Specifies the IP tag name to apply.
	IpTagName string `protobuf:"bytes,1,opt,name=ip_tag_name,json=ipTagName,proto3" json:"ip_tag_name,omitempty"`
	// A list of IP address subnets that will be tagged with
	// ip_tag_name. Both IPv4 and IPv6 are supported.
	IpList []*v3.CidrRange `protobuf:"bytes,2,rep,name=ip_list,json=ipList,proto3" json:"ip_list,omitempty"`
}

func (x *IPTagging_IPTag) Reset() {
	*x = IPTagging_IPTag{}
	if protoimpl.UnsafeEnabled {
		mi := &file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IPTagging_IPTag) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IPTagging_IPTag) ProtoMessage() {}

func (x *IPTagging_IPTag) ProtoReflect() protoreflect.Message {
	mi := &file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IPTagging_IPTag.ProtoReflect.Descriptor instead.
func (*IPTagging_IPTag) Descriptor() ([]byte, []int) {
	return file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescGZIP(), []int{0, 0}
}

func (x *IPTagging_IPTag) GetIpTagName() string {
	if x != nil {
		return x.IpTagName
	}
	return ""
}

func (x *IPTagging_IPTag) GetIpList() []*v3.CidrRange {
	if x != nil {
		return x.IpList
	}
	return nil
}

var File_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto protoreflect.FileDescriptor

var file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x2f, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2f,
	0x69, 0x70, 0x5f, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x33, 0x2f, 0x69, 0x70,
	0x5f, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x2b,
	0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x69, 0x70,
	0x5f, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x1a, 0x22, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x2f, 0x76,
	0x33, 0x2f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1d, 0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21,
	0x75, 0x64, 0x70, 0x61, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x03, 0x0a, 0x09, 0x49,
	0x50, 0x54, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x12, 0x6f, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x42,
	0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x69,
	0x70, 0x5f, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x50, 0x54,
	0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x82, 0x01, 0x02, 0x10, 0x01, 0x52, 0x0b, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x5f, 0x0a, 0x07, 0x69, 0x70, 0x5f,
	0x74, 0x61, 0x67, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x3c, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x69, 0x70, 0x5f, 0x74, 0x61,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x2e, 0x49, 0x50, 0x54, 0x61, 0x67, 0x67, 0x69,
	0x6e, 0x67, 0x2e, 0x49, 0x50, 0x54, 0x61, 0x67, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x92, 0x01, 0x02,
	0x08, 0x01, 0x52, 0x06, 0x69, 0x70, 0x54, 0x61, 0x67, 0x73, 0x1a, 0xa0, 0x01, 0x0a, 0x05, 0x49,
	0x50, 0x54, 0x61, 0x67, 0x12, 0x1e, 0x0a, 0x0b, 0x69, 0x70, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x54, 0x61, 0x67,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x38, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x33, 0x2e, 0x43, 0x69, 0x64,
	0x72, 0x52, 0x61, 0x6e, 0x67, 0x65, 0x52, 0x06, 0x69, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x3a, 0x3d,
	0x9a, 0xc5, 0x88, 0x1e, 0x38, 0x0a, 0x36, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x2e, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e,
	0x69, 0x70, 0x5f, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x32, 0x2e, 0x49, 0x50,
	0x54, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x49, 0x50, 0x54, 0x61, 0x67, 0x22, 0x33, 0x0a,
	0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04,
	0x42, 0x4f, 0x54, 0x48, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e,
	0x41, 0x4c, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x10, 0x02, 0x3a, 0x37, 0x9a, 0xc5, 0x88, 0x1e, 0x32, 0x0a, 0x30, 0x65, 0x6e, 0x76, 0x6f, 0x79,
	0x2e, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x68,
	0x74, 0x74, 0x70, 0x2e, 0x69, 0x70, 0x5f, 0x74, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76,
	0x32, 0x2e, 0x49, 0x50, 0x54, 0x61, 0x67, 0x67, 0x69, 0x6e, 0x67, 0x42, 0x55, 0x0a, 0x39, 0x69,
	0x6f, 0x2e, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78, 0x79, 0x2e, 0x65, 0x6e, 0x76,
	0x6f, 0x79, 0x2e, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x66, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x69, 0x70, 0x5f, 0x74, 0x61,
	0x67, 0x67, 0x69, 0x6e, 0x67, 0x2e, 0x76, 0x33, 0x42, 0x0e, 0x49, 0x70, 0x54, 0x61, 0x67, 0x67,
	0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0xba, 0x80, 0xc8, 0xd1, 0x06, 0x02,
	0x10, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescOnce sync.Once
	file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescData = file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDesc
)

func file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescGZIP() []byte {
	file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescOnce.Do(func() {
		file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescData = protoimpl.X.CompressGZIP(file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescData)
	})
	return file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDescData
}

var file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_goTypes = []interface{}{
	(IPTagging_RequestType)(0), // 0: envoy.extensions.filters.http.ip_tagging.v3.IPTagging.RequestType
	(*IPTagging)(nil),          // 1: envoy.extensions.filters.http.ip_tagging.v3.IPTagging
	(*IPTagging_IPTag)(nil),    // 2: envoy.extensions.filters.http.ip_tagging.v3.IPTagging.IPTag
	(*v3.CidrRange)(nil),       // 3: envoy.config.core.v3.CidrRange
}
var file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_depIdxs = []int32{
	0, // 0: envoy.extensions.filters.http.ip_tagging.v3.IPTagging.request_type:type_name -> envoy.extensions.filters.http.ip_tagging.v3.IPTagging.RequestType
	2, // 1: envoy.extensions.filters.http.ip_tagging.v3.IPTagging.ip_tags:type_name -> envoy.extensions.filters.http.ip_tagging.v3.IPTagging.IPTag
	3, // 2: envoy.extensions.filters.http.ip_tagging.v3.IPTagging.IPTag.ip_list:type_name -> envoy.config.core.v3.CidrRange
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_init() }
func file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_init() {
	if File_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPTagging); i {
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
		file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IPTagging_IPTag); i {
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
			RawDescriptor: file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_goTypes,
		DependencyIndexes: file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_depIdxs,
		EnumInfos:         file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_enumTypes,
		MessageInfos:      file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_msgTypes,
	}.Build()
	File_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto = out.File
	file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_rawDesc = nil
	file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_goTypes = nil
	file_envoy_extensions_filters_http_ip_tagging_v3_ip_tagging_proto_depIdxs = nil
}