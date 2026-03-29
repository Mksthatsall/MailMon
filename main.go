package main

import (
	"fmt"

	"sync"
)

type Recipient struct {
	Name  string
	Email string
}

func main() {

	recipientChannel := make(chan Recipient)
	fmt.Println("hello and welcome to my golang email campaign project")

	go func() {
		loadRecipient("./names_and_emails.csv", recipientChannel)
	}()
	workerCount := 3

	var wg sync.WaitGroup

	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go emailWorker(i, recipientChannel, &wg)
	}

	wg.Wait()
}
