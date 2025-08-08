package main

import "fmt"

type TestDrive struct{
	next CarDepartment
}

func (t *TestDrive)execute(c *Customer){
	// c.TestDrive=true
	if c.TestDrive{
		fmt.Println("Test Drive is succcessfully Completed ")
		t.next.execute(c)
		return
	}
	fmt.Println("The Customer is waiting for the Test Drive")
	c.TestDrive=true
	t.next.execute(c)
} 
func (t *TestDrive)setNext(next CarDepartment){
	t.next=next
}
