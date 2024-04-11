package main

import (
	"fmt"
)

func main() {
	// Ecrivez une interface Role avec quelques méthodes
	// Ecrivez les structures concrètes qui implémentent l'interface
	// Initialisez un tableau/slice de Role
	// Appelez la même méthode de modification sur chacun des roles du tableau/slice
	// Appelez la même méthode d'affichage sur chacun des roles du tableau/slice
	// Ecrivez un switch/case qui affiche un message différent suivant
	// le type de chaque Rôle

	roles := make([]Role, 0)
	root := Root{}
	user := User{
		name:   "Thomas",
		active: true,
	}
	roles = append(roles, &root, &user)

	for _, role := range roles {
		role.SetActive(true)
		fmt.Println(role.GetName())
		SwitchCase(role)
	}

}

type Role interface {
	GetName() string
	IsAdmin() bool
	SetActive(bool)
}

func SwitchCase(role Role) {
	switch role.(type) {
	case *Root:
		fmt.Println("Root! ")
	case *User:
		fmt.Println("User !")
	}
}
