package main



func getflight(name string)(Iflight,error){
	if(name=="Emirates"){
		return  Emirates1(),nil
	}
	return nil,nil
}