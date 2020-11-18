package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

//Bai2 abc
func Bai2() {
	var mutex sync.Mutex
	var number int64
	number = 0
	X := make(map[string]string)
	go func() {
		for i := 0; i < 3; i++ {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				number++
				key := "key" + strconv.FormatInt(number, 10)
				value := "value" + strconv.FormatInt(number, 10)
				X[key] = value
				mutex.Unlock()
			}
		}
	}()
	time.Sleep(5*time.Second)
	fmt.Println(X)
}
