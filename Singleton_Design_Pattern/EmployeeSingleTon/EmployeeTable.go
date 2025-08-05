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


func(e *Employee) CreateInstanceOnce()*Employee{
	if(Global==nil){
		once.Do(func(){
			fmt.Println("Employee Table is creaed Successfully")
			Global=&Employee{}

		})
	}
	Global=&Employee{EmpName: e.EmpName,EmpId: e.EmpId,Location: e.Location,Salary: e.Salary,Role: e.Role}
	
	return  Global
}
