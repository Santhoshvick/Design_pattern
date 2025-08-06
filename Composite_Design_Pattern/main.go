package main

import "fmt"

func main(){
	dev1:=&Developer{Name: "xyz",Position: "Frontend Developer"}
	dev1.display()
	fmt.Println()

	dev2:=&Developer{Name: "Aamir",Position: "backend Developer"}
	dev2.display()
	fmt.Println()

	manager1:=&Manager{Name: "www",Position: "Product Manager"}
	manager1.display()

	manager1.Add(dev1)
	manager1.Add(dev2)
	
}