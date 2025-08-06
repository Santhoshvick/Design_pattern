package main

import "fmt"



type Employee interface{
	getName()string
	getPosition()string
	display()
}

type Developer struct{
	Name string
	Position string
}


func (d *Developer)getName()string{
	return d.Name
}

func (d *Developer)getPosition()string{
	return d.Position
}

func(d *Developer)display(){
	fmt.Printf("the Name of the Developer is %s and The Position is %s",d.Name,d.Position)
}

type Manager struct{
	Name string
	Position string
	emp []Employee
}

func(m *Manager)getName()string{
	return m.Name
}
func(m *Manager)getPosition()string{
	return m.Position
}

func(m *Manager)Add(e Employee){
	m.emp=append(m.emp,e)
	
	for _,e:=range m.emp{
		fmt.Println(e.getName())
		fmt.Println(e.getPosition())
	}

}

func (m *Manager)display(){
	fmt.Printf("The Name of the Manager is %s and the position is %s",m.Name,m.Position,)
}