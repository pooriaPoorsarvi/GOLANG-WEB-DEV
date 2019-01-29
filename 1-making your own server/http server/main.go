package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
)

func main(){
	li, err  := net.Listen("tcp", ":8080")

	defer li.Close()

	if err != nil{
		log.Fatal(err)
	}


	for {
		conn, err := li.Accept()

		if err != nil {
			log.Println(err)
			continue
		}

		go request(conn)


	}


}



func request(conn net.Conn){

	i := 0

	scann := bufio.NewScanner(conn)

	var m string
	var u string

	for scann.Scan(){
		ln := scann.Text()
		fmt.Println(ln)
		if i == 0 {
			m = strings.Fields(ln)[0]
			u = strings.Fields(ln)[1]
		}
		if ln == ""{
			break
		}
		i++
	}

	respond(m, u, conn)

}


func respond(m string, u string, conn net.Conn){


	if(u == "/007"){
		body := `welcome in 007 !`
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}else{
		body := `Nothing To See Here !`
		fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
		fmt.Fprint(conn, "Content-Type: text/html\r\n")
		fmt.Fprint(conn, "\r\n")
		fmt.Fprint(conn, body)
	}
	defer conn.Close()

}



