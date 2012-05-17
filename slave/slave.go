package main

import (
	"fmt"
	"log"
	"net"
	"netchan"
	"time"
)

func main() {
	e := netchan.NewExporter()
	ch := make(chan string)
	err := e.Export("pingme", ch, netchan.Send)
	if err != nil {
		log.Fatal(err)
	}
	allDone := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- fmt.Sprintf("ping %d", i)
		}
		allDone <- true
	}()
	fmt.Printf("listening on localhost:8001\n")
	listener, err := net.Listen("tcp", "localhost:8001")
	if err != nil {
		log.Fatal(err)
	}
	go e.Serve(listener)
	_ = <-allDone
	e.Drain(0)
	e.Hangup("pingme")
	time.Sleep(time.Second)
}
