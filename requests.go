package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

// Client is a struct that implements
// client side logic
type Client struct {
	remoteHost string
}

func hasSlash(host string) bool {
	if host[len(host)-1] == '/' {
		return true
	}

	return false
}

// NewClient - creates new instance of the Client object
func NewClient(host string) (c Client, err error) {
	_, err = url.ParseRequestURI(host)
	if err != nil {
		return Client{}, err
	}

	if hasSlash(host) {
		return Client{}, errors.New("invalid URI is specified")
	}

	return Client{remoteHost: host}, nil
}

// ID holds information about users ID
// Used to deserialize json string which holds
// users id
type ID struct {
	Val int `json:"id"`
}

func canBeRegistered(u *User) bool {
	if u.Name == "" || u.Balance == 0 {
		return false
	}

	return true
}

func (c *Client) regURL() string {
	return c.remoteHost + "/user/register"
}

// Register - sends register POST message and gets
// id of new user
func (c *Client) Register(u *User) (id int, err error) {
	if !canBeRegistered(u) {
		return 0, errors.New("cannot create empty user")
	}

	body := map[string]interface{}{
		"name":    u.Name,
		"balance": u.Balance,
	}

	encodedBody, err := json.Marshal(body)
	if err != nil {
		return 0, err
	}
	resp, err := http.Post(c.regURL(), "application/json", bytes.NewBuffer(encodedBody))
	if err != nil {
		return 0, err
	}

	var decoded ID

	err = json.NewDecoder(resp.Body).Decode(&decoded)
	if err != nil {
		return 0, err
	}

	id = decoded.Val

	return id, nil
}

// User holds information about users account
type User struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
}

func (c *Client) opURL(id int, operation string) string {
	return fmt.Sprintf(c.remoteHost+"/user/%d/%s", id, operation)
}

func canModifyBalance(id, points int) bool {
	if id == 0 || points == 0 {
		return false
	}

	return true
}

// Take - takes points from users balance
func (c *Client) Take(id, points int) (gamer User, err error) {
	if !canModifyBalance(id, points) {
		return User{}, errors.New("Cannot modify users balance")
	}
	body := map[string]interface{}{
		"points": points,
	}

	encodedBody, err := json.Marshal(body)
	if err != nil {
		return User{}, err
	}

	resp, err := http.Post(c.opURL(id, "take"), "application/json", bytes.NewBuffer(encodedBody))
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
func (c *Client) Fund(id, points int) (u User, err error) {
	if !canModifyBalance(id, points) {
		return User{}, errors.New("Cannot modify users balance")
	}
	body := map[string]interface{}{
		"points": points,
	}

	encodedBody, err := json.Marshal(body)
	if err != nil {
		return User{}, nil
	}

	resp, err := http.Post(c.opURL(id, "fund"), "application/json", bytes.NewBuffer(encodedBody))
	if err != nil {
		return User{}, err
	}

	err = json.NewDecoder(resp.Body).Decode(&u)
	if err != nil {
		return User{}, err
	}

	return u, nil
}
