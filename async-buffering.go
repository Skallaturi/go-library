package main

import "fmt"

func GetBufferedIntSlice() []int {
	channel := make(chan int, 100) //TODO: buffer size?
	response := []int{}

	go dummyAsyncOutput(channel)
	fmt.Println("waiting for packets")
	for packet := range channel {
		fmt.Printf("recieved %v\n", packet)
		response = append(response, packet)

	}
	fmt.Println("no more packets")
	return response
}

func dummyAsyncOutput(outputChannel chan int) {
	for i := 0; i < 50; i++ {
		// time.Sleep(500 * time.Millisecond)
		fmt.Printf("sent %v\n", i)
		outputChannel <- i
	}
	close((outputChannel))
}
