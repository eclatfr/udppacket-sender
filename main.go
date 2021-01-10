package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

func main() {
	var ip = flag.String("ip", "localhost:25565", "please input ip address -ip [ip]")
	flag.Parse()
	fmt.Println("IP:", *ip)

	for {
		p := make([]byte, 8192)
		conn, err := net.Dial("udp", *ip)
		if err != nil {
			fmt.Printf("Some error %v", err)
			break
		}
		fmt.Fprintf(conn, "Hi Server, How are you doing?")
		_, err = bufio.NewReader(conn).Read(p)
		if err == nil {
			fmt.Printf("%s\n", p)
		} else {
			fmt.Printf("Some error %v\n", err)
		}
		conn.Close()
	}

}
