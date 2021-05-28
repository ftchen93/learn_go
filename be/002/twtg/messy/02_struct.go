package main

import (
	"fmt"
)

//A struct is a collection of fields
// struct fields are accessed using a dot

// defining our custom 'Address' data type
type Address struct {
	country    string
	city       string
	street     string
	postalCode int
}

func main() {
	varDeclare()
	// colonDeclare()
}

func varDeclare() {
	// declaring an address
	fmt.Println("--- create with var")
	var shopeeBuilding Address

	// setting fields of an address
	shopeeBuilding.country = "Singapore"
	shopeeBuilding.city = "Singapore"
	shopeeBuilding.street = "5 Science Park Dr"
	shopeeBuilding.postalCode = 118265

	// reading fields of an address
	fmt.Printf("Country: %s\n", shopeeBuilding.country)
	fmt.Printf("City: %s\n", shopeeBuilding.city)
	fmt.Printf("Street: %s\n", shopeeBuilding.street)
	fmt.Printf("Postal Code: %d\n", shopeeBuilding.postalCode)
	fmt.Println()
}

func colonDeclare() {
	fmt.Println("--- create with :=")
	orchardGateway := Address{ // ** we don't have to set all fields of a struct **
		country: "Singapore",
		city:    "Singapore",
	}

	fmt.Println(orchardGateway.country)
	fmt.Println(orchardGateway.city)
	fmt.Println(orchardGateway.street)
	fmt.Println(orchardGateway.postalCode)
}
