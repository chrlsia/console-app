package main

import (
	"fmt"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = keyboard.Close()
	}()

	fmt.Println("MENU")
	fmt.Println("---")
	fmt.Println("1- Cappucino")
	fmt.Println("2- Latte")
	fmt.Println("3- Americano")
	fmt.Println("4- Mocha")
	fmt.Println("5- Macchiato")
	fmt.Println("6- Espresso")
	fmt.Println("Q- Quit the program ...")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("You chose",char)

		if char == 'q' || char =='Q' {
			break
		}
	}

	fmt.Println("Program exiting ...")

}
