package shaping

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

//----------------------------------------------------------------------------------------------------------------------------//

func ___Test1(t *testing.T) {
	s := New(2)

	n := 10

	wg := new(sync.WaitGroup)
	wg.Add(n)

	for i := 0; i < n; i++ {
		i := i
		go func() {
			defer wg.Done()

			s.In()
			defer s.Out()

			fmt.Printf("%d\n", i)
			time.Sleep(time.Second)
		}()
	}

	wg.Wait()
}

//----------------------------------------------------------------------------------------------------------------------------//
