package main

import (
	"encoding/json"
	"fmt"
)

// main function
func main() {

	// defining a map
	var result map[string][]string

	// string json
	jsonString := `{
		"47": ["47", "47км"],
		"цветок": ["цветок", "цветочный"]
	  }`

	err := json.Unmarshal([]byte(jsonString), &result)

	if err != nil {
		// print out if error is not nil
		fmt.Println(err)
	}

	// printing details of map
	// iterate through the map
	for _, value := range result {
		fmt.Println(value[0])
	}
}
