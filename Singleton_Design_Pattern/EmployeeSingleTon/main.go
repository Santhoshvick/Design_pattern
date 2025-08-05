package main



import "fmt"

func main(){
	
	Employee1:=&Employee{EmpName: "Santhosh",EmpId: 1010,Location: "Chennai",Salary: 450000.00,Role: "Software Engineer"}
	fmt.Println(Employee1.CreateInstanceOnce())

	Employee2:=&Employee{EmpName: "NirmalRaj",EmpId: 1020,Location: "Thirumullaivoyal",Salary: 475000.90,Role: "Software Engineer"}
	fmt.Println(Employee2.CreateInstanceOnce())

}