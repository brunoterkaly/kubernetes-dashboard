package main

import (
	"net/http"
	"github.com/terkaly/projectweb/myweblib"
)

func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/view", myweblib.Handler)
	server.ListenAndServe()
}
