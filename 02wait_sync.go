// Place # 2 - Waiting & Synchronization
package golang_goroutine_tour

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	s  = rand.NewSource(time.Now().Unix())
	r  = rand.New(s)
	wg sync.WaitGroup
)

func WaitSync() {

	// non-ideal way to wait for all goroutines to complete
	start := time.Now()
	go producer(1)
	go producer(2)

	// Give goroutines time to complete work
	time.Sleep(1 * time.Second)
	elapse := time.Now().Sub(start)
	fmt.Printf("Non-idea wait took %v\n", elapse)

	// ideal way to wait for all goroutines to complete
	start = time.Now()
	wg.Add(2)
	go producer2(1)
	go producer2(2)

	// Give goroutines time to complete work
	wg.Wait()
	elapse = time.Now().Sub(start)
	fmt.Printf("Non-idea wait took %v\n", elapse)

	// pipeline
	start = time.Now()
	wg.Add(1)
	go producer2(1)

	// Give goroutines time to complete work
	wg.Wait()
	elapse = time.Now().Sub(start)
	fmt.Printf("Pipeline non-idea wait took %v\n", elapse)
}

func producer2(id int) {
	// n is a random int between 1 to 1000 inclusive
	n := (r.Int() % 1500) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer2 # %d ran for %v\n", id, d)
	defer wg.Done()
}

func producer(id int) {
	// n is a random int between 1 to 1000 inclusive
	n := (r.Int() % 1500) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer1 # %d ran for %v\n", id, d)
}
