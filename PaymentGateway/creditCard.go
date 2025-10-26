package paymentGateway

import "fmt"

type CreditCard struct {
	CardNumber string // 16 length string
	CVV        string // 3 length string
	ExpiryDate string // MM/YY
}

func (card CreditCard) IsValidCreditCard() bool {
	if len(card.CardNumber) != 16 {
		fmt.Println("Card number must be 16 digits")
		return false
	}
	if len(card.CVV) != 3 {
		fmt.Println("CVV must be 3 digits")
		return false
	}
	if len(card.ExpiryDate) != 5 {
		fmt.Println("Expiry date must be in MM/YY format")
		return false
	}
	if card.ExpiryDate[2] != '/' {
		fmt.Println("Expiry date must be in MM/YY format")
		return false
	}

	return true
}

func (card *CreditCard) ReadCardDetails() {
	fmt.Println("Enter the card number:")
	fmt.Scanln(&card.CardNumber)
	fmt.Println("Enter the CVV:")
	fmt.Scanln(&card.CVV)
	fmt.Println("Enter the expiry date:")
	fmt.Scanln(&card.ExpiryDate)
}
