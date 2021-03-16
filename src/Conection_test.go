package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"tests/"
)

func TestConexion(t *testing.T) {
	cliente, collection := connectToMongo()

}