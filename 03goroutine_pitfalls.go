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

func GoroutinePitfalls() {
	// case - 1 -> Reach deadlock. Because, Incorrect goroutine count initialization.
	// wg.Add(2)
	// go producer(1)
	// wg.Wait()

	// case - 2 -> Goroutine n't executed. Because, goroutine not initialized in wait group.
	// go producer(1)
	// wg.Wait()

	// case - 3 -> One goroutine n't executed at any case[{1, 2}, {2, 3}, {3, 1}].
	// wg.Add(2)
	// go producer(1)
	// go producer(2)
	// go producer(3)
	// wg.Wait()

	// case - 4 -> Through panic. Because, WaitGroup counter is negative.
	// wg.Add(2)
	// go producer(1)
	// go producer(2)
	// wg.Wait()
	// producer(3) // This is contain wg.Done() function, that is decrease the counter value -1.

	// Best to way to handle wait group pitfalls
	// producer1(1)
	// producer1(2)
	// producer1(3)
	// wg.Wait()

	// Launching goroutines from anonymous functions
	// launchWorkers(5)
	// wg.Wait()

	producer2(wg, 4) // wg act like a deep copy. So, goroutine not triggered.
	wg.Wait()
}

func producer2(wg sync.WaitGroup, id int) {
	wg.Add(1)
	go func() {
		// n is a random int between 1 to 1000 inclusive
		n := (r.Int() % 1000) + 1
		d := time.Duration(n) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("Producer1 # %d ran for %v\n", id, d)
		defer wg.Done()
	}()
}

func launchWorkers(workers int) {
	for w := 0; w < workers; w++ {
		wg.Add(1)
		id := w
		go func() {
			fmt.Printf("I am worker # %d\n", id)
			wg.Done()
		}()
	}
}

func producer1(id int) {
	wg.Add(1)
	go func() {
		// n is a random int between 1 to 1000 inclusive
		n := (r.Int() % 1000) + 1
		d := time.Duration(n) * time.Millisecond
		time.Sleep(d)
		fmt.Printf("Producer1 # %d ran for %v\n", id, d)
		defer wg.Done()
	}()
}

func producer(id int) {
	// n is a random int between 1 to 1000 inclusive
	n := (r.Int() % 1000) + 1
	d := time.Duration(n) * time.Millisecond
	time.Sleep(d)
	fmt.Printf("Producer1 # %d ran for %v\n", id, d)
	defer wg.Done()
}
