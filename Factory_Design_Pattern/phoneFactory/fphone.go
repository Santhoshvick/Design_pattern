package main

type FPhone interface{
	setName(name string)
	setPrice(price int)
	setModel(model string)
	setStorage(storage string)
	getName()string
	getPrice()int
	getModel()string
	getStorage()string
}
