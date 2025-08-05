package main

import "fmt"


type State interface {
    Press(s *Switch)
}


type Switch struct {
    state State
}

func (s *Switch) SetState(state State) {
    s.state = state
}

func (s *Switch) Press() {
    s.state.Press(s)
}



type OnState struct{}


func (o *OnState) Press(s *Switch) {
    fmt.Println("Switching OFF")
    s.SetState(&OffState{})
}

type OffState struct{}

func (o *OffState) Press(s *Switch) {
    fmt.Println("Switching ON")
    s.SetState(&OnState{})
}

func main() {
    sw := &Switch{}
	fmt.Println(sw)
    sw.SetState(&OffState{}) 
    
    sw.Press() 
    sw.Press() 
    sw.Press() 
}

