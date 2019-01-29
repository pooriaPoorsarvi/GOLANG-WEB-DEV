package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"time"
)

func main(){
	li, err := net.Listen("tcp", ":8080")
	defer li.Close()

	if err != nil{
		log.Fatal(err)
	}

	for{
		conn, err := li.Accept()
		if err != nil {
			continue
		}
		// we have to use go routines because the handle function is not closed until the recipient closes the connection
		go handle(conn)

	}


}

func handle(conn net.Conn){

	scanner := bufio.NewScanner(conn)
	// If we didn't set an explicit dead line the code of scanner.scan would have never reached its end
	err := conn.SetDeadline(time.Now().Add(4 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintln(conn, "I Heard You Say", ln)
	}
	defer conn.Close()

	fmt.Println("Code got Here !")

}



