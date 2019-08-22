package main

import (
	"log"
	"net"

)

func main(){
	listener, err := net.Listen("tcp", "localhost:10000")
	if err != nil{
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil{
			log.Print(err)
			continue
		}
		connection := NewConnection(conn)
		log.Print("new connection")
		go handleNewConnection(connection)
	}
}