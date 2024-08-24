package main

import (
	"chat_server/network"
	"log"
)

func init(){
	log.Println("main보다 먼저 실행")
}

func main(){
	n := network.NewServer()
	n.StartServer()
}