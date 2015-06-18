package main

//Copyright 2015 Graham Green

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
)

type donation struct {
	Id     uuid.UUID `json:"id"`
	From   uuid.UUID `json:"from"`
	To     uuid.UUID `json:"to"`
	Action uuid.UUID `json:"action"`
}

type account struct {
	Id     uuid.UUID `json:"id"`
	Action uuid.UUID `json:"action"`
	Name   string    `json:"name"`
	Stuff  string    `json:"stuff"`
}

type action struct {
	id     uuid.UUID `json:"id"`
	name   string    `json:"name"`
	thingy string    `json:"thingy"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/add/account/", AddAccount)
	router.HandleFunc("/add/action/", AddAction)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "use this to trigger an acction on an account")
}

//func AddAccount
