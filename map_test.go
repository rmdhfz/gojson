package gojson

import (
	"encoding/json"
	"fmt"
	"testing"
)

// untuk me-decode json ke map
func TestMapDecode(t *testing.T) {
	jsonString := `{"id": 1, "name":"Apple Mac Book Pro", "price": "123456"}`
	jsonBytes := []byte(jsonString)

	var result map[string]interface{}
	json.Unmarshal(jsonBytes, &result)

	fmt.Println(result)
	fmt.Println(result["id"])
	fmt.Println(result["name"])
	fmt.Println(result["price"])
}

// untuk mengkonversi map ke json
func TestMapEncode(t *testing.T) {
	product := map[string]interface{}{
		"id":    1,
		"name":  "Apple Mac Book Pro",
		"price": "123456",
	}
	bytes, _ := json.Marshal(product)
	fmt.Println(string(bytes))
}
