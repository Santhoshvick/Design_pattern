package main

import "fmt"

type Logging interface{
	validate(string,string)
}


type Login struct{
	Email string
	Password string
}
func(l *Login)validate(email string,password string){
	fmt.Println("Please Enter the Password to login")
	fmt.Scanf("%s",email)
	fmt.Println("Please Enter the Password")
	fmt.Scanf("%s",password)
	if(l.Email!=email||l.Password!=password){
		if(email!=l.Email){
			fmt.Println("Incorrect Email Id")
		}else if(password!=l.Password){
			fmt.Println("Incorrect Password")
		}else{
			fmt.Println("Incorrect Email Id and Password")
		}
	}else{
		fmt.Println("Log in Successfully")
	}
}



func main(){
	log1:=&Login{Email: "abc@gmail.com",Password: "1234"}

	log1.validate("abc@gmail.com","1234")
}

