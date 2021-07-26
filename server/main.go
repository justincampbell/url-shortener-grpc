package main

import (
	"log"
	"net"

	"github.com/justincampbell/url-shortener-go/urlstore"
	pb "github.com/justincampbell/url-shortener-grpc/server/shortener"
	"google.golang.org/grpc"
)

const port = ":1901"

type server struct {
	pb.UnimplementedShortenerServiceServer
}

var store *urlstore.URLStore

func main() {
	store = urlstore.NewURLStore()

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterShortenerServiceServer(s, &server{})

	log.Printf("gRPC %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
