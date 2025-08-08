package main

type CarDepartment interface{
	execute(*Customer)
	setNext(CarDepartment)
}