package main

import "fmt"

type Tea struct{}

func(t *Tea)boilWater(){
	fmt.Println("Water is boiling for the Tea ♨️☕️")
}

func(t *Tea)brew(){
	fmt.Println("Brew is mixed with the water")
}

func(t *Tea)pourInCup(){
	fmt.Println("The Tea is filled in the glass cup")
}

func(t *Tea)AddCondiments(){
	fmt.Println("The Extra Condiments is added to the tea flavour in the cup")
}


