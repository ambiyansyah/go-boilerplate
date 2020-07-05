package server

import "log"

// GRPCServer ...
type GRPCServer interface {
	Start()
}

type grpcServer struct{}

// NewGRPCServer ...
func NewGRPCServer() GRPCServer {
	return &grpcServer{}
}

func (h *grpcServer) Start() {
	log.Printf("GRPCServer Start func called")
}
