package main

import (
	"fmt"
	"./light"
)

func main() {

	light := light.IsLightOn()

	if (light) {
		fmt.Println("light on!")
	} else {
		fmt.Println("light off")
	}	
	
}
