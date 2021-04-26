// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: store.proto

package store

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

// Api Endpoints for Store service

func NewStoreEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Store service

type StoreService interface {
	Add(ctx context.Context, in *StoreDTO, opts ...client.CallOption) (*common.Response, error)
	List(ctx context.Context, in *StoreCondition, opts ...client.CallOption) (*common.Response, error)
	Update(ctx context.Context, in *StoreDTO, opts ...client.CallOption) (*common.Response, error)
	Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	Get(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	// 获取ids获取一组store.
	//  Data = common.Page {
	//    List = List<StoreSimpleDto>
	//  }
	ListSimpleInfoByIds(ctx context.Context, in *IdsDto, opts ...client.CallOption) (*common.Response, error)
	//查询storename以及address
	ListLocation(ctx context.Context, in *IdsDto, opts ...client.CallOption) (*common.Response, error)
	//查询商家禁止时间
	ListStoreInvalidTimeDTO(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error)
	//新增或者修改
	AddOrUpdate(ctx context.Context, in *StoreInvalidTimeDTO, opts ...client.CallOption) (*common.Response, error)
	//生成timeslot redis缓存
	GenTimeSlotCache(ctx context.Context, in *DateList, opts ...client.CallOption) (*common.Response, error)
	//获取 timeslot  列表
	GetTimeSlotCache(ctx context.Context, in *GetTimeSlotCon, opts ...client.CallOption) (*common.Response, error)
}

type storeService struct {
	c    client.Client
	name string
}

func NewStoreService(name string, c client.Client) StoreService {
	return &storeService{
		c:    c,
		name: name,
	}
}

func (c *storeService) Add(ctx context.Context, in *StoreDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.Add", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) List(ctx context.Context, in *StoreCondition, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.List", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Update(ctx context.Context, in *StoreDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.Update", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Delete(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.Delete", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) Get(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.Get", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) ListSimpleInfoByIds(ctx context.Context, in *IdsDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.ListSimpleInfoByIds", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) ListLocation(ctx context.Context, in *IdsDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.ListLocation", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) ListStoreInvalidTimeDTO(ctx context.Context, in *common.IdDto, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.ListStoreInvalidTimeDTO", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) AddOrUpdate(ctx context.Context, in *StoreInvalidTimeDTO, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.AddOrUpdate", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) GenTimeSlotCache(ctx context.Context, in *DateList, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.GenTimeSlotCache", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeService) GetTimeSlotCache(ctx context.Context, in *GetTimeSlotCon, opts ...client.CallOption) (*common.Response, error) {
	req := c.c.NewRequest(c.name, "Store.GetTimeSlotCache", in)
	out := new(common.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Store service

type StoreHandler interface {
	Add(context.Context, *StoreDTO, *common.Response) error
	List(context.Context, *StoreCondition, *common.Response) error
	Update(context.Context, *StoreDTO, *common.Response) error
	Delete(context.Context, *common.IdDto, *common.Response) error
	Get(context.Context, *common.IdDto, *common.Response) error
	// 获取ids获取一组store.
	//  Data = common.Page {
	//    List = List<StoreSimpleDto>
	//  }
	ListSimpleInfoByIds(context.Context, *IdsDto, *common.Response) error
	//查询storename以及address
	ListLocation(context.Context, *IdsDto, *common.Response) error
	//查询商家禁止时间
	ListStoreInvalidTimeDTO(context.Context, *common.IdDto, *common.Response) error
	//新增或者修改
	AddOrUpdate(context.Context, *StoreInvalidTimeDTO, *common.Response) error
	//生成timeslot redis缓存
	GenTimeSlotCache(context.Context, *DateList, *common.Response) error
	//获取 timeslot  列表
	GetTimeSlotCache(context.Context, *GetTimeSlotCon, *common.Response) error
}

func RegisterStoreHandler(s server.Server, hdlr StoreHandler, opts ...server.HandlerOption) error {
	type store interface {
		Add(ctx context.Context, in *StoreDTO, out *common.Response) error
		List(ctx context.Context, in *StoreCondition, out *common.Response) error
		Update(ctx context.Context, in *StoreDTO, out *common.Response) error
		Delete(ctx context.Context, in *common.IdDto, out *common.Response) error
		Get(ctx context.Context, in *common.IdDto, out *common.Response) error
		ListSimpleInfoByIds(ctx context.Context, in *IdsDto, out *common.Response) error
		ListLocation(ctx context.Context, in *IdsDto, out *common.Response) error
		ListStoreInvalidTimeDTO(ctx context.Context, in *common.IdDto, out *common.Response) error
		AddOrUpdate(ctx context.Context, in *StoreInvalidTimeDTO, out *common.Response) error
		GenTimeSlotCache(ctx context.Context, in *DateList, out *common.Response) error
		GetTimeSlotCache(ctx context.Context, in *GetTimeSlotCon, out *common.Response) error
	}
	type Store struct {
		store
	}
	h := &storeHandler{hdlr}
	return s.Handle(s.NewHandler(&Store{h}, opts...))
}

type storeHandler struct {
	StoreHandler
}

func (h *storeHandler) Add(ctx context.Context, in *StoreDTO, out *common.Response) error {
	return h.StoreHandler.Add(ctx, in, out)
}

func (h *storeHandler) List(ctx context.Context, in *StoreCondition, out *common.Response) error {
	return h.StoreHandler.List(ctx, in, out)
}

func (h *storeHandler) Update(ctx context.Context, in *StoreDTO, out *common.Response) error {
	return h.StoreHandler.Update(ctx, in, out)
}

func (h *storeHandler) Delete(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.StoreHandler.Delete(ctx, in, out)
}

func (h *storeHandler) Get(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.StoreHandler.Get(ctx, in, out)
}

func (h *storeHandler) ListSimpleInfoByIds(ctx context.Context, in *IdsDto, out *common.Response) error {
	return h.StoreHandler.ListSimpleInfoByIds(ctx, in, out)
}

func (h *storeHandler) ListLocation(ctx context.Context, in *IdsDto, out *common.Response) error {
	return h.StoreHandler.ListLocation(ctx, in, out)
}

func (h *storeHandler) ListStoreInvalidTimeDTO(ctx context.Context, in *common.IdDto, out *common.Response) error {
	return h.StoreHandler.ListStoreInvalidTimeDTO(ctx, in, out)
}

func (h *storeHandler) AddOrUpdate(ctx context.Context, in *StoreInvalidTimeDTO, out *common.Response) error {
	return h.StoreHandler.AddOrUpdate(ctx, in, out)
}

func (h *storeHandler) GenTimeSlotCache(ctx context.Context, in *DateList, out *common.Response) error {
	return h.StoreHandler.GenTimeSlotCache(ctx, in, out)
}

func (h *storeHandler) GetTimeSlotCache(ctx context.Context, in *GetTimeSlotCon, out *common.Response) error {
	return h.StoreHandler.GetTimeSlotCache(ctx, in, out)
}
