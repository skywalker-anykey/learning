package main

import (
	"log"
	"math"
	"net"
	"net/http"
	"net/rpc"
)

type Server int

type Point struct {
	X, Y float64
}

type Points struct {
	A, B Point
}

func (s *Server) Dist(p Points, resp *float64) error {
	*resp = (p.A.X-p.B.X)*(p.A.X-p.B.X) + (p.A.Y-p.B.Y)*(p.A.Y-p.B.Y)
	*resp = math.Sqrt(*resp)
	return nil
}

func main() {
	srv := new(Server)
	_ = rpc.Register(srv)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp4", ":1234")
	if err != nil {
		log.Fatal(err)
	}
	_ = http.Serve(listener, nil)
}
