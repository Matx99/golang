package main

import "fmt"

func main() {

	fmt.Println("Entrer un élément : ")

	var first string

	fmt.Scanln(&first)
	fmt.Println("Entrer un élément : ")
	var second string
	fmt.Scanln(&second)

	fmt.Print("Dans ma valise, il y a... ")

	fmt.Print(first + " et " + second)
}
