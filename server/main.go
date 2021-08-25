package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/justincampbell/url-shortener-go/urlstore"
	pb "github.com/justincampbell/url-shortener-grpc/server/shortener"
	"google.golang.org/grpc"
)

var port = "1901"

type server struct {
	pb.UnimplementedShortenerServiceServer
}

var store *urlstore.URLStore

func init() {
	if v := os.Getenv("PORT"); v != "" {
		port = v
	}
}

func main() {
	store = urlstore.NewURLStore()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
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
