package main

import (
	"fmt"
)

type Address struct {
	country    string
	city       string
	street     string
	postalCode int
}

// ** Exercise: define an 'Account' type with fields: username, password, age, address
type Account struct {
	username string
	password string
	age      int
	address  Address
}

func main() {
	// ** Exercise: declare and initialize an Account, name the variable "acc"
	acc := Account{
		username: "ftchen",
		password: "12345",
		age:      28,
		address: Address{
			country:    "Singapore",
			city:       "SG",
			street:     "Admiralty Link",
			postalCode: 750493,
		},
	}
	printAccount(acc)
}

func printAccount(account Account) {
	fmt.Println("--- printing account")
	fmt.Printf("username: %s\n", account.username)
	fmt.Printf("password: %s\n", account.password)
	fmt.Printf("age: %d\n", account.age)
	fmt.Printf("country: %s, city: %s, street: %s, postal code: %d\n", account.address.country, account.address.city, account.address.street, account.address.postalCode)
}
