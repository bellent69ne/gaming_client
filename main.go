package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("the web server address is expected as an argument")
	}

	//	id, err := register()
	//if err != nil {
	//log.Fatal(err)
	//}

	//gamer, err := take()
	//if err != nil {
	//log.Fatal(err)
	//}
	gamer, err := fund()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gamer)

}
