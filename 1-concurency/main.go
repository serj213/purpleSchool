package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func generateNum()int {
	rand.NewSource(time.Now().UnixNano())
	min := 1
	max := 100

	return rand.Intn(max + min - 1) + min
}

func main(){
	wg := sync.WaitGroup{}
	randNumCh := make(chan int, 10)
	squareCh := make(chan int, 10)
	result := []int{}
	wg.Add(1)
	go func ()  {
		defer wg.Done()
		countNumbers := 10
		for i := 0; i < countNumbers; i++ {
			randomNum := generateNum()
			randNumCh <- randomNum
		}
		close(randNumCh)
	}()
	wg.Add(1)
	go func(){
		defer wg.Done()
		for num := range randNumCh {
			squareCh <- num * num
		}
		close(squareCh)
		
	}()

	go func() {
        wg.Wait()       
    }()

	for squareItem := range squareCh{
		result = append(result, squareItem)
	}
	
	

	fmt.Println("finish result ", result)

}