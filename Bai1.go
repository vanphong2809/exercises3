package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

//Bai1  dùng kiến thức về go routine và chan đề func dưới in ra đủ 3 message.
func Bai1() {
	//Su dung channel
	//In ra message3 truoc message2
	fmt.Println("In ra message3 truoc message2")
	mychan := make(chan bool)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		<-mychan
	}()
	mychan <- true
	log.Print("hello 2")

	//In ra cac message theo thu tu
	fmt.Println("In ra cac message theo thu tu")
	log.Println("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Println("hello 3")
		<-mychan
	}()
	log.Println("hello 2")
	mychan <- true

	//Su dung mutex
	//In ra message3 truoc message2
	var mutex sync.Mutex
	fmt.Println("Su dung mutex , In ra message3 truoc message2: ")
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		mutex.Unlock()
	}()
	mutex.Lock()
	time.Sleep(2 * time.Second)
	log.Print("hello 2")

	var w sync.WaitGroup
	fmt.Println("Su dung waitGroup , In ra message3 truoc message2: ")
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		w.Done()
	}()
	w.Wait()
	log.Print("hello 2")
}
