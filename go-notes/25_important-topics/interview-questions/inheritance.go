// There is no inheritance present in Go but we can achieve it using composition

package main

import "fmt"

type person struct{
	name string
	age int
	role string
}

func (p *person) intro(){
	fmt.Printf("Employee name is %s", p.name)
}

type employee struct{
	*person           // if pointer is asked in interview then it will still be behaving same but we just need to pass address of type person
	employeeId int
}


func main(){
	per := person{
		name:"karan",
		age:23,
		role:"Software engineering intern - Platform",
	}

	emp := employee{ 
		person:&per,
		employeeId: 121,
	}
	per.intro()
	emp.intro()

}