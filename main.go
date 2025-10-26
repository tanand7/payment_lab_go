package main

import paymentGateway "payment_lab/PaymentGateway"

var (
	gatewayAccount paymentGateway.GatewayAccount
)

func main() {
	StartPaymentGateway(&gatewayAccount)
}
