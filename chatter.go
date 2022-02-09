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

	// raddy, _ := net.ResolveTCPAddr("tcp", "0.0.0.0:20012")
	// conn, _ := net.DialTCP("tcp", nil, raddy)
	// fmt.Println(conn.LocalAddr().String())

	// if err == nil {
	// }

	// defer conn.Close()
	defer listen.Close()

	for {
		inConn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println("Error connecting:", err.Error())
			return
		}
		fmt.Println("Connection established with adress: ", inConn.RemoteAddr().String())

		go write(inConn, "Player 1")
	}

	// go read(conn)

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
		fmt.Println(msg[1], ": ", msg[0])
		time.Sleep(1000 * time.Millisecond)
	}
}

func write(conn *net.TCPConn, user string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

		fmt.Println("skriv en melding")
		fmt.Print(user, ": ")
		msg := scanner.Text()

		msg += "\\x00" + user

		time.Sleep(10 * time.Millisecond)
		conn.Write([]byte(msg))
		time.Sleep(1000 * time.Millisecond)
	}
}
