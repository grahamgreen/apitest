package main

//Copyright 2015 Graham Green

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/garyburd/redigo/redis"
	"github.com/grahamgreen/goutils"
	"github.com/nu7hatch/gouuid"
)

//Index the primary view
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "use this to trigger an acction on an account")
}

//TODO will need so checking so that the accountIDs in the requests exist
//TODO all these add object workflows should be refactored to avoid duplication

//AddAccount view to add an account
//TODO how are we going to auth this?
//TODO will prly want a UI Form for this
func AddAccount(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "you'll add an account here")
	var account Account
	//This should look familiar
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &account); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else {
		//will eventually pass this write down a channel
		//js, _ := json.Marshal(account)
		c, err := redis.Dial("tcp", "192.168.59.103:6379")
		if err != nil {
			panic(err)
		}
		defer c.Close()
		u4, err := uuid.NewV4()
		goutils.Check(err)
		account.ID = *u4
		c.Do("SET", account.ID.String(), account.ToJSONString())
		c.Do("SADD", "accounts", account.ID.String())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(account.ToJSONString()))
	}

}

//AddAction view to add an action to an account
//TODO will prly want a UI Form for this
func AddAction(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "youll add an action to your account so you can be awesome")
	var action Action
	//This should look familiar
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &action); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	} else {
		c, err := redis.Dial("tcp", "192.168.59.103:6379")
		if err != nil {
			panic(err)
		}
		defer c.Close()

		//check that the account exists
		exists, _ := redis.Bool(c.Do("SISMEMBER", "accounts", action.Account))
		if !exists {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(404) // account not found TODO find better return code
			w.Write([]byte("account not found"))
		} else {
			u4, err := uuid.NewV4()
			goutils.Check(err)
			action.ID = *u4
			c.Do("SET", action.ID.String(), action.ToJSONString())
			c.Do("SADD", action.Account+":actions", action.ID.String())
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte(action.ToJSONString()))
		}
	}
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
		c, err := redis.Dial("tcp", "192.168.59.103:6379")
		if err != nil {
			panic(err)
		}
		defer c.Close()

		//account checks
		toExists, _ := redis.Bool(c.Do("SISMEMBER", "accounts", donation.To))
		fromExists, _ := redis.Bool(c.Do("SISMEMBER", "accounts", donation.From))
		if !(toExists && fromExists) {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(404) // account not found TODO find better return code
			w.Write([]byte("account not found"))
			return
		}
		validAction, _ := redis.Bool(c.Do("SISMEMBER", donation.To+":actions", donation.Action))
		if !validAction {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(400) // account not found TODO find better return code
			w.Write([]byte("action not valid for that account"))
			return
		}
		u4, err := uuid.NewV4()
		goutils.Check(err)
		donation.ID = *u4

		c.Do("SET", donation.ID.String(), donation.ToJSONString())
		c.Do("SADD", donation.From+":donations_given", donation.ID.String())
		c.Do("SADD", donation.To+":donations_rec", donation.ID.String())
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(donation.ToJSONString()))
	}
}
