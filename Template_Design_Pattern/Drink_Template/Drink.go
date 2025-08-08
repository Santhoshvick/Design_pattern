package main


type Drink interface{
	boilWater()
	brew()
	pourInCup()
	AddCondiments()
}

type Beverage struct{
	drink Drink
}

func(b *Beverage)prepare(){
	b.drink.boilWater()
	b.drink.brew()
	b.drink.pourInCup()
	b.drink.AddCondiments()
}