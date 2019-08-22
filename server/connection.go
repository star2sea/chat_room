package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type Connection struct {
	name string
	room *Room
	conn net.Conn
	sendBuffer chan string
}

func NewConnection(conn net.Conn) *Connection{
	c := &Connection{"", nil, conn, make(chan string)}
	return c
}

func (c *Connection) setName(name string){
	c.name = name
}

func (c *Connection) setRoom(room *Room){
	c.room = room
}

func (c *Connection) send(msg string){
	c.sendBuffer <- msg
}

func handleNewConnection(c *Connection){
	go c.handleWrite()
	input := bufio.NewScanner(c.conn)
	for input.Scan(){
		onMessage(c, input.Text())
	}
	c.handleClose()
}

func (c *Connection) handleClose(){
	log.Print("connection closed")
	c.conn.Close()
	if c.room != nil{
		c.room.LeaveRoom(c.name)
	}
}

func (c *Connection) handleWrite(){
	for msg := range c.sendBuffer{
		fmt.Fprintf(c.conn, msg)
	}
}