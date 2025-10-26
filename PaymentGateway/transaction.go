package paymentGateway

type Transaction struct {
	ID            int
	Amount        float64
	Status        string        // Success, Failed
	Method        PaymentMethod // Credit Card, PayPal, Crypto Currency
	Customer      Customer
	Type          string // Payment, Refund
	Description   string // Description of the transaction
	MethodDetails any
}
