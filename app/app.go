package main

import (
	"log"
	protocol "github.com/rathodc/todo/libservice"
)

func main()  {
	log.Println("Starting App.....")
	protocol.CreateApp()
}