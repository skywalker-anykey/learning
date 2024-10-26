package main

import (
	"fmt"
	"log"
	"net/rpc"
)

const (
	addr    = ":12345"
	network = "tcp4"
)

func main() {
	client, err := rpc.DialHTTP(network, addr)
	if err != nil {
		log.Fatal(err)
	}
	var resp string
	err = client.Call("Server.Time", "unknow message", &resp)
	if err != nil {
		log.Println("ошибка: ", err)
	}
	fmt.Println("time: ", resp)
	err = client.Call("Server.Time", "time", &resp)
	if err != nil {
		log.Println("ошибка: ", err)
	}
	fmt.Println("time: ", resp)
}
