package main

import (
	"fmt"
	"maps"
)

var schools = map[string]string{
	"PARIS":            "ESGI",
	"SOPHIA-ANTIPOLIS": "ESSI",
	"NANTES":           "ESGI",
}

func main() {

	if _, found := schools["PARIS"]; found {
		fmt.Println("TROUVE")
	}

	// fmt.Printf("%#v, %#v", town, found)

	maps.DeleteFunc(schools, FILTER)

	for town := range schools {
		fmt.Printf("Town: %s\n", town)
	}

	for town, school := range schools {
		fmt.Printf("Town: %s, School: %s\n", town, school)
	}
}

func FILTER(town string, school string) bool {
	return school != "ESGI"
}
