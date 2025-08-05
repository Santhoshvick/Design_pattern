package main

func main(){
	tv:=&Tv{}

	onCommand1:=&onCommand{
		device: tv,
	}

	offCommand1:=&offCommand{
		device: tv,
	}

	onButton:=&Button{
		command: onCommand1,
	}

	onButton.press()

	offButton1:=&Button{
		command: offCommand1,
	}

	offButton1.press()
}