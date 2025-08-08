package main

import "fmt"


type Employee struct{}

func(e *Employee)WakeUp(){
	fmt.Println("************Employeers DailyRouting*******************")
	fmt.Println("Employee usually wake up's at 6.00AM 🛏️")
}

func(e *Employee)GetReady(){
	fmt.Println("The Employee gets ready with his casual or formal outfit")
}

func(e *Employee)GoToWorkOrSchool(){
	fmt.Println("Travel by office cab or Personal vehicle")
}

func(e *Employee)DoMainTask(){
	fmt.Println("The Employee check's his all task and attends meeting in the morning session")
}

func(e *Employee)ComeHome(){
	fmt.Println("The employee drive back with his car back to the home")
}

func(e *Employee)Relax(){
	fmt.Println("Employee Relax in remaining of his time")
}