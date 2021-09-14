package main

import (
    "fmt"
    "math/rand"
    "time"
)
        
func main() {
    rand.Seed(time.Now().UnixNano())
    six()
	twenty()
	oneHundred()
}

func six() {
	min := 0
    max := 6
    fmt.Println("|",rand.Intn(max - min + 1) + min,"|")
}

func twenty() {
	min := 0
    max := 20
    fmt.Println("|",rand.Intn(max - min + 1) + min,"|")
}

func oneHundred() {
	min := 0
    max := 100
    fmt.Println("|",rand.Intn(max - min + 1) + min,"|")
}