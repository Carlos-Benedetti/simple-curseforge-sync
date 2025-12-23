package utils

import (
	"encoding/json"
	"fmt"
	"log"
)

func PrintObject(o any) {
	bytes, err := json.MarshalIndent(o, "", "  ") // Indent with two spaces
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bytes))
}
