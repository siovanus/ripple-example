package main

import (
	"github.com/rubblelabs/ripple/crypto"
	"github.com/rubblelabs/ripple/data"
)

func GeneratePayment(from, to data.Account, amount data.Amount, fee data.Value, sequence uint32) *data.Payment {
	payment := &data.Payment{
		Destination: to,
		Amount:      amount,
	}
	txBase := data.TxBase{
		TransactionType: data.PAYMENT,
		Account:         from,
		Sequence:        sequence,
		Fee:             fee,
	}
	payment.TxBase = txBase
	return payment
}

func main() {
	secret := "shtew2z1TRsEvpnYUGtiyvqPnYywt"
	accountSeedContent, err := crypto.NewRippleHash(secret)
	if err != nil {
		panic(err)
	}
	accountSeed := accountSeedContent.Payload()
	accountKey, err := crypto.NewECDSAKey(accountSeed)
	if err != nil {
		panic(err)
	}
	key := accountKey

	to, _ := data.NewAccountFromAddress("rT4vRkeJsgaq7t6TVJJPsbrQp5oKMGRfN")
	from, _ := data.NewAccountFromAddress("rsHYGX2AoQ4tXqFywzEeeTDgXFTUfL1Fw9")
	amount, _ := data.NewAmount("13/XRP")
	fee, _ := data.NewValue("0.00005", true)
	payment := GeneratePayment(*from, *to, *amount, *fee, 25336389)

	data.Sign(payment, key, nil)
}
