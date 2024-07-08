package lib

import (
	"fmt"
	"sync"
)

// push num into chan (queue)
func generatePipleline(nums []int) <-chan int {
	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}

		close(out)
	}()

	return out
}

// handle main business
func fanOut(in <-chan int, name string) <-chan int {
	out := make(chan int)

	go func() {
		for num := range in {
			fmt.Println(name)
			out <- num * num
		}
		close(out)
	}()

	return out
}

// merge out channels
func fanIn(inputChans ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	wg.Add(len(inputChans))

	in := make(chan int)

	for _, inputChan := range inputChans {
		go func() {
			for num := range inputChan {
				in <- num
			}

			wg.Done()
		}()

	}

	go func() {
		wg.Wait()
		close(in)
	}()

	return in
}

func PiplelineFanOutFanIn() {
	nums := []int{}

	for i := 1; i <= 100; i++ {
		nums = append(nums, i)
	}

	inputChan := generatePipleline(nums)

	c1 := fanOut(inputChan, "c1")
	c2 := fanOut(inputChan, "c2")
	c3 := fanOut(inputChan, "c3")

	sum := 0

	for n := range fanIn(c1, c2, c3) {
		sum += n
	}

	fmt.Println(sum)

}
