package main

import "fmt"

type Student struct{}

func(s *Student)WakeUp(){
	fmt.Println("************Students DailyRouting*******************")
	fmt.Println("Students  usually wake up's at 7.00AM 🛏️")
}

func(s *Student)GetReady(){
	fmt.Println("The Students gets ready with his school uniform outfit")
}

func(s *Student)GoToWorkOrSchool(){
	fmt.Println("Travel by School bus")
}

func(s *Student)DoMainTask(){
	fmt.Println("The Students will attend all the classes based on their scehuled time table ")
}

func(s *Student)ComeHome(){
	fmt.Println("The school students returns back to his home through school bus")
}

func(s *Student)Relax(){
	fmt.Println("Students Relax in remaining of his time with the family members")
}