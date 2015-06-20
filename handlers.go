package main

//Copyright 2015 Graham Green

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

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

//Donate commit action on an account
func Donate(w http.ResponseWriter, r *http.Request) {
	var donation Donation
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &donation); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else {
		js, _ := json.Marshal(donation)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write(js)
	}
}
