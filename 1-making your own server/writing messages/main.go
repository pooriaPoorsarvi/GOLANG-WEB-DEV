package main

import (
	"net"
	"log"
	"io"
	"fmt"
)

// To use this code you need to use telnet in windows, that is if you are using windows, and to activate it just
//Go to Activating features in windows, after you activated it you should run in your command line the following command
// telnet localhost 8080

func main(){

	li, err := net.Listen("tcp", ":8080")
	defer li.Close()

	if err != nil{
		log.Fatal(err)
	}


	for {
		conn, err := li.Accept()
		if err != nil{
			log.Println(err)
			continue
		}
		// Here we can see we can use any method that takes in a writer and writes to it in order to write messages
		io.WriteString(conn, "\nHello from TCP server\n")
		fmt.Fprintln(conn, "how you doing ??")
		fmt.Fprintf(conn, "how you doing ?? %s", "babe")

		conn.Close()

	}






}




