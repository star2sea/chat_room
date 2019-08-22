package main

import "fmt"

func onMessage(c *Connection, msg string){
	fmt.Printf("on message %s", msg)
}
