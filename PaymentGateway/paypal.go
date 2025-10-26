package paymentGateway

import (
	"fmt"
	helpers "payment_lab/Helpers"
	"strings"
)

type PayPal struct {
	Email     string
	AuthToken string
}

func (paypal PayPal) IsValidPayPal() bool {
	if len(paypal.AuthToken) < 10 {
		fmt.Println("Authentication token must be at least 10 characters long")
		return false
	}
	if !strings.Contains(paypal.Email, "@") {
		fmt.Println("Email must contain @")
		return false
	}
	return true
}

func (paypal *PayPal) ReadPayPalDetails() {
	paypal.Email = helpers.ReadNonEmptyString("Enter the email:")
	paypal.AuthToken = helpers.ReadNonEmptyString("Enter the authentication token:")
}
