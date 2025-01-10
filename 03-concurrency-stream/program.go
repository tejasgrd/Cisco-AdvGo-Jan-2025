package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	dataWg := &sync.WaitGroup{}
	dataCh := make(chan int)

	dataWg.Add(1)
	go Source(dataWg, "data1.dat", dataCh)

	dataWg.Add(1)
	go Source(dataWg, "data2.dat", dataCh)

	evenCh, oddCh := Splitter(dataCh)

	evenSumCh := Sum(evenCh)
	oddSumCh := Sum(oddCh)

	doneCh := Merger(evenSumCh, oddSumCh)

	dataWg.Wait()
	close(dataCh)

	<-doneCh

}

func Source(wg *sync.WaitGroup, fileName string, dataCh chan<- int) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if no, err := strconv.Atoi(line); err == nil {
			dataCh <- no
		}
	}
}

func Splitter(dataCh <-chan int) (<-chan int, <-chan int) {
	evenCh := make(chan int)
	oddCh := make(chan int)
	go func() {
		defer close(evenCh)
		defer close(oddCh)
		for no := range dataCh {
			if no%2 == 0 {
				evenCh <- no
			} else {
				oddCh <- no
			}
		}
	}()
	return evenCh, oddCh
}

func Sum(ch <-chan int) <-chan int {
	sumCh := make(chan int)
	go func() {
		defer close(sumCh)
		var total int
		for no := range ch {
			total += no
		}
		sumCh <- total
	}()
	return sumCh

}

func Merger(evenSumCh, oddSumCh <-chan int) <-chan struct{} {
	doneCh := make(chan struct{})
	go func() {
		file, err := os.Create("result.txt")
		if err != nil {
			log.Fatalln(err)
		}
		defer file.Close()

		for range 2 {
			select {
			case evenTotal := <-evenSumCh:
				fmt.Fprintln(file, fmt.Sprintf("Even Total : %d", evenTotal))
			case oddTotal := <-oddSumCh:
				fmt.Fprintln(file, fmt.Sprintf("Odd Total : %d", oddTotal))
			}
		}
		close(doneCh)
	}()
	return doneCh
}
