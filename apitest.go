package main

import "github.com/nu7hatch/gouuid"

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

type test struct {
	foo string
}
