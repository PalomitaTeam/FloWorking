package main

type subActivity struct {
	id     *string
	name   string
	status Status
}

func newSubActivity(name string) subActivity {
	return subActivity{nil,name, pending}
}