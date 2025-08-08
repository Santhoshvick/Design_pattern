package main

import "fmt"

type OnlineLearning interface{
	register(*Course)
}

type Course struct{
	onlinelearning OnlineLearning
}

func(c *Course)setState(ol OnlineLearning){
	c.onlinelearning=ol
}

func(c *Course)register(){
    c.onlinelearning.register(c)
}

type NotStarted struct{}

func(n *NotStarted)register(c *Course){
	fmt.Println("The Course you have register is not Started yet")
	c.setState(&InProgress{})
}

type InProgress struct{}

func(i *InProgress)register(c *Course){
	fmt.Println("The Course status is currently in progress")
	c.setState(&Completed{})
}

type Completed struct{}

func(c1 *Completed)register(c *Course){
	fmt.Println("The Course You have registered is finished successfully please upload your project details")
	c.setState(nil)
}

func main(){
	course:=&Course{}

	course.setState(&NotStarted{})
	course.register()
	course.register()
	course.register()
}