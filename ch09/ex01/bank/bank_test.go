// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package bank_test

import (
	"fmt"
	"testing"

	"github.com/YoshikiShibata/gpl/ch09/ex01/bank"
)

func TestWithdrawNormal(t *testing.T) {
	bank.Deposit(200)
	ok := bank.Withdraw(100)
	if !ok {
		t.Error("Result is false, want true")
		return
	}
	ok = bank.Withdraw(100)
	if !ok {
		t.Error("Result is false, want true")
		return
	}
	ok = bank.Withdraw(100)
	if ok {
		t.Error("Result is true, want false")
		return
	}
	if bank.Balance() != 0 {
		t.Errorf("Result is %d, want 0", bank.Balance())
	}
}

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
