package main
import "fmt"
type Coffee struct{}

func(c *Coffee)boilWater(){
	fmt.Println("Water is boiling for the Coffee ♨️☕️")
}

func(c *Coffee)brew(){
	fmt.Println("Brew is mixed with the water")
}



func(c *Coffee)pourInCup(){
	fmt.Println("The Coffee is filled in the glass cup")
}

func(c *Coffee)AddCondiments(){
	fmt.Println("The Extra Condiments is added to the tea flavour in the cup")
}
