package shortener

import (
	"context"

	__ "github.com/gihpee/linkShortener/pkg/api"
)

type GRPCServer_inmemory struct {
	storage map[string]string
}

func (s *GRPCServer_inmemory) Short(ctx context.Context, req *__.UrlRequest) (*__.UrlResponse, error) {
	tmp_short_url := short(req.Url)
	if s.storage == nil {
		s.storage = make(map[string]string)
	}
	s.storage[tmp_short_url] = req.Url
	return &__.UrlResponse{ShortUrl: tmp_short_url, OrigUrl: req.Url}, nil
}

func (s *GRPCServer_inmemory) Expand(ctx context.Context, req *__.UrlRequest) (*__.UrlResponse, error) {
	if s.storage == nil {
		s.storage = make(map[string]string)
	}
	return &__.UrlResponse{ShortUrl: req.Url, OrigUrl: s.storage[req.Url]}, nil
}
