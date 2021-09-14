package main

import (
    "fmt"
    "time"
) 

func main() {
	currentTime := time.Now()

    fmt.Println("Nous sommes le", currentTime.Format("02 January"), "il est", currentTime.Format("15h04"),".")
}
