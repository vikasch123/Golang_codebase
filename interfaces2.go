package main

import "fmt"

type Speaker interface{
	speak() string
}

type Person struct{
	Name string

}

type Dog struct{
	Name string 
}

func (p Person) speak() string{
	return "Hello , I'm " + p.Name
}

func (d Dog) speak() string{
	return "Woof! I'm " + d.Name 
}

func main(){
	var s Speaker 
	s= Person{Name:"Vikas"}
	fmt.Println(s.speak())
	s= Dog{Name:"Sheru"}
	fmt.Println(s.speak())
}
