package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(5)

		assertBalance(t, wallet, Bitcoin(15))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {

		startingFunds := Bitcoin(10)
		wallet := Wallet{startingFunds}

		err := wallet.Withdraw(15)

		assertBalance(t, wallet, startingFunds)
		assertError(t, err, ErrInsufficientFunds)
	})
}

// Helper function for asserting the balance
func assertBalance(t testing.TB, wallet Wallet, exp Bitcoin) {
	t.Helper()

	act := wallet.Balance

	if exp != act {
		t.Errorf("actual: %q, expected: %q", act, exp)
	}
}

// Helper func to assert we get no error
func assertNoError(t testing.TB, act error) {
	t.Helper()

	if act != nil {
		t.Fatal("got an error but wasn't expecting one")
	}
}

// Helper func to assert that we do get an error
func assertError(t testing.TB, exp, act error) {
	t.Helper()

	if act == nil {
		t.Fatal("did not get an error but expected one")
	}

	if exp != act {
		t.Errorf("actual %q, expected %q", act, exp)
	}
}
