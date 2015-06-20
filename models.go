package main

//Copyright 2015 Graham Green

import "github.com/nu7hatch/gouuid"

//Donation an instance of an action on an acctoun
type Donation struct {
	ID   uuid.UUID `json:"id"`
	From uuid.UUID `json:"from"`
	To   uuid.UUID `json:"to"`
}

//Account an account; can send donations by default
//must have an accociated action to receive actions
type Account struct {
	ID     uuid.UUID `json:"id"`
	Action uuid.UUID `json:"action"`
	Name   string    `json:"name"`
	Stuff  string    `json:"stuff"`
}

//Action a action associated w/ an account
type Action struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Thingy string    `json:"thingy"`
}
