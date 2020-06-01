package main

import (
	"fmt"
	"net"
	"os"

	"github.com/leegeobuk/GoServer/GoServer/env"
	"github.com/leegeobuk/GoServer/GoServer/pb/currency"
	"github.com/leegeobuk/GoServer/GoServer/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	gs := grpc.NewServer()
	cs := server.NewCurrency()

	currency.RegisterCurrencyServer(gs, cs)
	reflection.Register(gs)

	l, err := net.Listen("tcp", env.Addr)
	if err != nil {
		fmt.Printf("unable to listen %w", err)
		os.Exit(1)
	}

	gs.Serve(l)
}
