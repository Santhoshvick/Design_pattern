package main

import "fmt"

func main(){
	kia1,err:=buildcar("Kia")
	if err!=nil{
		return 
	}
	fmt.Println(kia1)

	hyundai1,err:=buildcar("Hyundai")
	if err!=nil{
		return
	}

	printCar(kia1)
	printCar(hyundai1)
}


func printCar(c IcarBuilder){
	fmt.Println("---------------------------")
	fmt.Printf("Car Engine Name: %s",c.getCar().engineName)
	fmt.Println()

	fmt.Printf("Car Color:%s",c.getCar().color1)
	fmt.Println()

	fmt.Printf("Car Tyre Name:%s",c.getCar().tyre1)
	fmt.Println()
	fmt.Printf("Car Model:%s",c.getCar().model)
	fmt.Println()
	fmt.Println("------------------------------")

}


