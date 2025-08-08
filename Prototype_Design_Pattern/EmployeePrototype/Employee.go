package main

import "fmt"


type clonable interface{
	copy()clonable
}


type Employee struct{
	Eid int64
	EmployeeName string
	EmployeeDesignation string
	EmployeeSalary float64
	EmployeeManager string
}

func(e *Employee)copy()clonable{
	return &Employee{
		EmployeeName: e.EmployeeName,
		Eid: e.Eid,
		EmployeeDesignation: e.EmployeeDesignation,
		EmployeeSalary: e.EmployeeSalary,
		EmployeeManager: e.EmployeeManager,
	}
	
}

func main(){
	employee1:=&Employee{
		Eid: 10212,
		EmployeeName: "John Maverick",
		EmployeeDesignation: "Software Engineer",
		EmployeeSalary:450000.00 ,
		EmployeeManager: "Mavin",
	}
	fmt.Println(employee1)

	employee2:=employee1.copy().(*Employee)
	employee2.EmployeeName="Dewald Brewis"
	employee2.EmployeeSalary=550000.00
	employee2.Eid=100001 
	fmt.Println(employee2)
}