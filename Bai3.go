package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func errFunc() {
	var mutex sync.Mutex
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mutex.Lock()
				_, ok := m[j]
				mutex.Unlock()
				if ok {
					mutex.Lock()
					delete(m, j)
					mutex.Unlock()
					continue
				}
				mutex.Lock()
				m[j] = j * 10

				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println(m)
	log.Print("done")
}
