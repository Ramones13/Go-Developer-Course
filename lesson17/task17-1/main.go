package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
	log.Println("завершениеработы")

}

func handleConn(c net.Conn) {
	defer c.Close()
	var i int
	for {
		_, err := io.WriteString(c, fmt.Sprintf("%d\n", i))
		if err != nil {
			log.Println(err)
			return
		}
		i++
		time.Sleep(time.Second)
	}
}
