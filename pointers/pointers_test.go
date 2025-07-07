package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, &wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		withdraw := Bitcoin(30)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(withdraw)

		fmt.Println(ErrEasterEgg.Error())
		assertError(t, err, ErrInsufficientFundsOnWithdraw(withdraw, startingBalance))
		assertBalance(t, &wallet, startingBalance)
	})
}

func assertBalance(t testing.TB, w *Wallet, want Bitcoin) {
	t.Helper()

	got := (*w).Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("encountered an error when there shouldn't have been one.")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		// t.Fatal exits the test
		t.Fatal("wanted an error but did not get one.")
	}

	// got.Error() MUST happen only when error is not nil
	// we guarantee this in the if block above this by using t.Fatal
	if got.Error() != want.Error() {
		t.Errorf("got %q, want %q", got, want)
	}
}
