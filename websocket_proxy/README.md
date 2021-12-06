# websocket proxy

The demo contains a websocket server and a client proxy which is responsible to transmit messages received from clients to the websocket server.

## wsserver
wsserver contains a implementation websocket based on gorilla/websocket. It will read message from client and print it on stdout. Also, it will continuously send message "hello" to the client per 5 second.  
Listening on 0.0.0.0:8080

## client proxy
It's implemented on the base of `UpgradeAwareHander` in `k8s.io/apimachinery`.
Listenning on 0.0.0.0:6443

## How to run
First, start the websocket server and the client proxy.
```bash
nohup go run ./wsserver/cmd/main.go &
nohup go run ./proxy/cmd/main.go &
```

Second, we need to download the [wscat image](https://registry.hub.docker.com/r/joshgubler/wscat)
```bash
docker pull joshgubler/wscat
alias wscat='docker run -it --rm --net=host joshgubler/wscat'
```
Then run the wscat command to send the websocket request.
```bash
wscat -n -c ws://127.0.0.1:6443
```


