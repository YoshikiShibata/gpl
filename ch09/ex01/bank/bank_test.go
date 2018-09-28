// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package bank_test

import (
	"fmt"
	"sync"
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

func TestConcurrentBankAccess(t *testing.T) {
	for i := 0; i < 10000; i++ {
		concurrentBankAccess(t)
		if t.Failed() {
			return
		}
	}
}
func concurrentBankAccess(t *testing.T) {
	balance := bank.Balance()
	if balance > 0 {
		_ = bank.Withdraw(balance)
	}
	balance = bank.Balance()
	if balance != 0 {
		t.Errorf("bank.Balance is %d, but want 0", balance)
		return
	}

	bank.Deposit(100)

	readyGo := make(chan struct{})
	result := make(chan bool)
	var wg sync.WaitGroup
	const numOfGoroutines = 10
	wg.Add(numOfGoroutines)
	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			wg.Done()
			<-readyGo
			result <- bank.Withdraw(60)
		}()
	}
	wg.Wait()
	close(readyGo)

	okCount := 0
	for i := 0; i < numOfGoroutines; i++ {
		if <-result {
			okCount++
		}
	}
	close(result)

	if okCount != 1 {
		t.Errorf("okCount is %d, but want 1", okCount)
	}
}
