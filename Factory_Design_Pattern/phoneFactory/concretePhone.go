package main

type Phone struct{
	Name string
	Price int
	Storage string
	Model string
}

func(p *Phone)setName(name string){
	p.Name=name
}
func(p *Phone)setPrice(price int){
	p.Price=price
}
func(p *Phone)setStorage(storage string){
	p.Storage=storage
}
func(p *Phone)setModel(model string){
	p.Model=model
}


func(p *Phone)getName()string{
	return p.Name
}

func(p *Phone)getPrice()int{
	return p.Price
}

func (p *Phone)getStorage()string{
	return p.Storage
}

func (p *Phone)getModel()string{
	return p.Model
}