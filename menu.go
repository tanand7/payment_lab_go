package main

import (
	"fmt"
	paymentGateway "payment_lab/PaymentGateway"
)

func StartPaymentGateway(gatewayAccount *paymentGateway.GatewayAccount) {
	for {
		displayMenu()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			gatewayAccount.InitializeTransaction()
		case 2:
			transactionID := paymentGateway.ReadTransactionID()
			gatewayAccount.RequestRefund(transactionID)
		case 3:
			gatewayAccount.ViewTransactionHistory()
		case 4:
			return
		}
	}
}

func displayMenu() {
	fmt.Println("\n\n\n------ Menu --------")
	fmt.Println("1. Make a payment")
	fmt.Println("2. Request a refund")
	fmt.Println("3. View transaction history")
	fmt.Println("4. Exit")
}
