package paymentGateway

import (
	"errors"
	"fmt"
	helpers "payment_lab/Helpers"
)

type CryptoCurrency struct {
	WalletAddress string // 10 length string
}

func (cryptoCurrency CryptoCurrency) IsValidCryptoCurrency() (bool, error) {
	if len(cryptoCurrency.WalletAddress) < 10 {
		fmt.Println("Wallet address must be at least 10 characters long")
		return false, errors.New("wallet address must be at least 10 characters long")
	}
	return true, nil
}

func (cryptoCurrency *CryptoCurrency) ReadCryptoCurrencyDetails() {
	cryptoCurrency.WalletAddress = helpers.ReadNonEmptyString("Enter the wallet address:")
	fmt.Scanln(&cryptoCurrency.WalletAddress)
}
