package main

import (
	"context"
	"github.com/bill-lhr/grpc-demo/proto/order"
	"google.golang.org/grpc"
	"log"
	"time"
)

const (
	address = "127.0.0.1:8090"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)

	c := order.NewOrderServiceClient(conn)
	req := order.OrderRequest{
		OrderId:   "201907310001",
		Timestamp: time.Now().Unix(),
	}
	res, err := c.GetOrderInfo(context.Background(), &req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("resp: %v", res)
}
