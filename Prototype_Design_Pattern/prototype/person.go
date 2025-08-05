package main

import "fmt"

type cloneable interface{
	clone()cloneable
}

type person struct{
	name string
	id int64
}

func (p *person) clone()cloneable{
	return &person{
		name: p.name,
		id: p.id,
	}
}

func main(){
	person1:=&person{name:"Santhosh",id: 101}
	fmt.Println(person1)

	clonePerson:=person1.clone().(*person)
	fmt.Println(clonePerson)

	
}
