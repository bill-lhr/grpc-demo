package main

import "github.com/bill-lhr/grpc-demo/server"

func main() {
	s := &server.Server{}
	s.Run()
}
