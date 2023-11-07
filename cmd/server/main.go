package main

import (
	"flag"
	"log"
	"net"

	__ "github.com/gihpee/linkShortener/pkg/api"
	"github.com/gihpee/linkShortener/pkg/shortener"
	"google.golang.org/grpc"
)

var storageType string

func init() {
	flag.StringVar(&storageType, "storage", "in-memory", "Define a memory type: in-memory or postgres")
	flag.Parse()
}

func main() {
	s := grpc.NewServer()
	var srv shortener.GRPCServer

	switch storageType {
	case "in-memory":
		srv = &shortener.GRPCServer_inmemory{}
	case "postgres":
		srv = &shortener.GRPCServer_postgres{}
	default:
		log.Fatalf("Invalid storage type: %s", storageType)
	}

	__.RegisterLinkShortenerServer(s, srv)

	l, err := net.Listen("tcp", ":9080")
	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
