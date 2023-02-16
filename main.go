package main

import (
	"fmt"
	"log"
	"net"

	gh "github.com/sylvainsausse/chess-server/gamehandling"
)

func handle(err error){
	if err != nil {
		log.Fatal(err.Error())
	}
}

func main(){
	server,err := net.Listen("tcp","0.0.0.0:3000")
	handle(err)
	for true {
		cli1, err := server.Accept()
		handle(err)
		fmt.Println("Waiting for second player")
		cli2, err :=  server.Accept()
		handle(err)
		m := gh.NewMatch()
		go m.Start(cli1,cli2)
	}
}