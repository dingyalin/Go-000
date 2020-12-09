package main

import (
	"context"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"log"
)

func newServer(addr string) *http.Server {
	log.Printf("server addr: %s", addr)
	server := &http.Server{
		Addr:    addr,
		Handler: nil,
	}

	return server
}

func listenAndServe(server *http.Server) {
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("listenAndServe err: %s", err)
	}
}

func listenSignal(server *http.Server) {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	sig := <-signalChan
	log.Printf("signal: %s", sig)
	timeoutCtx, cancal := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancal()
	server.Shutdown(timeoutCtx)
}

func main() {
	server := newServer("127.0.0.1:9090")
	go listenAndServe(server)
	listenSignal(server)
}
