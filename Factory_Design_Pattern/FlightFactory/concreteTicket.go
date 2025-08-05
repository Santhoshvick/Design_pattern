package main

type flight struct {
	name  string
	price int64
	class string
}

func (h *flight) setFlightName(name string) {
	h.name = name
}

func (h *flight) setFlightTicketPrice(price int64) {
	h.price = price
}

func (h *flight) setClass(class string){
	h.class=class
}

func (h *flight)getFlightName() string{
	return h.name
}



func (h *flight)getFlightTicketPrice() int64{
	return h.price
}


func (h *flight)getFlightClass() string{
	return h.class
}

