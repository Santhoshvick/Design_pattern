package main

type Iflight interface{
	setFlightName(name string)
	setFlightTicketPrice(price int64)
	setClass(class string)
	getFlightName()string
	getFlightTicketPrice()int64
	getFlightClass()string
}