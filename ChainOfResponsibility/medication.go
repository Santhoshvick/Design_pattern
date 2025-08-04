package main

import "fmt"

type Medication struct{
	next Department
}


func (m *Medication)execute(p *Patient){
	if p.MedicationDone{
		fmt.Println("The medicines has successfully given to the Patinet")
		m.next.execute(p)
		return
	}
	fmt.Println("The patient is waiting to receive the medicine")
	p.MedicationDone=true
	m.next.execute(p)
}

func (m *Medication)setNext(next Department){
	m.next=next
}