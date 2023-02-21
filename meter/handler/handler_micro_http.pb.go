// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// protoc-gen-go-micro version: v3.10.2
// source: handler.proto

package handler

import (
	context "context"
	v3 "go.unistack.org/micro-server-http/v3"
	codec "go.unistack.org/micro/v3/codec"
	server "go.unistack.org/micro/v3/server"
)

type meterServiceServer struct {
	MeterServiceServer
}

func (h *meterServiceServer) Metrics(ctx context.Context, req *codec.Frame, rsp *codec.Frame) error {
	return h.MeterServiceServer.Metrics(ctx, req, rsp)
}

func RegisterMeterServiceServer(s server.Server, sh MeterServiceServer, opts ...server.HandlerOption) error {
	type meterService interface {
		Metrics(ctx context.Context, req *codec.Frame, rsp *codec.Frame) error
	}
	type MeterService struct {
		meterService
	}
	h := &meterServiceServer{sh}
	var nopts []server.HandlerOption
	nopts = append(nopts, v3.HandlerEndpoints(MeterServiceServerEndpoints))
	return s.Handle(s.NewHandler(&MeterService{h}, append(nopts, opts...)...))
}
