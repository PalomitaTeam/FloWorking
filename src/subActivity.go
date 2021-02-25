package main

type subActivity struct {
	name  string
	state bool
}

func newSubActivity(name string) subActivity {
	return subActivity{name, true}
}