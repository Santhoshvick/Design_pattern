package main

func main(){
	t1:=&Dev{}
	t2:=&Production{}


	d1:=&buildTest{deployment: t1}
	d1.startDeployment()

	d2:=&buildTest{deployment: t2}
	d2.startDeployment()
}