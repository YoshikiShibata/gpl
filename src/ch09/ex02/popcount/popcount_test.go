// Copyright © 2016 Yoshiki Shibata. All rights reserved.

package popcount_test

import (
	"ch02/ex03/popcount"
	"fmt"
	"sync"
	"testing"
)

func TestInitPCOnce(t *testing.T) {
	done := make(chan struct{})

	var wg sync.WaitGroup

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			<-done
			count := popcount.PopCount(0x12345678901234)
			if count != 20 {
				panic(fmt.Sprintf("count is %d, want 20\n", count))
			}
			wg.Done()
		}()
	}
	close(done)
	wg.Wait()
}
