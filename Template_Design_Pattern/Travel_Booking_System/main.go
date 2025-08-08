package main

func main(){
	f1:=&Flight{FlightName: "Air Arabia",FlightNumber:8920312,FromLocation: "Dubai"}
	train:=&Train{TrainName: "Vandhe Bhaarath ",TrainNumber: 67123123143,FromLocation: "Kochi"}

	t1:=&Booking{travel: f1}

	t2:=&Booking{travel: train}
	t1.schedule()
	t2.schedule()
}