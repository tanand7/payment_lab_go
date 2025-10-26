package PaymentGateway

import (
	"errors"
	helpers "payment_lab/Helpers"
	"strings"
)

type PayPal struct {
	Email     string
	AuthToken string
}

func (paypal PayPal) IsValidPayPal() (bool, error) {
	if len(paypal.AuthToken) < 10 {
		return false, errors.New("authentication token must be at least 10 characters long")
	}
	if !strings.Contains(paypal.Email, "@") {
		return false, errors.New("email must contain @")
	}
	return true, nil
}

func (paypal *PayPal) ReadPaypalDetails() {
	paypal.Email = helpers.ReadNonEmptyString("Enter the email:")
	paypal.AuthToken = helpers.ReadNonEmptyString("Enter the authentication token:")
}
