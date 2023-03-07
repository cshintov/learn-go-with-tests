package pointers

import "testing"

func TestWallet(t *testing.T) {
    t.Run("deposit", func(t *testing.T){
        wallet := Wallet{}

        wallet.Deposit(Bitcoin(10))
        wallet.Deposit(Bitcoin(20))

        assertBalance(t, wallet, Bitcoin(30))
    })

    t.Run("withdraw", func(t *testing.T){
        wallet := Wallet{}

        wallet.Deposit(Bitcoin(10))
        wallet.Deposit(Bitcoin(20))
        err := wallet.Withdraw(Bitcoin(15))

        want := Bitcoin(15)

        assertNoError(t, err)
        assertBalance(t, wallet, want)
    })

    t.Run("withdraw insufficient funds", func(t *testing.T) {
        startingBalance := Bitcoin(20)
        wallet := Wallet{startingBalance}
        err := wallet.Withdraw(Bitcoin(100))

        assertError(t, err, "cannot withdraw, insufficient funds")
        assertBalance(t, wallet, startingBalance)
    })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()

    got := wallet.Balance()

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertError(t testing.TB, got error, want string) {
    t.Helper()

    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got.Error() != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

func assertNoError(t testing.TB, got error) {
    t.Helper()

    if got != nil {
        t.Fatal("got an error but didn't want one")
    }
}
