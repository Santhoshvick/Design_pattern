package main



func main(){

	gmessage:=GenerateNotification("Email")

	gmessage.sendNotification("s@gmail.com","Thanks for reaching to us")
}