package main

import (
	"flag"
	"fmt"
	"log"
	"netchan"
)

func main() {
	host := flag.String("host", "localhost:8001", "remote host")
	flag.Parse()
	fmt.Printf("attaching to %s\n", host)
	imp, err := netchan.Import("tcp", *host)
	if err != nil {
		log.Fatal(err)
	}
	ch := make(chan string)
	err = imp.ImportNValues("pingme", ch, netchan.Recv, 1, -1)
	if err != nil {
		log.Fatal(err)
	}
	errors := imp.Errors()
loop:
	for {
		select {
		case t, ok := <-ch:
			if !ok {
				fmt.Printf("channel closed\n")
				break loop
			}
			fmt.Printf("got %v\n", t)
		case e := <-errors:
			fmt.Printf("error %v\n", e)
			break loop
		}
	}
}
