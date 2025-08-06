package main

type Temper struct{
	temper1 OrderedProduct
}


func(t *Temper)getPrice()float64{
	price:=t.temper1.getPrice()
	return price+200.00
}