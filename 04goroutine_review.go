package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	WORKERS_COUNT = 10000 // 1000000 also worked
)

var (
	s       = rand.NewSource(time.Now().Unix())
	r       = rand.New(s)
	wg      sync.WaitGroup
	counter = 0
	mtx     sync.Mutex
	idMap   = make(map[int]int)
)

func main() {
	for w := 0; w < 50; w++ {
		initWorker(w)
	}
	wg.Wait()
	fmt.Printf("Counter : %d\n", counter)
	fmt.Printf("ID map length: %d value: %v\n", len(idMap), idMap)
}

func initWorker(id int) {
	wg.Add(1)
	go func() {
		time.Sleep(time.Duration((r.Int()%1000)+1) * time.Nanosecond)
		// Handling concurrent data access through sync.Mutex
		mtx.Lock()
		counter++
		idMap[id] = idMap[id] + 1
		mtx.Unlock()
		defer wg.Done()
	}()
}
