// Place #1 - Basic Goroutine

package golang_goroutine_tour

import (
	"fmt"
	"time"
)

func basicGoroutine() {

	// Creating goroutine from named function
	// go producer(1)
	// producer(2) // This alone executed

	// producer(1) // This alone executed
	// go producer(2)

	// Both producers were not executed. Reason, main function returned bafore producer execution
	// go producer(1)
	// go producer(2)

	go producer(1)
	go producer(2)

	// Creating goroutine from ananymous function
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Func() - Message # %d\n", i)
		}
		producer(3)
	}()
	// Give goroutines time to complete work
	time.Sleep(1 * time.Millisecond)
}

func producer(id int) {
	for i := 0; i < 5; i++ {
		fmt.Printf("Producer # %d - Message # %d\n", id, i)
	}
}
