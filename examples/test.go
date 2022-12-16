package main

import "fmt"

type Stu struct {
	name string
	age  uint8
}

func main() {
	stu := make(map[string]*Stu)
	stu["ranen"] = &Stu{"ranen", 12}
	s, ok := stu["ranen1"]
	fmt.Println(s, ok)
}
