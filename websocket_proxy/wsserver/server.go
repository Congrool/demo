package wsserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"k8s.io/apimachinery/pkg/util/httpstream"
)

var connid int64

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if httpstream.IsUpgradeRequest(r) {
		fmt.Printf("get upgrade request: %d\n", connid)
		s.serveUpgrade(w, r)
		return
	}
	fmt.Println("get normal request")
}

func (s *Server) serveUpgrade(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		HandshakeTimeout: 10 * time.Second,
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("failed to upgrade request")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context, cancel context.CancelFunc) {
		id := connid

		defer func() {
			fmt.Printf("stop read from %d\n", id)
			cancel()
		}()

		for {
			msgtype, buf, err := conn.ReadMessage()
			if err != nil {
				fmt.Printf("failed to read message, %v\n", err)
				if err.(*websocket.CloseError).Code >= 1000 {
					return
				}
			}
			if msgtype == websocket.CloseMessage {
				fmt.Printf("get close messge, closing %d\n", id)
				return
			}
			fmt.Printf("get msg from %d: %s ", id, string(buf))
		}
	}(ctx, cancel)

	go func(ctx context.Context) {
		id := connid

		ticker := time.NewTicker(5 * time.Second)
		buf := []byte("hello")

		defer func() {
			fmt.Printf("stop write to %d", id)
			ticker.Stop()
		}()

		for {
			select {
			case <-ticker.C:
				if err := conn.WriteMessage(websocket.TextMessage, buf); err != nil {
					fmt.Printf("failed to send message, %v", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}(ctx)

	connid += 1
}
