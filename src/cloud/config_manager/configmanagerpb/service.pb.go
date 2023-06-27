// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/cloud/config_manager/configmanagerpb/service.proto

package configmanagerpb

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_sortkeys "github.com/gogo/protobuf/sortkeys"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
	vizierconfigpb "px.dev/pixie/src/api/proto/vizierconfigpb"
	reflect "reflect"
	strings "strings"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type ConfigForVizierRequest struct {
	Namespace  string                     `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	VzSpec     *vizierconfigpb.VizierSpec `protobuf:"bytes,2,opt,name=vz_spec,json=vzSpec,proto3" json:"vz_spec,omitempty"`
	K8sVersion string                     `protobuf:"bytes,3,opt,name=k8s_version,json=k8sVersion,proto3" json:"k8s_version,omitempty"`
}

func (m *ConfigForVizierRequest) Reset()      { *m = ConfigForVizierRequest{} }
func (*ConfigForVizierRequest) ProtoMessage() {}
func (*ConfigForVizierRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c1886abb7750df, []int{0}
}
func (m *ConfigForVizierRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigForVizierRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigForVizierRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigForVizierRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigForVizierRequest.Merge(m, src)
}
func (m *ConfigForVizierRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConfigForVizierRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigForVizierRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigForVizierRequest proto.InternalMessageInfo

func (m *ConfigForVizierRequest) GetNamespace() string {
	if m != nil {
		return m.Namespace
	}
	return ""
}

func (m *ConfigForVizierRequest) GetVzSpec() *vizierconfigpb.VizierSpec {
	if m != nil {
		return m.VzSpec
	}
	return nil
}

func (m *ConfigForVizierRequest) GetK8sVersion() string {
	if m != nil {
		return m.K8sVersion
	}
	return ""
}

type ConfigForVizierResponse struct {
	NameToYamlContent map[string]string `protobuf:"bytes,1,rep,name=nameToYamlContent,proto3" json:"nameToYamlContent,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SentryDSN         string            `protobuf:"bytes,2,opt,name=sentry_dsn,json=sentryDsn,proto3" json:"sentry_dsn,omitempty"`
}

