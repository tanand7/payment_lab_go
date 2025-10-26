package PaymentGateway

import (
	"errors"
	"fmt"
)

type CreditCard struct {
	CardNumber string // 16 length string
	CVV        string // 3 length string
	ExpiryDate string // MM/YY
}

func (card CreditCard) IsValidCreditCard() (bool, error) {
	if len(card.CardNumber) != 16 {
		return false, errors.New("card number must be 16 digits")
	}
	if len(card.CVV) != 3 {
		return false, errors.New("CVV must be 3 digits")
	}
	if len(card.ExpiryDate) != 5 {
		return false, errors.New("expiry date must be in MM/YY format")
	}
	if card.ExpiryDate[2] != '/' {
		return false, errors.New("expiry date must be in MM/YY format")
	}

	return true, nil
}

func (card *CreditCard) ReadCardDetails() {
	fmt.Println("Enter the card number:")
	fmt.Scanln(&card.CardNumber)
	fmt.Println("Enter the CVV:")
	fmt.Scanln(&card.CVV)
	fmt.Println("Enter the expiry date:")
	fmt.Scanln(&card.ExpiryDate)
}
