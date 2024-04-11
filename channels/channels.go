package main

import (
	"fmt"
	"time"
)

func main() {
	// Créer une fonction Serveur, qui reçoit des commandes de boisson
	// et un chan à qui envoyer la boisson
	// Créer une fonction Client, qui reçoit des boissons (par un channel)
	// Main:
	// - crée tout le monde (Serveur et Clients)
	// - passe commande au serveur pour chaque client
	// - attends que tout le monde termine

	// Patron : main
	// Serveur
	// Client * 5
	// Patron -> tout le monde => 1 channel de contrôle
	// Client -> Serveur => 1 channel de commande
	// Commande: string

	doneServeur := make(chan bool)
	doneClient := make(chan bool)
	commandes := make(chan Commande)

	go Serveur(commandes, doneServeur)

	for range 5 {
		go Client(commandes, doneClient)
	}

	time.Sleep(5 * time.Second)
	for range 5 {
		doneClient <- true
	}
	doneServeur <- true
}

type Commande struct {
	Boisson string
	Retour  chan string
}

func Serveur(commandes chan Commande, done chan bool) {
	for {
		select {
		case commande := <-commandes:
			fmt.Println("Serveur : Préparation de " + commande.Boisson)
			boisson := commande.Boisson
			commande.Retour <- boisson
		case <-done:
			fmt.Println("Serveur : terminé !")
			return
		}
	}
}

func Client(commandes chan Commande, done chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Client : au revoir !")
			return
		default:
			retour := make(chan string)
			commandes <- Commande{
				Boisson: "Mojito",
				Retour:  retour,
			}
			boisson := <-retour
			fmt.Println("Client : J'ai reçu un " + boisson)
			time.Sleep(1 * time.Second)
		}
	}
}
