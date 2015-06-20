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

//ToJSONString return the json string of the struct w/out the id
func (d Donation) ToJSONString() string {
	//hack alert
	return `{"from": "` + d.From + `", "to": "` + d.To + `"}`
}

//Account an account; can send donations by default
//must have an accociated action to receive actions
type Account struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Stuff string `json:"stuff"`
}

//BADBADBAD NOT DRY

//ToJSONString return the json string of the struct w/out the id
func (a Account) ToJSONString() string {
	return `{"name": "` + a.Name + `", "stuff": "` + a.Stuff + `" }`
}

//Action a action associated w/ an account
type Action struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Thingy string `json:"thingy"`
}

//ToJSONString return the json string of the struct w/out the id
func (a Action) ToJSONString() string {
	return `{"name": "` + a.Name + `", "thingy": "` + a.Thingy + `" }`
}
