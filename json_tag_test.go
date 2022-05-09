package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Product struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

func TestJSONTag(t *testing.T) {
	product := Product{
		Id:       "1",
		Name:     "Apple Mac Book Pro",
		ImageURL: "https://example.com/image.png",
	}
	encoded, err := json.Marshal(product)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encoded))
}

func TestJSONTagDecode(t *testing.T) {
	jsonString := `{"id":"1","name":"Apple Mac Book Pro","image_url":"https://example.com/image.png"}`
	jsonByte := []byte(jsonString)
	product := &Product{}

	err := json.Unmarshal(jsonByte, product)
	if err != nil {
		panic(err)
	}

	fmt.Println(product)
	fmt.Println(product.Id)
	fmt.Println(product.Name)
	fmt.Println(product.ImageURL)
}
