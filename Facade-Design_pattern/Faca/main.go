package main

func main(){
	tv1:=&tv{}
	light1:=&light{}
	ms:=&musicSytem{}

	home:=newHouseSystem(tv1,light1,ms)

	home.Turnon()
	home.Turnoff()
}