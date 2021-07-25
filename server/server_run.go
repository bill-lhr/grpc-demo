package server

import (
	"fmt"
	"github.com/bill-lhr/grpc-demo/proto/order"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = 8090
)

type Server struct {
	order.UnimplementedOrderServiceServer
}

func (srv *Server) Run() {
	s := grpc.NewServer()
	order.RegisterOrderServiceServer(s, srv)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", port))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
