package main

import "fmt"

type tv struct{}

func (t *tv)tvOn(){
	fmt.Println("TV is Turned on")
}
func(t *tv)tvOff(){
	fmt.Println("Tv i s Turned OFF")
}

type light struct{}

func (l *light) lightsOn(){
	fmt.Println("Lights turned on")
}


func(l *light) lightsOff(){
	fmt.Println("Lights turned off")
}


type musicSytem struct{
}

func(m *musicSytem)musicOn(){
	fmt.Println("Music System is turned on")
}

func(m *musicSytem)musicOff(){
	fmt.Println("Music System is turned off")
}




