package main

import "fmt"

type Message struct {
	Text string
}

func (m *Message) Clone() *Message {
	return &Message{Text: m.Text}
}

func main() {
	msg1 := &Message{Text: "Hello"}
	msg2 := msg1.Clone()

	fmt.Println("Original:", msg1.Text)
	fmt.Println("Clone   :", msg2.Text)
}
