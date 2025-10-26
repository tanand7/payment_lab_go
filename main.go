package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// PaymentMethod is a type that represents the method of payment
type PaymentMethod int

const (
	creditCard     PaymentMethod = 1
	paypal         PaymentMethod = 2
	cryptocurrency PaymentMethod = 3
)

type CreditCard struct {
	CardNumber string    // 16 length string
	CVV        string    // 3 length string
	ExpiryDate time.Time // 10 length string
}

type PayPal struct {
	Email     string
	AuthToken string // 10 length string
}

type Cryptocurrency struct {
	WalletAddress string // 10 length string
}

type Transaction struct {
	ID            int
	Amount        float64
	Status        string        // Success, Failed, Pending
	Method        PaymentMethod // Credit Card, PayPal, Cryptocurrency
	Customer      Customer
	Type          string // Payment, Refund
	Description   string // Description of the transaction
	MethodDetails any
}

type PaymentGatewayAccount struct {
	Balance      float64
	Transactions []Transaction
}

type Customer struct {
	ID    int
	Name  string
	Phone string
}

var (
	gatewayAccount PaymentGatewayAccount
	reader         = bufio.NewReader(os.Stdin)
)

// find a random status from Success, Failed, Pending
var status = []string{"Success", "Failed", "Pending"}

func main() {
	for {
		displayMenu()
		var choice int
		fmt.Scanln(&choice)
		switch choice {
		case 1:
			customer := readCustomerDetails()
			transaction := customer.makePayment()
			gatewayAccount.Transactions = append(gatewayAccount.Transactions, transaction)
			gatewayAccount.Balance += transaction.Amount
			fmt.Println("Gateway account balance: ", gatewayAccount.Balance)
		case 2:
			requestRefund()
		case 3:
			viewTransactionHistory()
		case 4:
			return
		}
	}
}

func displayMenu() {
	fmt.Println("\n\n\n------ Welcome to Thoughtline payment gateway! ------")
	fmt.Println("\n\n\n------ Menu --------")
	fmt.Println("1. Make a payment")
	fmt.Println("2. Request a refund")
	fmt.Println("3. View transaction history")
	fmt.Println("4. Exit")
}

func (customer Customer) makePayment() Transaction {

	fmt.Println("Enter the amount:")
	var amount float64
	fmt.Scanln(&amount)

	fmt.Println("Select Payment method:")
	fmt.Println("1. Credit Card")
	fmt.Println("2. PayPal")
	fmt.Println("3. Cryptocurrency")
	var method PaymentMethod
	fmt.Scanln(&method)

	// Generating random status for transaction - for easiness
	transactionStatus := status[rand.Intn(len(status))]
	description := "Payment made successfully"
	var methodDetails any
	switch method {
	case creditCard:
		fmt.Println("Enter the credit card number:")
		var creditCard CreditCard
		fmt.Scanln(&creditCard.CardNumber)
		fmt.Println("Enter the CVV:")
		fmt.Scanln(&creditCard.CVV)
		fmt.Println("Enter the expiry date:")
		fmt.Scanln(&creditCard.ExpiryDate)
		if !creditCard.IsValidCreditCard() {
			transactionStatus = "Failed"
			description = "Invalid credit card details. Please try again"
			fmt.Println(description)
		}
		methodDetails = creditCard
	case paypal:
		var paypal PayPal
		paypal.Email = readNonEmptyString("Enter the email:")
		paypal.AuthToken = readNonEmptyString("Enter the authentication token:")
		methodDetails = paypal
		if !paypal.IsValidPayPal() {
			transactionStatus = "Failed"
			description = "Invalid PayPal details. Please try again"
			fmt.Println(description)
		}
	case cryptocurrency:
		var cryptocurrency Cryptocurrency
		cryptocurrency.WalletAddress = readNonEmptyString("Enter the wallet address:")
		methodDetails = cryptocurrency
		if !cryptocurrency.IsValidCrypto() {
			transactionStatus = "Failed"
			description = "Invalid cryptocurrency details. Please try again"
			fmt.Println(description)
		}
	default:
		transactionStatus = "Failed"
		description = "Invalid payment method. Please try again"
		fmt.Println(description)
	}

	transaction := Transaction{
		ID:            rand.Intn(1000000),
		Amount:        amount,
		Status:        transactionStatus,
		Method:        method,
		Customer:      customer,
		Type:          "Payment",
		Description:   description,
		MethodDetails: methodDetails,
	}

	if transactionStatus == "Success" {
		fmt.Println("Your transaction is successful. Your transaction ID is ", transaction.ID)
	} else {
		fmt.Println("Your transaction has been failed. If any amount has been deducted, it will be refunded within 24 hours. Your transaction ID is ", transaction.ID)
	}

	return transaction
}

func requestRefund() {

	fmt.Println("Not implemented yet")
	return

	fmt.Println("Enter the transaction ID:")
	var transactionID int
	fmt.Scanln(&transactionID)

	// TODO: Implement refund logic
}

func viewTransactionHistory() {

	fmt.Println("\n\n\n ------ All Transactions ------")
	totalTransactions := len(gatewayAccount.Transactions)
	if totalTransactions == 0 {
		fmt.Println("No transactions found")
		return
	}

	fmt.Println("\nTotal transactions: ", totalTransactions)

	for index, transaction := range gatewayAccount.Transactions {
		fmt.Println("\nTransaction ", index+1, ":")
		fmt.Println("Transaction ID: ", transaction.ID)
		fmt.Println("Amount: ", transaction.Amount)
		fmt.Println("Status: ", transaction.Status)
		fmt.Println("Customer Name: ", transaction.Customer.Name)
		fmt.Println("Customer Phone: ", transaction.Customer.Phone)
		fmt.Println("Type: ", transaction.Type)
		fmt.Println("Description: ", transaction.Description)
		fmt.Println("Method: ", transaction.Method)
		fmt.Println("Method Details: ", transaction.MethodDetails)
	}
	fmt.Println("--------------------------------")
}

func readCustomerDetails() Customer {
	name := readNonEmptyString("Enter name:")
	phone := readNonEmptyString("Enter phone:")
	return Customer{Name: name, Phone: phone, ID: rand.Intn(1000000)}
}

/// Common Methods

func readLine() string {
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func readNonEmptyString(prompt string) string {

	for {
		fmt.Println(prompt)

		line := readLine()
		if line != "" {
			return line
		}
		fmt.Println("Value cannot be empty. Please try again")
	}

}

func (c CreditCard) IsValidCreditCard() bool {
	return len(c.CardNumber) == 16 && len(c.CVV) == 3 && c.ExpiryDate.After(time.Now())
}

func (c Cryptocurrency) IsValidCrypto() bool {
	return len(c.WalletAddress) >= 10
}

func (p PayPal) IsValidPayPal() bool {
	return len(p.AuthToken) >= 10 && strings.Contains(p.Email, "@")
}
