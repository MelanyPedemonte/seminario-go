package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateResult(t *testing.T) {
	var cases = []struct {
		Input   string // input string in order to be parsed
		Success bool   // paser result
		Type    string // the input type
		Value   string // the input value
		Length  int    // value length
	}{
		{"TX02AB", true, "TX", "AB", 2},
		{"NN100987654321", true, "NN", "0987654321", 10},
		{"TX06ABCDE", false, "", "", 0},
		{"NN04000A", false, "", "", 0},
	}

	for _, testData := range cases {

		prueba := createResult(testData.Input)

		assert.Equal(t, testData.Type, prueba.Type)
		assert.Equal(t, testData.Value, prueba.Value)
		assert.Equal(t, testData.Length, prueba.Length)
		// acá agregar chequeos propios del test por ejemplo:
		// assert.Equal(t, err == nil, testData.Success)

		//todavia no funciona el test
	}
}
