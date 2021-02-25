package main


type Activity struct {
	id          uint64
	name        string
	duration    float32
	description string
	state       bool // true: not finish
	subActivity subActivity
}

func newActivity(nombre string, duration float32, description string,
										subActivity subActivity) Activity {

	return Activity{nil, nombre, duration,
		description, true, subActivity}

}

