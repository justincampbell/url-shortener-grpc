package main

import (
	"context"
	"log"
	"net/url"

	pb "github.com/justincampbell/url-shortener-grpc/server/shortener"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Shorten responds to gRPC shorten requests.
func (s *server) Shorten(ctx context.Context, request *pb.Shorten_Request) (*pb.Shorten_Response, error) {
	_, err := url.ParseRequestURI(request.Url)
	if err != nil {
		log.Printf("gRPC Shorten\n\trequest=%v\n\terror=invalid argument", request)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid URL: %s", request.Url)
	}

	token := store.Shorten(request.Url)
	response := &pb.Shorten_Response{Token: token}
	log.Printf("gRPC Shorten\n\trequest=%v\n\tresponse=%v", request, response)
	return response, nil
}

// Expand responds to gRPC expand requests.
func (s *server) Expand(ctx context.Context, request *pb.Expand_Request) (*pb.Expand_Response, error) {
	url := store.Expand(request.Token)
	if url == "" {
		log.Printf("gRPC Expand\n\trequest=%v\n\terror=not found", request)
		return nil, status.Errorf(codes.NotFound, "Token not found: %s", request.Token)
	}
	response := &pb.Expand_Response{Url: url}
	log.Printf("gRPC Expand\n\trequest=%v\n\tresponse=%v", request, response)
	return response, nil
}
