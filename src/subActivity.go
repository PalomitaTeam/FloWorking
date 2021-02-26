package main

type subActivity struct {
	name  string
	status Status
}

func newSubActivity(name string) subActivity {
	return subActivity{name, pending}
}