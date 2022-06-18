package main

import (
	"encoding/json"
	"log"
	"os"
)

func main() {
	var problems Problems
	err := json.NewDecoder(os.Stdin).Decode(&problems)
	if err != nil {
		log.Fatal(err)
	}

	// pprint(problems)

	var results Solutions

	// TODO: solve!

	err = json.NewEncoder(os.Stdout).Encode(results)
	if err != nil {
		log.Fatal(err)
	}
}
