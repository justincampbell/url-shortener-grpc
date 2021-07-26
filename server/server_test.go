package main

import (
	"context"
	"os"
	"testing"

	"github.com/justincampbell/url-shortener-go/urlstore"
	pb "github.com/justincampbell/url-shortener-grpc/server/shortener"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	store = urlstore.NewURLStore()
	os.Exit(m.Run())
}

func TestServerShorten(t *testing.T) {
	s := &server{}
	req := &pb.Shorten_Request{Url: "http://example.com"}
	resp, err := s.Shorten(context.Background(), req)
	require.NoError(t, err)
	assert.NotEmpty(t, resp.Token)
}

func TestServerShorten_invalidUrl(t *testing.T) {
	s := &server{}
	req := &pb.Shorten_Request{Url: "invalid"}
	resp, err := s.Shorten(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, resp)
}

func TestServerExpand(t *testing.T) {
	url := "http://example.com"
	token := store.Shorten(url)
	s := &server{}
	req := &pb.Expand_Request{Token: token}
	resp, err := s.Expand(context.Background(), req)
	require.NoError(t, err)
	assert.Equal(t, url, resp.Url)
}

func TestServerExpand_tokenNotFound(t *testing.T) {
	s := &server{}
	req := &pb.Expand_Request{Token: "123"}
	resp, err := s.Expand(context.Background(), req)
	require.Error(t, err)
	assert.Nil(t, resp)
}
