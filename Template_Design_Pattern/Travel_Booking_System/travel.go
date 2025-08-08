package main



type Travel interface{
	ChooseDestination(string)
	TravelMode()
	Book()
	Confirm()
}

type Booking struct{
	travel Travel
}

func(b *Booking)schedule(){
	b.travel.ChooseDestination("Chennai")
	b.travel.TravelMode()
	b.travel.Book()
	b.travel.Confirm()
}