package main

func main(){
	student:=&Student{}
	employee:=&Employee{}

	t1:=&process{routine: student}
	t2:=&process{routine: employee}
	t1.DailyRoutine()
	t2.DailyRoutine()
}