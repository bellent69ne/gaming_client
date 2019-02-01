package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

// ID holds information about users ID
// Used to deserialize json string which holds
// users id
type ID struct {
	Val int `json:"id"`
}

// register - sends register POST message and gets
// id of new user
func register() (id ID, err error) {
	regURL := os.Args[1] + "/user/register"
	body := map[string]interface{}{
		"name":    "stalker",
		"balance": 1000,
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

// take - takes points from users balance
func take() (gamer User, err error) {
	takeURL := os.Args[1] + "/user/1/take"
	body := map[string]interface{}{
		"points": 300,
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

// fund - funds points to users balance
func fund() (gamer User, err error) {
	fundURL := os.Args[1] + "/user/1/fund"
	body := map[string]interface{}{
		"points": 400,
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
