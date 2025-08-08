package main

type Ecommerce interface{
	execute(*Customer)
	setNext(Ecommerce)
}