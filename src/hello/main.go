package main

import (
	"fmt"
)

func main() {
	s := "hello"
	//	var b bool
	fmt.Println(s)
	a, e := fmt.Printf("%v \n", s)
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", Data{"lengxin", 19})
}

type Data struct {
	name string
	age  int
}

func (self Data) String() string {
	return fmt.Sprintf("%v, %d", self.name, self.age)
}
