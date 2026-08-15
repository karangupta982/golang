package main

import "fmt"

// an interface is a type that defines a set of method signatures (behaviors), but does not provide the actual implementation for them
// Instead of defining what an object is (like data fields inside a struct), an interface defines what an object can do
type paymenter interface{
	pay(amount float32)
	refund(amount float32, account string)
}

type payment struct {
	gateway paymenter
}

func (p payment) makePayment(amount float32){
	// razorpayGateway := razorpay{}
	// razorpayGateway.pay(amount)
	p.gateway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) {
	fmt.Println("Payment is being processed via razorpay", amount)
}

type stripe struct{}

func (s stripe) pay(amount float32){
	fmt.Println("Making payment using stripe", amount)
}

type fakePayment struct {}

func (f fakePayment) pay(amount float32){
	fmt.Println("Making payment using Fake payment service", amount)
}

func (f fakePayment) refund(amount float32, account string){
	
}

func main(){

	// newPayment := payment{}
	// newPayment.makePayment(100)

	fakePaymentGateway := fakePayment{}
	newPayment := payment{
		gateway: fakePaymentGateway,
	}
	newPayment.makePayment(100)

	 
}