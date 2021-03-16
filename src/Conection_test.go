package main

import (
	"testing"
)

func TestConexion(t *testing.T) {
	_, _, err := connectToMongo()
	if err != nil {
		t.Fail()
	}
}