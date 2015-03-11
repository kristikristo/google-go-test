package main

import (
	"github.com/go-martini/martini"
	"net/http"
	)

func main() {
  client := martini.Classic()

  client.Get("/", func() string {
    return "Hello world!"
  })

  client.Get("/hello/:name", Auth,  func(params martini.Params) string {
  	return "Hello " + params["name"]
  	})
  client.Run()
}

func Auth(res http.ResponseWriter, req *http.Request) {
	if req.Header.Get("API-KEY") != "kristi" {
		http.Error(res, "Nope", 401);
	}
}