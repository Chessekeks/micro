// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// protoc-gen-go-micro version: v3.4.2
// source: handler.proto

package handler

import (
	context "context"
	api "github.com/unistack-org/micro/v3/api"
	codec "github.com/unistack-org/micro/v3/codec"
	server "github.com/unistack-org/micro/v3/server"
)

type meterServer struct {
	MeterServer
}

func (h *meterServer) Metrics(ctx context.Context, req *codec.Frame, rsp *codec.Frame) error {
	return h.MeterServer.Metrics(ctx, req, rsp)
}

func RegisterMeterServer(s server.Server, sh MeterServer, opts ...server.HandlerOption) error {
	type meter interface {
		Metrics(ctx context.Context, req *codec.Frame, rsp *codec.Frame) error
	}
	type Meter struct {
		meter
	}
	h := &meterServer{sh}
	var nopts []server.HandlerOption
	for _, endpoint := range MeterEndpoints {
		nopts = append(nopts, api.WithEndpoint(&endpoint))
	}
	return s.Handle(s.NewHandler(&Meter{h}, append(nopts, opts...)...))
}
