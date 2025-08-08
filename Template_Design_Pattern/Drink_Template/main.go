package main
// import "fmt"

func main(){
	tea:=&Tea{}
	coffee:=&Coffee{}

	t1:=&Beverage{drink: tea}

	t1.prepare()

	t2:=&Beverage{drink: coffee}
	t2.prepare()
}