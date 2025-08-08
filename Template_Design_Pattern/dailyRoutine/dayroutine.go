package main

type Routine interface{
	WakeUp()
	GetReady()
	GoToWorkOrSchool()
	DoMainTask()
	ComeHome()
	Relax()
}

type process struct{
	routine Routine
}



func(p *process)DailyRoutine(){
	p.routine.WakeUp()
	p.routine.GetReady()
	p.routine.GoToWorkOrSchool()
	p.routine.ComeHome()
	p.routine.Relax()
}

