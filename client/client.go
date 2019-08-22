package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"log"
)

func main(){
	sendBuffer := make(chan string)

	serverAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:10000")
	conn, err := net.DialTCP("tcp", nil, serverAddr)

	if err != nil{
		log.Fatal(err)
	}

	defer conn.Close()

	go handleRead(conn)

	go handleWrite(conn, sendBuffer)

	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil{
			break
		}
		fmt.Printf("send %s", input)
		sendBuffer <- input
	}
	close(sendBuffer)
}


func onMessage(c net.Conn, msg string){
	fmt.Printf("on message %s", msg)
}

func handleRead(c net.Conn){
	input := bufio.NewScanner(c)
	for input.Scan(){
		onMessage(c, input.Text())
	}
	fmt.Printf("handle read over")
}

func handleWrite(c net.Conn, ch chan string){
	for msg := range ch{
		log.Print("start send")
		fmt.Fprintf(c, msg)
	}
	fmt.Printf("handle write over")
}