package main

import (
	"flag"
	"log"
	"fmt"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {	
	port := flag.Uint("port", 5000, "TCP Port Number for Blockchain Server")
	flag.Parse()
	app := NewBlockchainServer(uint16(*port))
	fmt.Printf("Running Blockchain Server on port %d", app.Port())
	app.Run()
}
