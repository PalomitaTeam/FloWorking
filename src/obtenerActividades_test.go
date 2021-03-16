package main

import (
	"testing"
)

func TestObntenerActividades(t *testing.T) {

	_, colec, _ := connectToMongo()
	err := getAllActivities(colec)
	if err != nil {
		t.Fail()
	}
}