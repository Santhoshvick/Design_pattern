package main

func getPhone(name string)(FPhone){
	if(name=="Apple"){
		return newApple()
	}else if(name=="Samsung"){
        return newSamsung()
	}

	return nil
}