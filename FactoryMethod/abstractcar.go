package main

type Icar interface{
	setName(name string)
	setPrice(price int64)
	getName()string
	getPrice()int64
}