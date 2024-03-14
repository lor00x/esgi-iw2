package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	s := GetStringSlice()
	FillRandomStrings(&s)
	s = RemoveEchar(s)
	DisplayList(s)
	// TODO: déclarer une slice de strings

	// Créer une fonction pour remplir votre slice de
	// chaines aléatoire

	// Supprimer du tableau tout ce qui contient un caractère "E"

	// Affichez le résultat
}

func GetStringSlice() []string {
	return []string{}
}

func FillRandomStrings(s *[]string) {
	*s = append(*s, "euirhg", "oerijoigr", "odiujhiu", "iduh")
}

func RemoveEchar(s []string) []string {
	del := func(s string) bool {
		return strings.Contains(s, "o") || strings.Contains(s, "O")
	}
	slices.DeleteFunc(s, del)
	return s
}

func DisplayList(s []string) {
	fmt.Println(s)
}
