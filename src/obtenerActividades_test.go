package main

import (
	"github.com/stretchr/testify/assert"
	"src/Activity"
	"testing"
	"fmt"
)

func TestCreaAct(t *testing.T) {
	actividad := newActivity(
			"Actividad 1",
			60,
			"descripción actividad de prueba",
			newSubActivity("SUbActivity1"),
		)
	client, colec, err := connectToMongo()


}