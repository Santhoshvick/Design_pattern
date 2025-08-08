package main

import "fmt"

type Document interface {
	Clone() Document
	Print()
}


type Resume struct {
	Name    string
	Details string
}

func (r *Resume) Clone() Document {

	return &Resume{
		Name:    r.Name,
		Details: r.Details,
	}
}

func (r *Resume) Print() {
	fmt.Println("Resume of", r.Name, ":", r.Details)
}

func main() {
	original := &Resume{Name: "xyz", Details: "software tester"}
	copy := original.Clone()

	original.Print()
	copy.Print()
}

