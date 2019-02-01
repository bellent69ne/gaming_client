package main

import (
	"errors"
	"fmt"
)

// test - tests gaming web server for existing operations
// which include registration of the user,
// taking and funding points from users account
func test(cl *Client) error {
	stalker := User{Name: "ni", Balance: 1000}
	id, err := cl.Register(&stalker)
	if err != nil {
		return errors.New("Failed registering new user")
	}
	fmt.Println("id: ", id)

	gmUser, err := cl.Take(id, 300)
	if err != nil {
		return errors.New("Failed taking points from users balance")
	}
	fmt.Println("After take: ", gmUser)

	gmUser, err = cl.Fund(id, 400)
	if err != nil {
		return errors.New("Failed funding points to users balance")
	}
	fmt.Println("After fund: ", gmUser)

	return nil
}
