package main

import (
	"fmt"
)

type Status int
const(
	pending		Status = iota
	in_progress
	completed
)


type Activity struct {
	id          string
	name        string
	duration    int
	description string
	status		Status
	subActivity subActivity
}

func newActivity(nombre string, duration int, description string,
	subActivity subActivity) Activity {
	// ¿Asignar id preguntandole a mongo?
	return Activity{"0", nombre, duration,
		description, pending, subActivity}

}

func (a Activity) String() string {
	return fmt.Sprintf(
		"ID: %b, \n"+
			"Name: %s, \n"+
			"Duration: %.2f min, \n"+
			"Description: %s, \n"+
			"Status: %s, \n"+
			"SubAct: {%s}, \n",
		a.id, a.name, a.duration, a.description, a.status, a.subActivity.name,
	)
}
