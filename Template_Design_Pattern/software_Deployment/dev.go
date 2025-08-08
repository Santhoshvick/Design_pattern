package main

import "fmt"

type Dev struct{}

func(d *Dev)Build(){
	fmt.Println("************DEV*****************")
	fmt.Println("The Build is attached to the dev environment")
}

func(d *Dev)Test(){
	fmt.Println("The Testing is Done in the Dev")
}

func(d *Dev)Deploy(){
	fmt.Println("The Deplyment is done successfully")
}

func(d *Dev)Notify(){
	fmt.Println("Notify the changes in the dev")
}