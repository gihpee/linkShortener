package shortener

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	__ "github.com/gihpee/linkShortener/pkg/api"
)

type GRPCServer_postgres struct {
}

const connStr = "postgres://root:rootroot@db:5432/postgres?sslmode=disable"

func (s *GRPCServer_postgres) Short(ctx context.Context, req *__.UrlRequest) (*__.UrlResponse, error) {
	tmp_short_url := short(req.Url)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("create table if not exists links(link VARCHAR, short VARCHAR)")
	if err != nil {
		panic(err)
	}

	res, err := db.Exec("insert into links (link, short) values ($1, $2)", req.Url, tmp_short_url)
	if err != nil {
		log.Fatal(err)
	}
	_ = res

	return &__.UrlResponse{ShortUrl: tmp_short_url, OrigUrl: req.Url}, nil
}

func (s *GRPCServer_postgres) Expand(ctx context.Context, req *__.UrlRequest) (*__.UrlResponse, error) {
	var orig_url string

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("create table if not exists links(link VARCHAR, short VARCHAR)")
	if err != nil {
		panic(err)
	}

	rows := db.QueryRow("select link from links where short=$1 limit 1", req.Url)
	rows.Scan(&orig_url)

	return &__.UrlResponse{ShortUrl: req.Url, OrigUrl: orig_url}, nil
}
