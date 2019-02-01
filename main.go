package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("the web server address is expected as an argument")
		os.Exit(1)
	}

	gmClient := &Client{os.Args[1]}

	err := test(gmClient)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
