package main

import "fmt"

type Doctor struct{
	next Department
}

func (d *Doctor)execute(p *Patient){
	if p.medicalDone{
		fmt.Println("The Patient is successfully consulted the doctor")
		d.next.execute(p)
		return
	}
	fmt.Println("The patient is waiting to consult the doctor")
	p.medicalDone=true
	d.next.execute(p)
}

func (d *Doctor)setNext(next Department){
	d.next=next
}