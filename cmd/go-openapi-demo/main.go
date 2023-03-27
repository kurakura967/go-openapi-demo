package main

import (
	"net/http"

	"github.com/kurakura967/go-openapi-demo/api"
)

func main() {
	r := api.NewRouter()

	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: r,
	}
	s.ListenAndServe()
}
