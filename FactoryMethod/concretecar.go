package main

type Car struct{
	name string
	price int64
}

func (c *Car)setName(name string){
	c.name=name
}
func (c *Car)setPrice(price int64){
	c.price=price
}

func (c *Car) getName()string{
	return c.name
}

func(c *Car)getPrice()int64{
	return c.price
}
