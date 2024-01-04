package exercise_3

import (
	"errors"
	"fmt"
)

type Customer struct {
	ID          int
	File        string
	Name        string
	PhoneNumber string
	Home        string
}

func New(id int, file, name, phoneNumber, home string) Customer {
	return Customer{
		ID:          id,
		File:        file,
		Name:        name,
		PhoneNumber: phoneNumber,
		Home:        home,
	}
}

// Methods

// ExistingClients simulates a database of customers
func existingClients() []Customer {
	return []Customer{
		New(1, "./exercise_3/customers.txt", "John Doe", "123456789", "Cl fake 123"),
		New(2, "./exercise_3/customers.txt", "Jane Doe", "987654321", "Cl fake 456"),
		New(3, "./exercise_3/customers.txt", "Mary Doe", "123456789", "Cl fake 897"),
	}
}

func validateNewCustomerData(customer Customer) (err error) {
	switch {
	case customer.ID == 0:
		err = errors.New("0 is a invalid value")
	case customer.Name == "", customer.File == "", customer.Home == "", customer.PhoneNumber == "":
		err = errors.New(" '' is a invalid value")
	}
	return
}

func validateIfCustomerExists(clientsList []Customer, customer Customer) {
	for _, client := range clientsList {
		if customer.PhoneNumber == client.PhoneNumber || customer.ID == client.ID {
			fmt.Println("Error: client already exists")
			panic("")
		}
	}
}

func RegisterClient(customer Customer) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Several errors were detected at runtime")
		}
	}()

	clientsList := existingClients()
	fmt.Println("Clients list before", clientsList)

	validateIfCustomerExists(clientsList, customer)

	// Check values in new Customer:
	err := validateNewCustomerData(customer)
	if err != nil {
		panic("Invalid data in some field of customer!")
	}

	// Add new customer
	clientsList = append(clientsList, customer)
	fmt.Println("Clients list after", clientsList)

}
