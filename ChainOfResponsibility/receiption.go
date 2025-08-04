package main

import "fmt"

type Reception struct{
	next Department
}


func(r * Reception)execute(p *Patient){
	if p.registrationDone{
		fmt.Println("Patient has been successfully  registerred")
		r.next.execute(p)
		return
	}
	fmt.Println("Patient is On the admission process")
	p.registrationDone=true
	r.next.execute(p)
}

func (r *Reception) setNext(next Department){
	r.next=next
}