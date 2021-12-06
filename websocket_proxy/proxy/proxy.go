package proxy

import (
	"fmt"
	"net/http"
	"net/url"

	"k8s.io/apimachinery/pkg/util/httpstream"
	"k8s.io/apimachinery/pkg/util/proxy"
)

type Server struct {
	upgradeAwareHandler *proxy.UpgradeAwareHandler
}

type responder struct{}

func (r *responder) Error(w http.ResponseWriter, req *http.Request, err error) {
	fmt.Printf("error responder: %v", err)
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(err.Error()))
}

func New() *Server {
	return &Server{
		upgradeAwareHandler: proxy.NewUpgradeAwareHandler(
			&url.URL{
				Host: "127.0.0.1:8080",
			},
			proxy.MirrorRequest,
			false,
			false,
			&responder{},
		),
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if httpstream.IsUpgradeRequest(r) {
		fmt.Println("get upgrade request")
		s.upgradeAwareHandler.ServeHTTP(w, r)
		return
	}

	fmt.Printf("get request: %s\n", r.URL.String())
	fmt.Printf("scheme: %s\n", r.URL.Scheme)
	for k, v := range r.Header {
		fmt.Printf("%s:%s\n", k, v)
	}
}
