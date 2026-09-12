package concurrency

import (
	"fmt"
	"time"
)

type ListMessages struct {
	Name    string
	Message string
}

func RunChannel() {
	sendMessage()
}

func boardMessage(mes chan ListMessages, done chan bool, id int) {

	defer func() {
		done <- true
	}()
	fmt.Println("menunggu pesan...")
	fmt.Printf("worker %d\n", id)
	for m := range mes {
		fmt.Printf("name: %s\n", m.Name)
		fmt.Printf("message: %s\n", m.Message)
		time.Sleep(500 * time.Millisecond)
	}
}
func sendMessage() {
	message := []ListMessages{
		{
			Name:    "mulyono",
			Message: "saya akan lawan",
		},
		{
			Name:    "praboro",
			Message: "mbg bagus atau tidak?",
		},
		{
			Name:    "abah",
			Message: "angin tidak punya ktp",
		},
	}
	mboard := make(chan ListMessages)
	done := make(chan bool)

	go boardMessage(mboard, done, 1)
	go boardMessage(mboard, done, 2)
	go boardMessage(mboard, done, 3)
	// time.Sleep(1 * time.Second)
	fmt.Println("mengirim pesan...")
	for _, m := range message {
		mboard <- m
	}
	close(mboard)
	<-done
	fmt.Println("selesai")
}
