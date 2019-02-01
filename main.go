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

	id, err := register()
	if err != nil {
		fmt.Println("Couldn't register new user")
		os.Exit(1)
	}
	fmt.Println(id)

	gamer, err := take()
	if err != nil {
		fmt.Println("Failed taking points from users balance")
		os.Exit(1)
	}
	fmt.Println(gamer)
	gamer, err = fund()
	if err != nil {
		fmt.Println("Failed funding points to users balance")
		os.Exit(1)
	}
	fmt.Println(gamer)

}
