// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: location/location.proto

package location

import (
	fmt "fmt"
	common "github.com/cargod-bj/b2c-proto-common/common"
	store "github.com/cargod-bj/b2c-store-api/store"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/anypb"
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

// Api Endpoints for Location service

func NewLocationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Location service

type LocationService interface {
	Add(ctx context.Context, in *LocationDto, opts ...client.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *LocationDto, opts ...client.CallOption) (*common.Response, error)
	Update(ctx context.Context, in *LocationDto, opts ...client.CallOption) (*common.Response, error)
	List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error)
	// 获取ids获取一组store.
	//  Data = common.Page {
	//    List = List<LocationDto>
	//  }
	ListByIds(ctx context.Context, in *store.IdsDto, opts ...client.CallOption) (*common.Response, error)
}

type locationService struct {
	c    client.Client
	name string
}

func NewLocationService(name string, c client.Client) LocationService {
	return &locationService{
		c:    c,
		name: name,
	}
}

func (c *locationService) Add(ctx context.Context, in *LocationDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Location.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) Delete(ctx context.Context, in *LocationDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Location.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) Update(ctx context.Context, in *LocationDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Location.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) List(ctx context.Context, in *common.Page, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Location.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationService) ListByIds(ctx context.Context, in *store.IdsDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Location.ListByIds", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Location service

type LocationHandler interface {
	Add(context.Context, *LocationDto, *common.Response) error
	Delete(context.Context, *LocationDto, *common.Response) error
	Update(context.Context, *LocationDto, *common.Response) error
	List(context.Context, *common.Page, *common.Response) error
	// 获取ids获取一组store.
	//  Data = common.Page {
	//    List = List<LocationDto>
	//  }
	ListByIds(context.Context, *store.IdsDto, *common.Response) error
}

func RegisterLocationHandler(s server.Server, hdlr LocationHandler, opts ...server.HandlerOption) error {
	type location interface {
		Add(ctx context.Context, in *LocationDto, out *common.Response) error
		Delete(ctx context.Context, in *LocationDto, out *common.Response) error
		Update(ctx context.Context, in *LocationDto, out *common.Response) error
		List(ctx context.Context, in *common.Page, out *common.Response) error
		ListByIds(ctx context.Context, in *store.IdsDto, out *common.Response) error
	}
	type Location struct {
		location
	}
	h := &locationHandler{hdlr}
	return s.Handle(s.NewHandler(&Location{h}, opts...))
}

type locationHandler struct {
	LocationHandler
}

func (h *locationHandler) Add(ctx context.Context, in *LocationDto, out *common.Response) error {
	return h.LocationHandler.Add(ctx, in, out)
}

func (h *locationHandler) Delete(ctx context.Context, in *LocationDto, out *common.Response) error {
	return h.LocationHandler.Delete(ctx, in, out)
}

func (h *locationHandler) Update(ctx context.Context, in *LocationDto, out *common.Response) error {
	return h.LocationHandler.Update(ctx, in, out)
}

func (h *locationHandler) List(ctx context.Context, in *common.Page, out *common.Response) error {
	return h.LocationHandler.List(ctx, in, out)
}

func (h *locationHandler) ListByIds(ctx context.Context, in *store.IdsDto, out *common.Response) error {
	return h.LocationHandler.ListByIds(ctx, in, out)
}
