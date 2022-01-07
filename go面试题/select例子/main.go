package main

import (
	"log"
	"time"
)

func main() {
	stopChan := make(chan struct{})
	ticker := time.NewTicker(100 * time.Millisecond)

	go func() {
		time.Sleep(2 * time.Second)
		close(stopChan)
	}()

	for {
		//加以下代码
		select {
		case <-stopChan:
		default:

		}

		time.Sleep(time.Second)
		log.Print("working hard")

		select {
		case <-ticker.C:
			log.Print("next")
		case <-stopChan:
			log.Print("Exit")
			return
		}
	}
}
