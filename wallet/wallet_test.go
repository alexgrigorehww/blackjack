package wallet_test

import (
	"blackjack/wallet"
	"testing"
)

func TestWallet_LostMoneyValue(t *testing.T) {
	w := new(wallet.Wallet)
	w.SetAmount(100)
	if w.LostMoney(50) != 50 {
		t.Error("Remaining amount should be 50")
	}
}

func TestWallet_SetAmount(t *testing.T) {
	w := new(wallet.Wallet)
	w.SetAmount(500)
	if w.GetAmount() != 500 {
		t.Error("Amount on set should be 500")
	}
}

func TestWallet_GetAmount(t *testing.T) {
	w := new(wallet.Wallet)
	w.SetAmount(500)
	if w.GetAmount() != 500 {
		t.Error("Amount on get should be 500")
	}
}
