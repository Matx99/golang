package main

import "fmt"

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

		fmt.Println("Voulez-vous ajouter un joueur supplémentaire ? (O, N)\n")
		var response string
		fmt.Scanln(&response)
		if response == "O" {
			ok = true
		} else {
			ok = false
		}
	}

}
