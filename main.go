package main

import "fmt"

type record struct {
	name           string
	quantity       int
	price          float32
	expirationDate int64
	dateOfRecord   int64
}

func main() {
	fmt.Println("Welcome to Smart Pharma")
}

func createARecord(name string, quantity int, price float32, expirationDate int64, dateOfRecord int64) record {
	return record{name, quantity, price, expirationDate, dateOfRecord}
}
