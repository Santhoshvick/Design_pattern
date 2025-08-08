package main

import (
	"fmt"
	"sync"
)

type Employee struct{
	EmpName string
	EmpId int64
	Location string
	Salary float64
	Role string
}


var Global *Employee
var once sync.Once


func CreateInstanceOnce()*Employee{
	if(Global==nil){
		once.Do(func(){
			fmt.Println("Employee Table is creaed Successfully")
			Global=&Employee{EmpName: "John",EmpId: 1901,Location: "Chennai",Salary: 450000.00,Role: "Software Developer"}
		})
	}

	return  Global
}

func (e *Employee)Update(ename string,eid int64,loc string,sal float64,role string){
	e.EmpName=ename
	e.EmpId=eid
	e.Location=loc
	e.Salary=sal
	e.Role=role

	fmt.Print("Eid:",e.EmpId)
	fmt.Print("Ename:",e.EmpName)
	fmt.Print("Location:",e.Location)
	fmt.Printf("Salary:%1f",e.Salary)
	fmt.Print("Role:",e.Role)

}
