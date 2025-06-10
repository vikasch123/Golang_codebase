package main

import "fmt"


type Payment interface {
	Process() string
}

type CreditCard struct{
	CardNumber string
	Amount float64
}

type UPI struct{
	UPiId string
	Amount float64 
}

func (cc CreditCard) Process() string{
	return fmt.Sprintf("Paid %.2f via Credit Card ending in %s", cc.Amount, cc.CardNumber[len(cc.CardNumber)-4:])
}

func (upi UPI) Process() string {
    return fmt.Sprintf("Paid %.2f via UPI ID %s", upi.Amount, upi.UPiId)
}

func ExecutePayment(p Payment) string{
	return p.Process()

}



func main(){
	cc := CreditCard{CardNumber: "1234567890123456", Amount: 500.00}
    upi := UPI{UPiId: "rahul@upi", Amount: 300.00}

    // Execute payments using the interface
    fmt.Println(ExecutePayment(cc))  
	fmt.Println(ExecutePayment(upi))
}
