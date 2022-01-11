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
	
	char:= ' '
	for char != 'q' && char !='Q'{
		char, _, err = keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err)
		}
		


		i,_:=strconv.Atoi(string(char))

		_,ok:=coffees[i]
		if ok{
			fmt.Println(fmt.Sprintf("You chose %s",coffees[i]))
		}

	}

	fmt.Println("Program exiting ...")

}
