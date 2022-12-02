package main

import "fmt"

type People struct {
	Name string
	Age int
}

func main() {
	Jake := People {"Jake", 31}
	fmt.Println(Jake.Name, "-", Jake.Age)
}
