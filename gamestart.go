package main

import "fmt"
import "math/rand"

func main() {

	type player struct {
		firstname string
		name string
		nickname string
		score int 
		life int
	}

	var playersList []player
	var newPlayer player
	var ok = true

	for ok == true {

		fmt.Println("Entrer votre prénom : ")
		var firstname string
		fmt.Scanln(&firstname)
		newPlayer.firstname = firstname

		fmt.Println("Entrer votre nom : ")
		var name string
		fmt.Scanln(&name)
		newPlayer.name = name

		fmt.Println("Entrer votre pseudo : ")
		var nickname string
		fmt.Scanln(&nickname)
		newPlayer.nickname = nickname

		playersList = append(playersList, newPlayer)

		fmt.Printf("Il y a %v joueurs.\n", len(playersList))
		y := 1
		for i := 0; i < len(playersList); i++ {
			fmt.Printf("Le joueur %v s’appelle %v %v, il porte le pseudo %v et possède un score de %v.\n", 
														y, 
														playersList[i].firstname,
														playersList[i].name,
														playersList[i].nickname,
														playersList[i].score,
			)
			y++
		}

		fmt.Println("Voulez-vous ajouter un joueur supplémentaire ? Si vous répondez N, alors la partie commencera (O, N)\n")
		var response string
		fmt.Scanln(&response)
		if response == "O" {
			ok = true
		} else {
			if len(playersList) > 0 {
				fmt.Println("Partie lancée !\n")
				random := rand.Intn(100 - 0) + 0

				fmt.Print(random"\n")

				for i := 1; i < 5; i++ {
					fmt.Println("Tentative %v : ", i)
					fmt.Println("Joueur 1 : ")
					var tryp1 int
					fmt.Scanln(&tryp1)
					fmt.Println("Joueur 2 : ")
					var tryp1 int
					fmt.Scanln(&tryp2)
					if 
				}	
			} else {
				fmt.Println("Il faut un minimum de un joueur pour commencer la partie.")
			}
		}
	}

}
