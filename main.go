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

	fmt.Println("Press any key on the keyboard.Press ESC to quit.")

	for {
		char, key, err := keyboard.GetSingleKey()
		fmt.Println(char,key, err)

		//check for an error
		if err != nil {
			log.Fatal(err)
		}

		if key != 0 {
			//for a: char is 97 and key is 0
			//for ESC: char is 0 and key is 27
			
			fmt.Println("You pressed", char, key)
		} else {
			fmt.Println("You pressed", char)
		}

		if key == keyboard.KeyEsc {
			break
		}
	}

	fmt.Println("Program exiting ...")

}
