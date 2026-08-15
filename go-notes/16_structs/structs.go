package main

import (
	"fmt"
	"time"
)


type customer struct{
	name string
	phone string
}


// embedding (customer is embedded in order struct)
type order struct{
	id string
	amount float32
	status string
	createdat time.Time
	customer
}

// constructor function
func newOrder(id string, amount float32, status string) *order {
	myOrder := order{
		id: id,
		amount: amount,
		status: status,
	}
	return &myOrder
}

// receive type 
// methods
func (o *order) setStatus(status string){
	o.status = status
}

// getting value not modifying so does not matter if done by copy or actual address
func (o order) getAmount() float32 {
	return  o.amount
}


func main(){

	// If you don't set any field, default value is zero value
	// int -> 0, float->0, string -> "", bool-> false
	// myOrder := order{
	// 	id: "1",
	// 	amount: 323.00,
	// 	status: "received",
	// }

	// via contructor
	myOrder := newOrder("1",323.00,"received")
	myOrder.setStatus("confirmed")
	fmt.Println(myOrder)

	fmt.Println(myOrder.getAmount())
	// myOrder.createdat = time.Now()

	// myOrder2 := order{
	// 	id: "2",
	// 	amount: 43.00,
	// 	status: "delivered",
	// 	createdat: time.Now(),
	// }

	// // fmt.Println(myOrder.status)
	// // fmt.Println(myOrder)
	// fmt.Println(myOrder2)	



	// if we have to use struct only one time, so we can create a struct without name and assign values right after that
	language := struct{
		name string
		isGood bool
	}{"golang", true}

	fmt.Println(language)

	newCustomer := customer{
		name: "karan",
		phone: "1023456789",
	}

	myOrder2 := order{
		id: "1",
		amount: 323.00,
		status: "received",
		customer: customer{
			name: "karan",
			phone: "1023456789",
		},
		
	}	
}



