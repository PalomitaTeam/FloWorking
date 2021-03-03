package main

import "fmt"

type Status int

const(
	pending		Status = iota
	in_progress
	completed
)


type Activity struct {
	id          *string `bson:"_id"`
	name        string  `bson:"name,omitempty"`
	duration    int     `bson:"duration"`
	description string  `bson:"description,omitempty"`
	status		Status  `bson:"status"`
}

func newActivity(nombre string, duration int, description string) Activity {
	// ¿Asignar id preguntandole a mongo?
	return Activity{nil, nombre, duration,
		description, pending}

}

func (a Activity) String() string {
	return fmt.Sprintf(
		"ID: %s, \n"+
			"Name: %s, \n"+
			"Duration: %d min, \n"+
			"Description: %s, \n"+
			"Status: %d, \n"+
			"SubAct: {nil}, \n",
		a.id, a.name, a.duration, a.description, a.status,
	)
}

