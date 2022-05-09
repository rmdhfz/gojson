package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONArray(t *testing.T) {
	bytes, _ := json.Marshal(Customer{
		FirstName:  "Hafiz Ramadhan",
		MiddleName: "Ganteng",
		LastName:   "Banget",
		Age:        30,
		Married:    true,
		Hobbies: []string{
			"Gaming",
			"Belajar",
			"Coding",
		},
	})
	fmt.Println(string(bytes))
}

func TestJSONArrayDecode(t *testing.T) {
	jsonString := `{"FirstName":"Hafiz Ramadhan","MiddleName":"Ganteng","LastName":"Banget","Age":30,"Married":true,"Hobbies":["Gaming","Belajar","Coding"]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
}

func TestJSONArrayComplex(t *testing.T) {
	bytes, _ := json.Marshal(Customer{
		FirstName:  "Hafiz Ramadhan",
		MiddleName: "Ganteng",
		LastName:   "Banget",
		Age:        30,
		Married:    true,
		Hobbies: []string{
			"Gaming",
			"Belajar",
			"Coding",
		},
		Addresses: []Address{
			{
				Street:     "Bogor",
				Country:    "Indonesia",
				PostalCode: "10000",
			},
			{
				Street:     "Depok",
				Country:    "Indonesia",
				PostalCode: "2000",
			},
		},
	})
	fmt.Println(string(bytes))
}

func TestJSONArrayComplexDecode(t *testing.T) {
	jsonString := `{"FirstName":"Hafiz Ramadhan","MiddleName":"Ganteng","LastName":"Banget","Age":30,"Married":true,"Hobbies":["Gaming","Belajar","Coding"],"Addresses":[{"Street":"Bogor","Country":"Indonesia","PostalCode":"10000"},{"Street":"Depok","Country":"Indonesia","PostalCode":"2000"}]}`
	jsonBytes := []byte(jsonString)

	customer := &Customer{}
	err := json.Unmarshal(jsonBytes, customer)
	if err != nil {
		panic(err)
	}
	fmt.Println(customer)
	fmt.Println(customer.FirstName)
	fmt.Println(customer.MiddleName)
	fmt.Println(customer.LastName)
	fmt.Println(customer.Hobbies)
	fmt.Println(customer.Addresses)
}

func TestOnlyJSONArrayComplexDecode(t *testing.T) {
	jsonString := `[{"Street":"Bogor","Country":"Indonesia","PostalCode":"10000"},{"Street":"Depok","Country":"Indonesia","PostalCode":"2000"}]`
	jsonBytes := []byte(jsonString)

	address := &[]Address{}
	err := json.Unmarshal(jsonBytes, address)
	if err != nil {
		panic(err)
	}
	fmt.Println(address)
}
