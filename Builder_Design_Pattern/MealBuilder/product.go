package main

import "fmt"

type Product struct{
	MealName string
	Drink string
	Fries string
}


type MealBuilder interface{
     addBurger()
	 fillDrinks()
	 friesType()
}

type VegmealBuilder struct{
	Product
}


func (v *VegmealBuilder)addBurger(burger string)*VegmealBuilder{
	v.Product.MealName=burger
	return v
}

func(v *VegmealBuilder)fillDrinks(drinkName string)*VegmealBuilder{
	v.Product.Drink=drinkName
	return v
}

func(v *VegmealBuilder)friesType(fries string)*VegmealBuilder{
	v.Product.Fries=fries
	return v
}

type NonVegMealBuilder struct{
	Product
}

func (n *NonVegMealBuilder)addBurger(burger string)*NonVegMealBuilder{
	n.Product.MealName=burger
	return n
}

func(n *NonVegMealBuilder)fillDrinks(drinkName string)*NonVegMealBuilder{
	n.Product.Drink=drinkName
	return n
}

func(n *NonVegMealBuilder)friesType(fries string)*NonVegMealBuilder{
	n.Product.Fries=fries
	return n
}

func(m *VegmealBuilder)build()Product{
	return m.Product
}

func(m *NonVegMealBuilder)build1()Product{
	return m.Product
}

func main(){


	vegBuilder:=&VegmealBuilder{}

	NonvegBuilder:=&NonVegMealBuilder{}

	meal1:=vegBuilder.addBurger("Panner Burger").fillDrinks("Coke").friesType("peri peri fries").build()
	fmt.Println("Veg meal 1",meal1)


	meal2:=NonvegBuilder.addBurger("Classic zinger doubl patty crispy chciken burger").fillDrinks("Coke").friesType("Cheesy loaded fries with crishpy chicken").build1()
	fmt.Println(meal2)
}

