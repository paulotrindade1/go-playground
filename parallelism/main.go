package main

import (
	"fmt"
	"sync"
	"time"
)

func printTime(msg string) {
	fmt.Println(msg, time.Now().Format("15:04:25"))
}

func writeMail1(wg *sync.WaitGroup) {
	printTime("Done writing mail #1")
	wg.Done()
}

func writeMail2(wg *sync.WaitGroup) {
	printTime("Done writing mail #2")
	wg.Done()
}

func writeMail3(wg *sync.WaitGroup) {
	printTime("Done writing mail #3")
	wg.Done()
}

func printCars(cars []string) {
	fmt.Println(cars)
}

func listenForever() {
	for {
		printTime("Listening...")
	}
}

func main() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(3)
	var cars = []string{"🚗", "🚙", "🚕"}

	go listenForever()

	time.Sleep(time.Second * 1)

	writeMail1(&waitGroup)
	writeMail2(&waitGroup)
	printCars(cars)
	writeMail3(&waitGroup)

	waitGroup.Wait()
}
