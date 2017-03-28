// Code generated by protoc-gen-go.
// source: base.proto
// DO NOT EDIT!

/*
Package base is a generated protocol buffer package.

It is generated from these files:
	base.proto

It has these top-level messages:
	ActionRequest
	ActionResponse
*/
package base

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ActionRequest struct {
	Properties map[string]string `protobuf:"bytes,1,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ActionRequest) Reset()                    { *m = ActionRequest{} }
func (m *ActionRequest) String() string            { return proto.CompactTextString(m) }
func (*ActionRequest) ProtoMessage()               {}
func (*ActionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActionRequest) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

type ActionResponse struct {
	Properties map[string]string `protobuf:"bytes,1,rep,name=properties" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ActionResponse) Reset()                    { *m = ActionResponse{} }
func (m *ActionResponse) String() string            { return proto.CompactTextString(m) }
func (*ActionResponse) ProtoMessage()               {}
func (*ActionResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ActionResponse) GetProperties() map[string]string {
	if m != nil {
		return m.Properties
	}
	return nil
}

func init() {
	proto.RegisterType((*ActionRequest)(nil), "ActionRequest")
	proto.RegisterType((*ActionResponse)(nil), "ActionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Publisher API

type Publisher interface {
	Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error
}

type publisher struct {
	c     client.Client
	topic string
}

func (p *publisher) Publish(ctx context.Context, msg interface{}, opts ...client.PublishOption) error {
	return p.c.Publish(ctx, p.c.NewPublication(p.topic, msg), opts...)
}

func NewPublisher(topic string, c client.Client) Publisher {
	if c == nil {
		c = client.NewClient()
	}
	return &publisher{c, topic}
}

// Subscriber API

func RegisterSubscriber(topic string, s server.Server, h interface{}, opts ...server.SubscriberOption) error {
	return s.Subscribe(s.NewSubscriber(topic, h, opts...))
}

// Client API for Service service

type ServiceClient interface {
	RunAction(ctx context.Context, in *ActionRequest, opts ...client.CallOption) (*ActionResponse, error)
}

type serviceClient struct {
	c           client.Client
	serviceName string
}

func NewServiceClient(serviceName string, c client.Client) ServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "service"
	}
	return &serviceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *serviceClient) RunAction(ctx context.Context, in *ActionRequest, opts ...client.CallOption) (*ActionResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Service.RunAction", in)
	out := new(ActionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	RunAction(context.Context, *ActionRequest, *ActionResponse) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Service{hdlr}, opts...))
}

type Service struct {
	ServiceHandler
}

func (h *Service) RunAction(ctx context.Context, in *ActionRequest, out *ActionResponse) error {
	return h.ServiceHandler.RunAction(ctx, in, out)
}

func init() { proto.RegisterFile("base.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 193 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0x4a, 0x2c, 0x4e,
	0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0xea, 0x63, 0xe4, 0xe2, 0x75, 0x4c, 0x2e, 0xc9, 0xcc,
	0xcf, 0x0b, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xb2, 0xe3, 0xe2, 0x2a, 0x28, 0xca, 0x2f,
	0x48, 0x2d, 0x2a, 0xc9, 0x4c, 0x2d, 0x96, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0x92, 0xd3, 0x43,
	0x51, 0xa3, 0x17, 0x00, 0x57, 0xe0, 0x9a, 0x57, 0x52, 0x54, 0x19, 0x84, 0xa4, 0x43, 0xca, 0x96,
	0x8b, 0x1f, 0x4d, 0x5a, 0x48, 0x80, 0x8b, 0x39, 0x3b, 0xb5, 0x52, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xc4, 0x14, 0x12, 0xe1, 0x62, 0x2d, 0x4b, 0xcc, 0x29, 0x4d, 0x95, 0x60, 0x02, 0x8b,
	0x41, 0x38, 0x56, 0x4c, 0x16, 0x8c, 0x4a, 0x13, 0x18, 0xb9, 0xf8, 0x60, 0x96, 0x15, 0x17, 0xe4,
	0xe7, 0x15, 0xa7, 0x0a, 0xd9, 0x63, 0x71, 0x91, 0xbc, 0x1e, 0xaa, 0x22, 0x1a, 0x3a, 0xc9, 0xc8,
	0x92, 0x8b, 0x3d, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x8f, 0x8b, 0x33, 0xa8, 0x34,
	0x0f, 0x62, 0xb5, 0x10, 0x1f, 0x6a, 0xa8, 0x48, 0xf1, 0xa3, 0xb9, 0x49, 0x89, 0x21, 0x89, 0x0d,
	0x1c, 0xca, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x15, 0xf4, 0xbd, 0x56, 0x73, 0x01, 0x00,
	0x00,
}