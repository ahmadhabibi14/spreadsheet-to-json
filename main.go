package main

import (
	"fmt"
	"log"
	"spreadsheet-json/tojson"
)

func main() {
	err := tojson.ConvertToJSON(
		"1TmAjrclFHUwDA1487ifQjX4FzYt9y7eJ0gwyxtwZMJU",
		522117981, "tsv", "data.json",
	)
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println("Success")
}
