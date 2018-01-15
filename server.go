// Initial code from https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/

package main

import (
	"flag"
	"fmt"
	"net"
	"os"
)

/* A Simple function to verify error */
func CheckError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(0)
	}
}

func main() {

	host := flag.String("host", "", "hostname:port")
	flag.Parse()

	/* Lets prepare a address at any address at port 1700*/
	ServerAddr, err := net.ResolveUDPAddr("udp", ":1700")
	CheckError(err)

	/* Now listen at selected port */
	ServerConn, err := net.ListenUDP("udp", ServerAddr)
	CheckError(err)
	defer ServerConn.Close()

	conn, err := net.Dial("udp", *host)
	if err != nil {
		fmt.Printf("%v", err)
	}

	buf := make([]byte, 1024)

	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		fmt.Fprintf(conn, string(buf[0:n]))
		fmt.Println("Received ", string(buf[0:n]), " from ", addr)

		if err != nil {
			fmt.Println("Error: ", err)
		}
	}
}
