package main

import (
	"fmt"
	"net"
)

const port = "8080"
const stdresponse = `
HTTP/1.1 200 OK
Server: gohttps
Content-Type: text; charset=utf-8

This is an awesome response!
`

func httpServer(c net.Conn) {
	buf := make([]byte, 5120)
	nr, _ := c.Read(buf)
	fmt.Println(string(buf[:nr]))
	c.Write([]byte(stdresponse))
	c.Close()
}

func main() {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		panic(err)
	}
	fmt.Println("started listening on", port)
	for {
		fd, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		go httpServer(fd)
	}
}
