package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"github.com/rarecircles/backend-challenge-go/tokenhandler"
)

var flagHTTPListenAddr = flag.String("http-listen-port", ":8080", "HTTP listen address, if blank will default to ENV PORT")
var flagRPCURL = flag.String("rpc-url", "https://eth-mainnet.alchemyapi.io/v2/", "RPC URL")

func main() {
	flag.Parse()
	httpListenAddr := *flagHTTPListenAddr
	rpcURL := *flagRPCURL
	r := mux.NewRouter()
	r.HandleFunc("/", tokenhandler.TokenHandler)
	log.Fatal(http.ListenAndServe(httpListenAddr, r))

	zlog.Info("Running Challenge",
		zap.String("httpL_listen_addr", httpListenAddr),
		zap.String("rpc_url", rpcURL),
	)
}