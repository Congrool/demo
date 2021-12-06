package main

import (
	"net/http"

	"github.com/Congrool/demo/websocketproxy/wsserver"
)

func main() {
	http.ListenAndServe("127.0.0.1:8080", &wsserver.Server{})
}
