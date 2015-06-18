package main

//Copyright 2015 Graham Green

import "github.com/nu7hatch/gouuid"

type Donation struct {
	ID     uuid.UUID `json:"id"`
	From   uuid.UUID `json:"from"`
	To     uuid.UUID `json:"to"`
	Action uuid.UUID `json:"action"`
}

type Account struct {
	ID     uuid.UUID `json:"id"`
	Action uuid.UUID `json:"action"`
	Name   string    `json:"name"`
	Stuff  string    `json:"stuff"`
}

type Action struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Thingy string    `json:"thingy"`
}
