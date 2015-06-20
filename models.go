package main

//Copyright 2015 Graham Green

//I'm taking out the UUID for now
//import "github.com/nu7hatch/gouuid"

//Donation an instance of an action on an acctoun
type Donation struct {
	ID   string `json:"id"`
	From string `json:"from"`
	To   string `json:"to"`
}

//Account an account; can send donations by default
//must have an accociated action to receive actions
type Account struct {
	ID     string `json:"id"`
	Action string `json:"action"`
	Name   string `json:"name"`
	Stuff  string `json:"stuff"`
}

//Action a action associated w/ an account
type Action struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Thingy string `json:"thingy"`
}
