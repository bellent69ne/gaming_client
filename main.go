package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("the web server address is expected as an argument")
	}

	//	gmClient := &Client{remoteHost: os.Args[1]}
	gmClient, err := NewClient(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	err = test(&gmClient)
	if err != nil {
		log.Fatal(err)
	}

}
