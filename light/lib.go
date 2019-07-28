package light

import (
	"fmt"
	"os"
	"github.com/stianeikeland/go-rpio"
)



func IsLightOn() bool {
	err := rpio.Open()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pin := rpio.Pin(18) // Physical PIN 12, BCM pin 18
	pin.Input()

	input := pin.Read()

	fmt.Println(input)

	rpio.Close()

	lightOn := false
	if input == 1 {
		lightOn = true
	}

	return lightOn
}
