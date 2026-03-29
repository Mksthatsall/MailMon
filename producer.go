package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func loadRecipient(filePath string, ch chan Recipient) error {
	defer close(ch)
	f, err := os.Open(filePath)

	if err != nil {
		fmt.Println("Error opening file:", err)
		return err
	}
	defer f.Close()
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		fmt.Println(record)
		ch <- Recipient{
			Name:  record[0],
			Email: record[1],
		}
	}
	return nil

}
