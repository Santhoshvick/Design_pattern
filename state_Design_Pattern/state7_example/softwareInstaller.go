package main

import "fmt"

type Software interface{
	download(*App)
}

type App struct{
	software Software
}

func(a *App)setState(soft Software){
	a.software=soft
}



func(a *App)download(){
	a.software.download(a)
}

type NotInstalled struct{}

func(n *NotInstalled)download(a *App){
	fmt.Println("Download The Applciation  by clicking the Below ")
	a.setState(&Installing{})
}

type Installing struct{
	installation bool
}

func(i *Installing)download(a *App){
	i.installation=true
	if i.installation{
	fmt.Println("The desktop is started it downloading process")
	a.setState(&Installed{})
	}else{
        a.setState(&Failed{})
	}
}

type Installed struct{
	Software string
	Version string
	OS string
}

func(i1 *Installed)download(a *App){
	fmt.Printf("The softwate %s and the version %s is Downloaded successfully",i1.Software,i1.Version)

}

type Failed struct{}

func(f *Failed)download(a *App){
	fmt.Println("Installation Failed‼️")
	a.setState(nil)
}

func main(){

	t1:=&App{}
	t1.setState(&NotInstalled{})
	soft1:=&Installed{Software: "Teams",Version: "10.0.12"}

	t1.download()
	t1.download()
	soft1.download(t1)
	
}

