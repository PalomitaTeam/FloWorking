package main

import "fmt"

type Activity struct {
	id          uint64
	name        string
	duration    float32
	description string
	finish       bool
	subActivity subActivity
}

func newActivity(nombre string, duration float32, description string,
	subActivity subActivity) Activity {
	// ¿Asignar id preguntandole a mongo?
	return Activity{0, nombre, duration,
		description, false, subActivity}

}

func (a Activity) String() string {
	return fmt.Sprintf(
		"ID: %b, \n"+
			"Name: %s, \n"+
			"Duration: %.2f min, \n"+
			"Description: %s, \n"+
			"Finish?: %t, \n"+
			"SubAct: {%s}, \n",
		a.id, a.name, a.duration, a.description, a.finish, a.subActivity.name,
	)
}
