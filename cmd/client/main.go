package main

import (
	"context"
	_ "flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"

	__ "github.com/gihpee/linkShortener/pkg/api"
	"google.golang.org/grpc"
)

type Result struct {
	Link   string
	Code   string
	Status string
}

func isValidUrl(token string) bool {
	_, err := url.ParseRequestURI(token)
	if err != nil {
		return false
	}
	u, err := url.Parse(token)
	if err != nil || u.Host == "" {
		return false
	}
	return true
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	templ, _ := template.ParseFiles("./templates/index.html")
	result := Result{}

	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := __.NewLinkShortenerClient(conn)

	switch r.Method {
	case "POST":
		url := r.FormValue("p")
		if !isValidUrl(url) {
			result.Status = "Ссылка имеет неправильный формат"
			result.Link = ""
		} else {
			result.Link = url
			res, err := c.Short(context.Background(), &__.UrlRequest{Url: url})
			if err != nil {
				log.Fatal(err)
			}
			result.Code = res.ShortUrl
			result.Status = "Сокращение выполнено"
		}
	case "GET":
		url := r.FormValue("g")
		if url != "" {
			result.Code = url
			res, err := c.Expand(context.Background(), &__.UrlRequest{Url: url})
			if err != nil {
				log.Fatal(err)
			}
			if res.OrigUrl == "" {
				result.Status = "Ссылка не найдена в базе"
			} else {
				result.Link = res.OrigUrl
				result.Status = "Ссылка найдена"
			}
		}
	}
	templ.Execute(w, result)
}

func redirectTo(w http.ResponseWriter, r *http.Request) {
	conn, err := grpc.Dial(":9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := __.NewLinkShortenerClient(conn)

	vars := mux.Vars(r)
	res, err := c.Expand(context.Background(), &__.UrlRequest{Url: vars["key"]})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "<script>location='%s';</script>", res.OrigUrl)
}

func main() {
	/* КЛИЕНТСКАЯ ЧАСТЬ ИЗ КОНСОЛИ, БЕЗ HTML СТРАНИЦЫ
	flag.Parse()
	if flag.NArg() < 2 {
		log.Fatal("not enough args")
	}

	method := flag.Arg(0)

	url := flag.Arg(1)
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	c := __.NewLinkShortenerClient(conn)

	switch method {
	case "POST":
		res, err := c.Short(context.Background(), &__.UrlRequest{Url: url})
		if err != nil {
			log.Fatal(err)
		}

		log.Println(res.ShortUrl)
	case "GET":
		res, err := c.Expand(context.Background(), &__.UrlRequest{Url: url})
		if err != nil {
			log.Fatal(err)
		}

		log.Println(res.OrigUrl)
	default:
		log.Fatalf("Incorrect type of method: %s, must be POST or GET", method)
	}*/
	router := mux.NewRouter()
	router.HandleFunc("/", indexPage)
	router.HandleFunc("/{key}", redirectTo)
	log.Fatal(http.ListenAndServe(":8080", router))
}
