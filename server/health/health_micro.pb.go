// Code generated by protoc-gen-go-micro. DO NOT EDIT.
// protoc-gen-go-micro version: v3.4.2
// source: health.proto

package health

import (
	context "context"

	api "go.unistack.org/micro/v3/api"
	codec "go.unistack.org/micro/v3/codec"
)

var (
	HealthName = "Health"

	HealthEndpoints = []api.Endpoint{
		{
			Name:    "Health.Live",
			Path:    []string{"/live"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "Health.Ready",
			Path:    []string{"/ready"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
		{
			Name:    "Health.Version",
			Path:    []string{"/version"},
			Method:  []string{"GET"},
			Handler: "rpc",
		},
	}
)

func NewHealthEndpoints() []api.Endpoint {
	return HealthEndpoints
}

type HealthServer interface {
	Live(ctx context.Context, req *codec.Frame, rsp *codec.Frame) error
	Ready(ctx context.Context, req *codec.Frame, rsp *codec.Frame) error
	Version(ctx context.Context, req *codec.Frame, rsp *codec.Frame) error
}
