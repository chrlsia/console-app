package main

import (
	"fmt"
	"log"
	"strconv"

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
	
	coffees:=make(map[int]string)
	coffees[1]="Cappucino"
	coffees[2]="Latte"
	coffees[3]="Americano"
	coffees[4]="Mocha"
	coffees[5]="Macchiato"
	coffees[6]="Espresso"
	
	fmt.Println("MENU")
	fmt.Println("---")
	fmt.Println("1-",coffees[1])
	fmt.Println("2-",coffees[2])
	fmt.Println("3-",coffees[3])
	fmt.Println("4-",coffees[4])
	fmt.Println("5-",coffees[5])
	fmt.Println("6-",coffees[6])
	fmt.Println("Q- Quit the program ...")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		
		if char == 'q' || char =='Q' {
			break
		}

		i,_:=strconv.Atoi(string(char))
		if i<=6 && i>=1 {
			fmt.Println(fmt.Sprintf("You chose %s",coffees[i]))
		} else {
			fmt.Println("Number out of bounds.Type a numbers from 1 to 6")
		}

	}

	fmt.Println("Program exiting ...")

}
