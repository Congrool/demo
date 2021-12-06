package main

import (
	"net/http"

	"github.com/Congrool/demo/websocketproxy/proxy"
)

func main() {
	http.ListenAndServe("0.0.0.0:6443", proxy.New())
}
