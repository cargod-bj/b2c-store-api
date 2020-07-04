// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: area/area.proto

package are

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Area service

func NewAreaEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Area service

type AreaService interface {
	Add(ctx context.Context, in *AreaDto, opts ...client.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	Update(ctx context.Context, in *AreaDto, opts ...client.CallOption) (*common.Response, error)
	TopList(ctx context.Context, in *common.LocalDto, opts ...client.CallOption) (*common.Response, error)
	ChildList(ctx context.Context, in *common.IdLocalDTO, opts ...client.CallOption) (*common.Response, error)
}

type areaService struct {
	c    client.Client
	name string
}

func NewAreaService(name string, c client.Client) AreaService {
	return &areaService{
		c:    c,
		name: name,
	}
}

func (c *areaService) Add(ctx context.Context, in *AreaDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Area.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Area.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) Update(ctx context.Context, in *AreaDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Area.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) TopList(ctx context.Context, in *common.LocalDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Area.TopList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *areaService) ChildList(ctx context.Context, in *common.IdLocalDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Area.ChildList", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Area service

type AreaHandler interface {
	Add(context.Context, *AreaDto, *common.Response) error
	Delete(context.Context, *common.IdDto, *common.Response) error
	Update(context.Context, *AreaDto, *common.Response) error
	TopList(context.Context, *common.LocalDto, *common.Response) error
	ChildList(context.Context, *common.IdLocalDTO, *common.Response) error
}

func RegisterAreaHandler(s server.Server, hdlr AreaHandler, opts ...server.HandlerOption) error {
	type area interface {
		Add(ctx context.Context, in *AreaDto, out *common.Response) error
		Delete(ctx context.Context, in *common.IdDto, out *common.Response) error
		Update(ctx context.Context, in *AreaDto, out *common.Response) error
		TopList(ctx context.Context, in *common.LocalDto, out *common.Response) error
		ChildList(ctx context.Context, in *common.IdLocalDTO, out *common.Response) error
	}
	type Area struct {
		area
	}
	h := &areaHandler{hdlr}
	return s.Handle(s.NewHandler(&Area{h}, opts...))
}

type areaHandler struct {
	AreaHandler
}

func (h *areaHandler) Add(ctx context.Context, in *AreaDto, out *common.Response) error {
	return h.AreaHandler.Add(ctx, in, out)
}

func (h *areaHandler) Delete(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.AreaHandler.Delete(ctx, in, out)
}

func (h *areaHandler) Update(ctx context.Context, in *AreaDto, out *common.Response) error {
	return h.AreaHandler.Update(ctx, in, out)
}

func (h *areaHandler) TopList(ctx context.Context, in *common.LocalDto, out *common.Response) error {
	return h.AreaHandler.TopList(ctx, in, out)
}

func (h *areaHandler) ChildList(ctx context.Context, in *common.IdLocalDTO, out *common.Response) error {
	return h.AreaHandler.ChildList(ctx, in, out)
}
