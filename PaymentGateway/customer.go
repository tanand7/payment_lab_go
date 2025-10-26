package PaymentGateway

import (
	"math/rand"
	helpers "payment_lab/Helpers"
)

type Customer struct {
	ID    int
	Name  string
	Phone string
}

func ReadCustomerDetails() Customer {
	name := helpers.ReadNonEmptyString("Enter name:")
	phone := helpers.ReadNonEmptyString("Enter phone:")
	return Customer{Name: name, Phone: phone, ID: rand.Intn(1000000)}
}
