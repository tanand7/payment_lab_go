package paymentGateway

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
)

type PaymentMethod int

const (
	creditCard     PaymentMethod = 1
	paypal         PaymentMethod = 2
	cryptoCurrency PaymentMethod = 3
)

type GatewayAccount struct {
	Balance      float64
	Transactions []Transaction
}

type PaymentGateway interface {
	InitializeTransaction() Transaction
	ReadAmount() float64
	ReadPaymentMethod() PaymentMethod
	ReadPaymentSource(method PaymentMethod) (any, error)
	MakePayment(amount float64, method PaymentMethod, customer Customer, source any) Transaction
	ReadTransactionID() int
	RequestRefund(transactionID int)
	ViewTransactionHistory()
}

func (gatewayAccount *GatewayAccount) InitializeTransaction() {

	customer := ReadCustomerDetails()
	method := gatewayAccount.ReadPaymentMethod()
	amount := gatewayAccount.ReadAmount()
	paymentSource, err := gatewayAccount.ReadPaymentSource(method)

	if err != nil {
		fmt.Println("Transaction failed: ", err)
		return
	}

	transaction := gatewayAccount.MakePayment(amount, method, customer, paymentSource)
	gatewayAccount.Transactions = append(gatewayAccount.Transactions, transaction)
	if transaction.Status == "Success" {
		gatewayAccount.Balance += transaction.Amount
	}
	fmt.Println("Gateway account balance: ", gatewayAccount.Balance)
}

func (gatewayAccount GatewayAccount) ReadAmount() float64 {
	fmt.Println("Enter the amount:")
	var amount float64
	fmt.Scanln(&amount)
	return amount
}

func (gatewayAccount GatewayAccount) ReadPaymentMethod() PaymentMethod {
	fmt.Println("Select Payment method:")
	fmt.Println("1. Credit Card")
	fmt.Println("2. PayPal")
	fmt.Println("3. Cryptocurrency")
	var method PaymentMethod
	fmt.Scanln(&method)
	return method
}

func (gatewayAccount GatewayAccount) ReadPaymentSource(method PaymentMethod) (any, error) {
	switch method {
	case creditCard:
		var creditCard CreditCard
		creditCard.ReadCardDetails()
		valid, err := creditCard.IsValidCreditCard()
		if !valid {
			fmt.Println(err)
			return nil, err
		}
		return creditCard, nil
	case paypal:
		var paypal PayPal
		paypal.ReadPaypalDetails()
		valid, err := paypal.IsValidPayPal()
		if !valid {
			fmt.Println(err)
			return nil, err
		}
		return paypal, nil
	case cryptoCurrency:
		var cryptoCurrency CryptoCurrency
		cryptoCurrency.ReadCryptoCurrencyDetails()
		valid, err := cryptoCurrency.IsValidCryptoCurrency()
		if !valid {
			fmt.Println(err)
			return nil, err
		}
		return cryptoCurrency, nil
	default:
		return nil, errors.New("invalid payment method, please try again")
	}
}

func (gatewayAccount *GatewayAccount) MakePayment(amount float64, method PaymentMethod, customer Customer, source any) Transaction {

	transactionStatus := "Success"
	description := "Payment made successfully"

	transaction := Transaction{
		ID:            rand.Intn(1000000),
		Amount:        amount,
		Status:        transactionStatus,
		Method:        method,
		Customer:      customer,
		Type:          "Payment",
		Description:   description,
		MethodDetails: source,
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
