package main

import (
	"github.com/nailcui/cloud-tool/routers"
	"net/http"
)

func main() {
	routersInit := routers.InitRouter()
	server := &http.Server{
		Addr:           ":80",
		Handler:        routersInit,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
