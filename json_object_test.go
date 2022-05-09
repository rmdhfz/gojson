package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
}

func TestJSONObject(t *testing.T) {
	bytes, _ := json.Marshal(Customer{
		FirstName:  "Hafiz Ramadhan",
		MiddleName: "Ganteng",
		LastName:   "Banget",
		Age:        30,
		Married:    true,
	})
	fmt.Println(string(bytes))
}
