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
	ID     uuid.UUID `json:"id"`
	From   uuid.UUID `json:"from"`
	To     uuid.UUID `json:"to"`
	Action uuid.UUID `json:"action"`
}

type account struct {
	ID     uuid.UUID `json:"id"`
	Action uuid.UUID `json:"action"`
	Name   string    `json:"name"`
	Stuff  string    `json:"stuff"`
}

type action struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Thingy string    `json:"thingy"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	router.HandleFunc("/add/account/", AddAccount)
	router.HandleFunc("/add/action/", AddAction)
	log.Fatal(http.ListenAndServe(":8080", router))
}

//Index the primary view
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "use this to trigger an acction on an account")
}

//AddAccount view to add an account
//TODO how are we going to auth this?
//TODO will prly want a UI Form for this
func AddAccount(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "you'll add an account here")
}

//AddAction view to add an action to an account
//TODO will prly want a UI Form for this
func AddAction(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "youll add an action to your account so you can be awesome")
}
