// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package bank_test

import (
	"fmt"
	"sync"
	"testing"

	. "github.com/YoshikiShibata/gpl/ch09/ex01/bank"
)

func TestWithdrawNormal(t *testing.T) {
	Deposit(200)
	ok := Withdraw(100)
	if !ok {
		t.Error("Result is false, want true")
		return
	}
	ok = Withdraw(100)
	if !ok {
		t.Error("Result is false, want true")
		return
	}
	ok = Withdraw(100)
	if ok {
		t.Error("Result is true, want false")
		return
	}
	if Balance() != 0 {
		t.Errorf("Result is %d, want 0", Balance())
	}
}

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
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
	balance := Balance()
	if balance > 0 {
		_ = Withdraw(balance)
	}
	balance = Balance()
	if balance != 0 {
		t.Errorf("Balance is %d, but want 0", balance)
		return
	}

	Deposit(100)

	readyGo := make(chan struct{})
	result := make(chan bool)
	var wg sync.WaitGroup
	const numOfGoroutines = 10
	wg.Add(numOfGoroutines)
	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			wg.Done()
			<-readyGo
			result <- Withdraw(60)
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
