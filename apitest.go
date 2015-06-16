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
	uuid   uuid.UUID
	from   uuid.UUID
	to     uuid.UUID
	action uuid.UUID
}

type account struct {
	id     uuid.UUID
	action uuid.UUID
	name   string
	stuff  string
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
