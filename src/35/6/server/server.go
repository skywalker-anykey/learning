package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

const (
	addr    = ":12345"
	network = "tcp4"
)

type Server int

func (s *Server) Time(msg string, resp *string) error {
	if msg != "time" {
		return errors.New("неизвестная команда")
	}
	*resp = time.Now().String()
	return nil
}

func main() {
	srv := new(Server)
	rpc.Register(srv)
	rpc.HandleHTTP()
	l, e := net.Listen(network, addr)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
