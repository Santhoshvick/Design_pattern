package main

import (
	"fmt"
	"time"
)

type State interface{
	controller(*switch1)
}


type switch1 struct{
	state State
}

func(s *switch1)setState(state State){
	s.state=state
}

func (s *switch1)controller(){
	s.state.controller(s)
}


type Red struct{}

func(r *Red)controller(s *switch1){
	time.Sleep(10*time.Second)
	s.setState(&Yellow{})
	fmt.Println("The Signal tured into yellow")
}


type Yellow struct{}

func(y *Yellow)controller(s *switch1){
	time.Sleep(10*time.Second)
	s.setState(&Green{})
	fmt.Println("The Signal Turned into green")
}


type Green struct{}

func(g *Green)controller(s *switch1){
	s.setState(&Red{})
	fmt.Println("The signal Turned into red")
}


func main(){
	sw:=&switch1{}
	sw.setState(&Green{})
	sw.controller()
	sw.controller()
	sw.controller()
}



