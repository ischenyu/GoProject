package main

import "fmt"

type Person struct{
    Name string
    Age int
}

func (p Person) test() {
    fmt.Println(p.Name, p.Age)
}

func main(){
    var person Person
    person.Name = "Paimon"
    person.Age = 5
    person.test()
}
