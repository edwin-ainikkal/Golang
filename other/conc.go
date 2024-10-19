package main

import (
	"fmt"

	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)

	// Wait for both goroutines to finish
	wg.Wait()
}

func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i < 6; i++ {
		fmt.Println(i)
	}
	wg.Done()
}

func printLetters(wg *sync.WaitGroup) {

	for _, word := range words {
		fmt.Println(word)
	}
	wg.Done()

}
