package shortener

import (
	"context"

	__ "github.com/gihpee/linkShortener/pkg/api"
)

type GRPCServer interface {
	Short(context.Context, *__.UrlRequest) (*__.UrlResponse, error)
	Expand(context.Context, *__.UrlRequest) (*__.UrlResponse, error)
}
