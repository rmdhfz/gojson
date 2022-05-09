package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Address struct {
	Street     string
	Country    string
	PostalCode string
}

type Customer struct {
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Married    bool
	Hobbies    []string
	Addresses  []Address
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
