package paymentGateway

import (
	"fmt"
	"math/rand"
	"strconv"
)

type PaymentMethod string

const (
	creditCard     PaymentMethod = "Credit Card"
	paypal         PaymentMethod = "PayPal"
	cryptoCurrency PaymentMethod = "Crypto Currency"
)

type GatewayAccount struct {
	Balance      float64
	Transactions []Transaction
}

type PaymentGateway interface {
	InitializeTransaction() Transaction
	ReadAmount() float64
	ReadPaymentMethod() PaymentMethod
	MakePayment(amount float64, method PaymentMethod, customer Customer) Transaction
	ReadTransactionID() int
	RequestRefund(transactionID int)
	ViewTransactionHistory()
}

func (gatewayAccount *GatewayAccount) InitializeTransaction() {


	
	customer := readCustomerDetails()
	transaction := customer.makePayment()
	gatewayAccount.Transactions = append(gatewayAccount.Transactions, transaction)
	if transaction.Status == "Success" {
		gatewayAccount.Balance += transaction.Amount
	}
	fmt.Println("Gateway account balance: ", gatewayAccount.Balance)

}

func ReadAmount() float64 {
	fmt.Println("Enter the amount:")
	var amount float64
	fmt.Scanln(&amount)
	return amount
}

func ReadPaymentMethod() PaymentMethod {
	fmt.Println("Select Payment method:")
	fmt.Println("1. Credit Card")
	fmt.Println("2. PayPal")
	fmt.Println("3. Cryptocurrency")
	var method PaymentMethod
	fmt.Scanln(&method)
	return method
}

func (customer Customer) MakePayment() Transaction {

	// Generating random status for transaction - for easiness
	transactionStatus := "Success"
	description := "Payment made successfully"
	var methodDetails any
	switch method {
	case creditCard:
		// TODO: Replace with the new method
		// fmt.Println("Enter the credit card number:")
		// var creditCard CreditCard
		// fmt.Scanln(&creditCard.CardNumber)
		// fmt.Println("Enter the CVV:")
		// fmt.Scanln(&creditCard.CVV)
		// fmt.Println("Enter the expiry date:")
		// fmt.Scanln(&creditCard.ExpiryDate)
		// if !creditCard.IsValidCreditCard() {
		// 	transactionStatus = "Failed"
		// 	description = "Invalid credit card details. Please try again"
		// 	fmt.Println(description)
		// }
		methodDetails = creditCard
	case paypal:
		// TODO: Replace with the new method
		// var paypal PayPal
		// paypal.Email = readNonEmptyString("Enter the email:")
		// paypal.AuthToken = readNonEmptyString("Enter the authentication token:")
		// methodDetails = paypal
		// if !paypal.IsValidPayPal() {
		// 	transactionStatus = "Failed"
		// 	description = "Invalid PayPal details. Please try again"
		// 	fmt.Println(description)
		// }
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

// Reads the transaction ID from the user
func ReadTransactionID() int {
	fmt.Println("Enter the transaction ID:")
	var transactionID int
	fmt.Scanln(&transactionID)
	return transactionID
}

// Requests a refund for a given transaction ID
func (gatewayAccount *GatewayAccount) RequestRefund(transactionID int) {

	for _, transaction := range gatewayAccount.Transactions {
		if transaction.ID == transactionID {

			refundTransaction := Transaction{
				ID:            rand.Intn(1000000),
				Amount:        transaction.Amount,
				Status:        "Refunded",
				Method:        transaction.Method,
				Customer:      transaction.Customer,
				Type:          "Refund",
				Description:   "Refunded for transaction ID " + strconv.Itoa(transaction.ID),
				MethodDetails: transaction.MethodDetails,
			}

			gatewayAccount.Balance -= transaction.Amount
			gatewayAccount.Transactions = append(gatewayAccount.Transactions, refundTransaction)
			fmt.Println("Refunded amount: ", transaction.Amount, " successfully to your original source")
			return
		}
	}
	fmt.Println("Could not find the transaction with ID: ", transactionID, ". Please try again with a valid transaction ID.")
}

// Prints the transaction history of the gateway account
func (gatewayAccount GatewayAccount) ViewTransactionHistory() {

	fmt.Println("\n\n\n ------ All Transactions ------")
	totalTransactions := len(gatewayAccount.Transactions)
	if totalTransactions == 0 {
		fmt.Println("No transactions found")
		return
	}

	fmt.Println("\nTotal transactions till date: ", totalTransactions)
	fmt.Println("\nGateway account balance: ", gatewayAccount.Balance)

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
