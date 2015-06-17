package main

import (
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nu7hatch/gouuid"
)

type donation struct {
	Id     uuid.UUID
	From   uuid.UUID
	To     uuid.UUID
	Action uuid.UUID
}

type account struct {
	Id     uuid.UUID
	Action uuid.UUID
	Name   string
	Stuff  string
}

type action struct {
	id     uuid.UUID
	name   string
	thingy string
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
