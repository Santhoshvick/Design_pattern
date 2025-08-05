package main

import "fmt"

type Tv struct{
	isBoolean bool
}

func(t *Tv)On(){
	t.isBoolean=true
	fmt.Println("TV is Turned on")
}

func(t *Tv)off(){
	t.isBoolean=false
	fmt.Println("Tv is Turned Off")
}