func (m *ConfigForVizierResponse) Reset()      { *m = ConfigForVizierResponse{} }
func (*ConfigForVizierResponse) ProtoMessage() {}
func (*ConfigForVizierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c1886abb7750df, []int{1}
}
func (m *ConfigForVizierResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigForVizierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigForVizierResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigForVizierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigForVizierResponse.Merge(m, src)
}
func (m *ConfigForVizierResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConfigForVizierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigForVizierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigForVizierResponse proto.InternalMessageInfo

func (m *ConfigForVizierResponse) GetNameToYamlContent() map[string]string {
	if m != nil {
		return m.NameToYamlContent
	}
	return nil
}

func (m *ConfigForVizierResponse) GetSentryDSN() string {
	if m != nil {
		return m.SentryDSN
	}
	return ""
}

type ConfigForOperatorRequest struct {
}

func (m *ConfigForOperatorRequest) Reset()      { *m = ConfigForOperatorRequest{} }
func (*ConfigForOperatorRequest) ProtoMessage() {}
func (*ConfigForOperatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c1886abb7750df, []int{2}
}
func (m *ConfigForOperatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigForOperatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigForOperatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigForOperatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigForOperatorRequest.Merge(m, src)
}
func (m *ConfigForOperatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *ConfigForOperatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigForOperatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigForOperatorRequest proto.InternalMessageInfo

type ConfigForOperatorResponse struct {
	SentryOperatorDSN string `protobuf:"bytes,1,opt,name=sentry_operator_dsn,json=sentryOperatorDsn,proto3" json:"sentry_operator_dsn,omitempty"`
}

func (m *ConfigForOperatorResponse) Reset()      { *m = ConfigForOperatorResponse{} }
func (*ConfigForOperatorResponse) ProtoMessage() {}
func (*ConfigForOperatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_11c1886abb7750df, []int{3}
}
func (m *ConfigForOperatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConfigForOperatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConfigForOperatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConfigForOperatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigForOperatorResponse.Merge(m, src)
}
func (m *ConfigForOperatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *ConfigForOperatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigForOperatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigForOperatorResponse proto.InternalMessageInfo

func (m *ConfigForOperatorResponse) GetSentryOperatorDSN() string {
	if m != nil {
		return m.SentryOperatorDSN
	}
	return ""
}

func init() {
	proto.RegisterType((*ConfigForVizierRequest)(nil), "px.services.ConfigForVizierRequest")
	proto.RegisterType((*ConfigForVizierResponse)(nil), "px.services.ConfigForVizierResponse")
	proto.RegisterMapType((map[string]string)(nil), "px.services.ConfigForVizierResponse.NameToYamlContentEntry")
	proto.RegisterType((*ConfigForOperatorRequest)(nil), "px.services.ConfigForOperatorRequest")
	proto.RegisterType((*ConfigForOperatorResponse)(nil), "px.services.ConfigForOperatorResponse")
}

func init() {
	proto.RegisterFile("src/cloud/config_manager/configmanagerpb/service.proto", fileDescriptor_11c1886abb7750df)
}

var fileDescriptor_11c1886abb7750df = []byte{
	// 538 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x8f, 0xd2, 0x4e,
	0x1c, 0xee, 0x2c, 0xf9, 0xef, 0x3f, 0x0c, 0xf1, 0x85, 0x11, 0x57, 0x6c, 0x74, 0x20, 0xf8, 0x12,
	0x0e, 0xda, 0x26, 0x98, 0x6c, 0x88, 0xde, 0x58, 0xd0, 0x83, 0x71, 0x4d, 0x8a, 0xd9, 0x44, 0x2f,
	0x4d, 0xdb, 0x9d, 0xc5, 0x06, 0x98, 0x19, 0x3b, 0x6d, 0xb3, 0x70, 0xf2, 0x23, 0xf8, 0x0d, 0xbc,
	0xfa, 0x51, 0x3c, 0x72, 0xdc, 0x13, 0x91, 0xe1, 0xe2, 0x71, 0x8f, 0x1e, 0x4d, 0x67, 0xca, 0xba,
	0xcb, 0xe2, 0xcb, 0x89, 0x79, 0x7e, 0x2f, 0xcf, 0xf3, 0xf0, 0x4c, 0x07, 0xee, 0x8a, 0x28, 0xb0,
	0x83, 0x11, 0x4b, 0x0e, 0xed, 0x80, 0xd1, 0xa3, 0x70, 0xe0, 0x8e, 0x3d, 0xea, 0x0d, 0x48, 0x94,
	0xc3, 0x1c, 0x71, 0xdf, 0x16, 0x24, 0x4a, 0xc3, 0x80, 0x58, 0x3c, 0x62, 0x31, 0x43, 0x25, 0x7e,
	0x6c, 0xe5, 0x15, 0x61, 0x3e, 0x1e, 0x84, 0xf1, 0xfb, 0xc4, 0xb7, 0x02, 0x36, 0xb6, 0x07, 0x6c,
	0xc0, 0x6c, 0x35, 0xe3, 0x27, 0x47, 0x0a, 0x29, 0xa0, 0x4e, 0x7a, 0xd7, 0xb4, 0x33, 0x4d, 0x8f,
	0x87, 0x7a, 0xcc, 0x4e, 0xc3, 0x69, 0x48, 0x22, 0x2d, 0xc7, 0xfd, 0x1c, 0xba, 0xf1, 0x84, 0x13,
	0xa1, 0x17, 0x1a, 0x9f, 0x01, 0xdc, 0xd9, 0x53, 0xfd, 0xe7, 0x2c, 0x3a, 0x50, 0x7d, 0x87, 0x7c,
	0x48, 0x88, 0x88, 0xd1, 0x1d, 0x58, 0xa4, 0xde, 0x98, 0x08, 0xee, 0x05, 0xa4, 0x0a, 0xea, 0xa0,
	0x59, 0x74, 0x7e, 0x15, 0xd0, 0x2e, 0xfc, 0x3f, 0x9d, 0xba, 0x82, 0x93, 0xa0, 0xba, 0x55, 0x07,
	0xcd, 0x52, 0xeb, 0xae, 0xc5, 0x8f, 0xad, 0x8b, 0x82, 0x96, 0x26, 0xec, 0x73, 0x12, 0x38, 0xdb,
	0xe9, 0x34, 0xfb, 0x45, 0x36, 0x2c, 0x0d, 0xdb, 0xc2, 0x4d, 0x49, 0x24, 0x42, 0x46, 0xab, 0x85,
	0x8c, 0xb7, 0x73, 0x55, 0xce, 0x6b, 0xf0, 0x65, 0x5b, 0x1c, 0xe8, 0xaa, 0x03, 0x87, 0x67, 0xe7,
	0xc6, 0x0f, 0x00, 0x6f, 0x5d, 0x72, 0x28, 0x38, 0xa3, 0x82, 0xa0, 0x10, 0x96, 0x33, 0x47, 0x6f,
	0xd8, 0x5b, 0x6f, 0x3c, 0xda, 0x63, 0x34, 0x26, 0x34, 0xae, 0x82, 0x7a, 0xa1, 0x59, 0x6a, 0x3d,
	0xb3, 0xce, 0xc5, 0x68, 0xfd, 0x86, 0xc0, 0xda, 0x5f, 0xdf, 0xee, 0xd1, 0x38, 0x9a, 0x38, 0x97,
	0x59, 0xd1, 0x23, 0x08, 0x05, 0xc9, 0x9a, 0xee, 0xa1, 0xa0, 0xea, 0x2f, 0x17, 0x3b, 0x57, 0xe4,
	0xbc, 0x56, 0xec, 0xab, 0x6a, 0xb7, 0xbf, 0xef, 0x14, 0xf5, 0x40, 0x57, 0x50, 0xb3, 0x0b, 0x77,
	0x36, 0x53, 0xa3, 0xeb, 0xb0, 0x30, 0x24, 0x93, 0x3c, 0xcf, 0xec, 0x88, 0x2a, 0xf0, 0xbf, 0xd4,
	0x1b, 0x25, 0x44, 0x93, 0x3a, 0x1a, 0x3c, 0xdd, 0x6a, 0x83, 0x86, 0x09, 0xab, 0x67, 0xc6, 0x5f,
	0x73, 0x12, 0x79, 0x31, 0x5b, 0xdd, 0x4e, 0xc3, 0x87, 0xb7, 0x37, 0xf4, 0xf2, 0x5c, 0x7a, 0xf0,
	0x46, 0x6e, 0x96, 0xe5, 0x2d, 0xe5, 0x5a, 0x89, 0x76, 0x6e, 0xca, 0x79, 0xad, 0xac, 0x5d, 0xaf,
	0x16, 0x33, 0xf7, 0x65, 0x71, 0xb1, 0x24, 0x68, 0x6b, 0x0e, 0x60, 0x45, 0x8b, 0xbc, 0xd2, 0xdf,
	0x6a, 0x5f, 0xe7, 0x89, 0x5c, 0x88, 0x5e, 0x90, 0x78, 0x2d, 0x54, 0x74, 0xef, 0xcf, 0x91, 0x2b,
	0xdf, 0xe6, 0xfd, 0x7f, 0xb9, 0x17, 0x44, 0x60, 0xe5, 0xbc, 0xc0, 0xca, 0x14, 0x7a, 0xb0, 0x79,
	0x7b, 0x2d, 0x1c, 0xf3, 0xe1, 0xdf, 0xc6, 0xb4, 0x4c, 0xa7, 0x37, 0x5b, 0x60, 0xe3, 0x64, 0x81,
	0x8d, 0xd3, 0x05, 0x06, 0x1f, 0x25, 0x06, 0x5f, 0x24, 0x06, 0x5f, 0x25, 0x06, 0x33, 0x89, 0xc1,
	0x37, 0x89, 0xc1, 0x77, 0x89, 0x8d, 0x53, 0x89, 0xc1, 0xa7, 0x25, 0x36, 0x66, 0x4b, 0x6c, 0x9c,
	0x2c, 0xb1, 0xf1, 0xee, 0xda, 0xda, 0xfb, 0xf5, 0xb7, 0xd5, 0x5b, 0x7a, 0xf2, 0x33, 0x00, 0x00,
	0xff, 0xff, 0x78, 0x32, 0x6b, 0x8a, 0xf2, 0x03, 0x00, 0x00,
}

func (this *ConfigForVizierRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfigForVizierRequest)
	if !ok {
		that2, ok := that.(ConfigForVizierRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Namespace != that1.Namespace {
		return false
	}
	if !this.VzSpec.Equal(that1.VzSpec) {
		return false
	}
	if this.K8sVersion != that1.K8sVersion {
		return false
	}
	return true
}
func (this *ConfigForVizierResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfigForVizierResponse)
	if !ok {
		that2, ok := that.(ConfigForVizierResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.NameToYamlContent) != len(that1.NameToYamlContent) {
		return false
	}
	for i := range this.NameToYamlContent {
		if this.NameToYamlContent[i] != that1.NameToYamlContent[i] {
			return false
		}
	}
	if this.SentryDSN != that1.SentryDSN {
		return false
	}
	return true
}
func (this *ConfigForOperatorRequest) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfigForOperatorRequest)
	if !ok {
		that2, ok := that.(ConfigForOperatorRequest)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	return true
}
func (this *ConfigForOperatorResponse) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ConfigForOperatorResponse)
	if !ok {
		that2, ok := that.(ConfigForOperatorResponse)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.SentryOperatorDSN != that1.SentryOperatorDSN {
		return false
	}
	return true
}
func (this *ConfigForVizierRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&configmanagerpb.ConfigForVizierRequest{")
	s = append(s, "Namespace: "+fmt.Sprintf("%#v", this.Namespace)+",\n")
	if this.VzSpec != nil {
		s = append(s, "VzSpec: "+fmt.Sprintf("%#v", this.VzSpec)+",\n")
	}
	s = append(s, "K8sVersion: "+fmt.Sprintf("%#v", this.K8sVersion)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ConfigForVizierResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&configmanagerpb.ConfigForVizierResponse{")
	keysForNameToYamlContent := make([]string, 0, len(this.NameToYamlContent))
	for k, _ := range this.NameToYamlContent {
		keysForNameToYamlContent = append(keysForNameToYamlContent, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForNameToYamlContent)
	mapStringForNameToYamlContent := "map[string]string{"
	for _, k := range keysForNameToYamlContent {
		mapStringForNameToYamlContent += fmt.Sprintf("%#v: %#v,", k, this.NameToYamlContent[k])
	}
	mapStringForNameToYamlContent += "}"
	if this.NameToYamlContent != nil {
		s = append(s, "NameToYamlContent: "+mapStringForNameToYamlContent+",\n")
	}
	s = append(s, "SentryDSN: "+fmt.Sprintf("%#v", this.SentryDSN)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ConfigForOperatorRequest) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 4)
	s = append(s, "&configmanagerpb.ConfigForOperatorRequest{")
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *ConfigForOperatorResponse) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 5)
	s = append(s, "&configmanagerpb.ConfigForOperatorResponse{")
	s = append(s, "SentryOperatorDSN: "+fmt.Sprintf("%#v", this.SentryOperatorDSN)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringService(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConfigManagerServiceClient is the client API for ConfigManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConfigManagerServiceClient interface {
	GetConfigForVizier(ctx context.Context, in *ConfigForVizierRequest, opts ...grpc.CallOption) (*ConfigForVizierResponse, error)
	GetConfigForOperator(ctx context.Context, in *ConfigForOperatorRequest, opts ...grpc.CallOption) (*ConfigForOperatorResponse, error)
}

type configManagerServiceClient struct {
	cc *grpc.ClientConn
}

func NewConfigManagerServiceClient(cc *grpc.ClientConn) ConfigManagerServiceClient {
	return &configManagerServiceClient{cc}
}

func (c *configManagerServiceClient) GetConfigForVizier(ctx context.Context, in *ConfigForVizierRequest, opts ...grpc.CallOption) (*ConfigForVizierResponse, error) {
	out := new(ConfigForVizierResponse)
	err := c.cc.Invoke(ctx, "/px.services.ConfigManagerService/GetConfigForVizier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *configManagerServiceClient) GetConfigForOperator(ctx context.Context, in *ConfigForOperatorRequest, opts ...grpc.CallOption) (*ConfigForOperatorResponse, error) {
	out := new(ConfigForOperatorResponse)
	err := c.cc.Invoke(ctx, "/px.services.ConfigManagerService/GetConfigForOperator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConfigManagerServiceServer is the server API for ConfigManagerService service.
type ConfigManagerServiceServer interface {
	GetConfigForVizier(context.Context, *ConfigForVizierRequest) (*ConfigForVizierResponse, error)
	GetConfigForOperator(context.Context, *ConfigForOperatorRequest) (*ConfigForOperatorResponse, error)
}

// UnimplementedConfigManagerServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConfigManagerServiceServer struct {
}

func (*UnimplementedConfigManagerServiceServer) GetConfigForVizier(ctx context.Context, req *ConfigForVizierRequest) (*ConfigForVizierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigForVizier not implemented")
}
func (*UnimplementedConfigManagerServiceServer) GetConfigForOperator(ctx context.Context, req *ConfigForOperatorRequest) (*ConfigForOperatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigForOperator not implemented")
}

func RegisterConfigManagerServiceServer(s *grpc.Server, srv ConfigManagerServiceServer) {
	s.RegisterService(&_ConfigManagerService_serviceDesc, srv)
}

func _ConfigManagerService_GetConfigForVizier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigForVizierRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManagerServiceServer).GetConfigForVizier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/px.services.ConfigManagerService/GetConfigForVizier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManagerServiceServer).GetConfigForVizier(ctx, req.(*ConfigForVizierRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConfigManagerService_GetConfigForOperator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfigForOperatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConfigManagerServiceServer).GetConfigForOperator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/px.services.ConfigManagerService/GetConfigForOperator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConfigManagerServiceServer).GetConfigForOperator(ctx, req.(*ConfigForOperatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConfigManagerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "px.services.ConfigManagerService",
	HandlerType: (*ConfigManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConfigForVizier",
			Handler:    _ConfigManagerService_GetConfigForVizier_Handler,
		},
		{
			MethodName: "GetConfigForOperator",
			Handler:    _ConfigManagerService_GetConfigForOperator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/cloud/config_manager/configmanagerpb/service.proto",
}

func (m *ConfigForVizierRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigForVizierRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigForVizierRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.K8sVersion) > 0 {
		i -= len(m.K8sVersion)
		copy(dAtA[i:], m.K8sVersion)
		i = encodeVarintService(dAtA, i, uint64(len(m.K8sVersion)))
		i--
		dAtA[i] = 0x1a
	}
	if m.VzSpec != nil {
		{
			size, err := m.VzSpec.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintService(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Namespace) > 0 {
		i -= len(m.Namespace)
		copy(dAtA[i:], m.Namespace)
		i = encodeVarintService(dAtA, i, uint64(len(m.Namespace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConfigForVizierResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigForVizierResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigForVizierResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SentryDSN) > 0 {
		i -= len(m.SentryDSN)
		copy(dAtA[i:], m.SentryDSN)
		i = encodeVarintService(dAtA, i, uint64(len(m.SentryDSN)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.NameToYamlContent) > 0 {
		for k := range m.NameToYamlContent {
			v := m.NameToYamlContent[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintService(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintService(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintService(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ConfigForOperatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigForOperatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigForOperatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ConfigForOperatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConfigForOperatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConfigForOperatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SentryOperatorDSN) > 0 {
		i -= len(m.SentryOperatorDSN)
		copy(dAtA[i:], m.SentryOperatorDSN)
		i = encodeVarintService(dAtA, i, uint64(len(m.SentryOperatorDSN)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintService(dAtA []byte, offset int, v uint64) int {
	offset -= sovService(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ConfigForVizierRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Namespace)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	if m.VzSpec != nil {
		l = m.VzSpec.Size()
		n += 1 + l + sovService(uint64(l))
	}
	l = len(m.K8sVersion)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *ConfigForVizierResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.NameToYamlContent) > 0 {
		for k, v := range m.NameToYamlContent {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovService(uint64(len(k))) + 1 + len(v) + sovService(uint64(len(v)))
			n += mapEntrySize + 1 + sovService(uint64(mapEntrySize))
		}
	}
	l = len(m.SentryDSN)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func (m *ConfigForOperatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ConfigForOperatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SentryOperatorDSN)
	if l > 0 {
		n += 1 + l + sovService(uint64(l))
	}
	return n
}

func sovService(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozService(x uint64) (n int) {
	return sovService(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *ConfigForVizierRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConfigForVizierRequest{`,
		`Namespace:` + fmt.Sprintf("%v", this.Namespace) + `,`,
		`VzSpec:` + strings.Replace(fmt.Sprintf("%v", this.VzSpec), "VizierSpec", "vizierconfigpb.VizierSpec", 1) + `,`,
		`K8sVersion:` + fmt.Sprintf("%v", this.K8sVersion) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConfigForVizierResponse) String() string {
	if this == nil {
		return "nil"
	}
	keysForNameToYamlContent := make([]string, 0, len(this.NameToYamlContent))
	for k, _ := range this.NameToYamlContent {
		keysForNameToYamlContent = append(keysForNameToYamlContent, k)
	}
	github_com_gogo_protobuf_sortkeys.Strings(keysForNameToYamlContent)
	mapStringForNameToYamlContent := "map[string]string{"
	for _, k := range keysForNameToYamlContent {
		mapStringForNameToYamlContent += fmt.Sprintf("%v: %v,", k, this.NameToYamlContent[k])
	}
	mapStringForNameToYamlContent += "}"
	s := strings.Join([]string{`&ConfigForVizierResponse{`,
		`NameToYamlContent:` + mapStringForNameToYamlContent + `,`,
		`SentryDSN:` + fmt.Sprintf("%v", this.SentryDSN) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ConfigForOperatorRequest) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConfigForOperatorRequest{`,
		`}`,
	}, "")
	return s
}
func (this *ConfigForOperatorResponse) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ConfigForOperatorResponse{`,
		`SentryOperatorDSN:` + fmt.Sprintf("%v", this.SentryOperatorDSN) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringService(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *ConfigForVizierRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigForVizierRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigForVizierRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Namespace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Namespace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VzSpec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.VzSpec == nil {
				m.VzSpec = &vizierconfigpb.VizierSpec{}
			}
			if err := m.VzSpec.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field K8sVersion", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.K8sVersion = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigForVizierResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigForVizierResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigForVizierResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NameToYamlContent", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NameToYamlContent == nil {
				m.NameToYamlContent = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowService
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowService
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthService
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthService
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowService
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthService
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthService
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipService(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthService
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.NameToYamlContent[mapkey] = mapvalue
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentryDSN", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SentryDSN = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigForOperatorRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigForOperatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigForOperatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ConfigForOperatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowService
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ConfigForOperatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConfigForOperatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SentryOperatorDSN", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowService
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthService
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthService
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SentryOperatorDSN = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipService(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthService
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipService(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowService
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowService
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthService
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupService
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthService
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthService        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowService          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupService = fmt.Errorf("proto: unexpected end of group")
)