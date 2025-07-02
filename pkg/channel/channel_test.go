package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("data before dikirim")
		channel <- "Ini data yang dikirim"
		fmt.Println("data after dikirim")
	}()

	data := <-channel
	fmt.Println(data)
	close(channel)
}

func GiveMeData(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Data 123456"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	go GiveMeData(channel)
	data := <-channel
	fmt.Println(data)
	close(channel)
}
