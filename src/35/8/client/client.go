package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Point struct {
	X, Y float64
}

type Points struct {
	A, B Point
}

func main() {
	client, err := rpc.DialHTTP("tcp4", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	var p Points = Points{
		A: Point{X: 1, Y: 1},
		B: Point{X: 4, Y: 5},
	}

	var resp float64

	err = client.Call("Server.Dist", p, &resp)
	if err != nil {
		log.Println("call failed:", err)
	}

	fmt.Println("Dist:", resp)
}
