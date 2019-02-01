package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Client is a struct that implements
// client side logic
type Client struct {
	remoteHost string
}

// ID holds information about users ID
// Used to deserialize json string which holds
// users id
type ID struct {
	Val int `json:"id"`
}

// Register - sends register POST message and gets
// id of new user
func (cl *Client) Register(gm User) (id ID, err error) {
	regURL := cl.remoteHost + "/user/register"
	body := map[string]interface{}{
		"name":    gm.Name,
		"balance": gm.Balance,
	}

	encodedBody, err := json.Marshal(body)
	if err != nil {
		return ID{}, err
	}
	resp, err := http.Post(regURL, "application/json", bytes.NewBuffer(encodedBody))
	if err != nil {
		return ID{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&id)
	if err != nil {
		return ID{}, err
	}

	return id, nil
}

// User holds information about users account
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func (cl *Client) path(id int, operation string) string {
	return fmt.Sprintf("/user/%d/%s", id, operation)
}

// Take - takes points from users balance
func (cl *Client) Take(id, points int) (gamer User, err error) {
	takeURL := cl.remoteHost + cl.path(id, "take")
	body := map[string]interface{}{
		"points": points,
	}

	encodedBody, err := json.Marshal(body)
	if err != nil {
		return User{}, err
	}

	resp, err := http.Post(takeURL, "application/json", bytes.NewBuffer(encodedBody))
	if err != nil {
		return User{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&gamer)
	if err != nil {
		return User{}, err
	}

	return gamer, nil
}

// Fund - funds points to users balance
func (cl *Client) Fund(id, points int) (gamer User, err error) {
	fundURL := cl.remoteHost + cl.path(id, "fund")
	body := map[string]interface{}{
		"points": points,
	}

	encodedBody, err := json.Marshal(body)
	if err != nil {
		return User{}, nil
	}

	resp, err := http.Post(fundURL, "application/json", bytes.NewBuffer(encodedBody))
	if err != nil {
		return User{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&gamer)
	if err != nil {
		return User{}, err
	}

	return gamer, nil
}
