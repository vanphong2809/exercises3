package main

import (
	"bufio"
	"log"
	"os"
	"sync"
)
type Line struct {
	line_number int
	value       string
}

func worker(result chan *Line, wg *sync.WaitGroup) {
	for {
		select {
		case data := <-result:
			log.Printf("%d giá trị là: %s xong!", data.line_number, data.value)
			wg.Done() //Bao cho goroutine chinh la goroutine nay da chay xong
		}
	}
}

//Bai4 workerpool
func Bai4() {
	var wg sync.WaitGroup
	result := make(chan *Line, 10)
	fileIO, _ := os.Open("file.txt")
	scanner := bufio.NewScanner(fileIO)
	for i := 1; i <= 3; i++ {
		go worker(result, &wg)
	}
	j := 1
	for scanner.Scan() {
		data := Line{line_number: j, value: scanner.Text()}
		result <- &data
		j++
		wg.Add(1) //Them 1 goroutine can doi
	}
	wg.Wait() //Doi den khi nao ma tat ca cac goroutine chay xong
}
