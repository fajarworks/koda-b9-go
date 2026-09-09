package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func GoToWork() {
	var wg sync.WaitGroup
	fmt.Println("start")
	wg.Add(4)
	go takeABath(&wg)
	go breakFast(&wg)
	go makeACoffe(&wg)
	go cleanUpRoom(&wg)
	wg.Wait()
	fmt.Println("berangkat")

}

func breakFast(wg *sync.WaitGroup) {
	defer fmt.Println("selesai sarapan")
	defer wg.Done()
	fmt.Println("Sarapan Duls")
	time.Sleep(1 * time.Second)
}
func makeACoffe(wg *sync.WaitGroup) {
	defer fmt.Println("Selesai Ngopi")
	defer wg.Done()
	fmt.Println("Ngopi Dulu")
	time.Sleep(2 * time.Second)
}
func takeABath(wg *sync.WaitGroup) {
	defer fmt.Println("Selesai Mandi")
	defer wg.Done()
	fmt.Println("Mandi Dulu")
	time.Sleep(3 * time.Second)
}
func cleanUpRoom(wg *sync.WaitGroup) {
	defer fmt.Println("Selesai beresin kamar")
	defer wg.Done()
	fmt.Println("beres-beres kamar dulu")
	time.Sleep(4 * time.Second)
}
