package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Listening on %s\n", ln.Addr())
	for {
		if conn, err := ln.Accept(); err != nil {
			log.Println(err)
		} else {
			go io.Copy(conn, conn)
		}
	}
}
