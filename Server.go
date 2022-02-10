package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

func main() {

	laddy, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:20013")
	listen, _ := net.ListenTCP("tcp", laddy)

	conn, _ := listen.AcceptTCP()

	defer listen.Close()

	go write(conn, "player 1")
	go read(conn)

	for {
	}
}

func read(inconn *net.TCPConn) {
	for {
		buffer := make([]byte, 1024)
		n, err := inconn.Read(buffer)
		if err != nil {
			fmt.Println("Connection lost")
			break
		}
		//fmt.Println(n, "bytes recieved. Local:", conn.LocalAddr().String(), " Remote:", conn.RemoteAddr().String())
		msg := strings.Split(string(buffer[:n]), "\\x00")
		fmt.Println()
		fmt.Println(msg[1], ": ", msg[0])
		time.Sleep(1000 * time.Millisecond)
	}
}

func write(conn *net.TCPConn, user string) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Me: ")
	for scanner.Scan() {

		// fmt.Println("skriv en melding")
		fmt.Print("Me: ")
		msg := scanner.Text()

		msg += "\\x00" + user

		time.Sleep(10 * time.Millisecond)
		conn.Write([]byte(msg))
		time.Sleep(1000 * time.Millisecond)
	}
}
