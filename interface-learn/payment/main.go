package main

import "fmt"

type Payment interface {
	Pay(amount int)
}

type OVO struct {}
type Dana struct {}
type Gopay struct {}

func (OVO) Pay(amount int) {
	fmt.Println("Paying with OVO");
}

func (Dana) Pay(amount int) {
	fmt.Println("Paying with Dana");
}

func (Gopay) Pay(amount int) {
	fmt.Println("Paying with Gopay");
}

func Paying(p Payment, amount int) {
	p.Pay(amount);
}

func main(){
	var inp int64
	var amount int64
	
	for true {
		fmt.Println("Payment methods");
		fmt.Println("1. OVO");
		fmt.Println("2. Dana");
		fmt.Println("3. Gopay");
		fmt.Print("Input: ");
		fmt.Scan(&inp);
		
		if inp > 3 || inp < 1 {
			fmt.Println("Invalid");
			continue
		}

		for true {
			fmt.Print("Amount: ");
			fmt.Scan(&amount);

			if amount > 0 {
				break
			}
			
			fmt.Println("Invalid amount!")
		}
		break
	}

	var payment Payment;
	switch inp {
	case 1: payment = OVO{}
	case 2: payment = Dana{}
	case 3: payment = Gopay{}
	}

	Paying(payment, int(amount)) 
}
