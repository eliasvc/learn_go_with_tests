package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(50))
		wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(40)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
