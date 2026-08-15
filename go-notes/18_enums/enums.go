package main

import "fmt"

// custom types
// type Mytype string

// enums are enumerated types

// type OrderStatus int

// const (
// 	Received OrderStatus = iota //iota is an auto increamented number starting from 0
// 	Confirmed
// 	Prepared
// 	Delivered 
// )

// func changeStatus(status OrderStatus){
// 	fmt.Println("Changing Status to", status)
// }



type OrderStatus string

const (
	Received OrderStatus = "received"
	Confirmed = "confirmed"
	Prepared = "prepared"
	Delivered = "delivered"
)

func changeOrderStatus(status OrderStatus){
	fmt.Println("Changing order status", status)
}


func main(){
	// changeStatus(Prepared)
	changeOrderStatus(Confirmed)
}
