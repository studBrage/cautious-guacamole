package main

import (
	"fmt"
	"net"
)

func main() {

	laddy, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:20013")
	listen, _ := net.ListenTCP("tcp", laddy)

	inConn, _ := listen.AcceptTCP()

	buffer := make([]byte, 100)
	inConn.Read(buffer)
	fmt.Println(string(buffer))

}
