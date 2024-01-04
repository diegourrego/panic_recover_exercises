package exercise_1

import (
	"fmt"
	"os"
)

// Datos clientes

func OpenFile() {
	defer func() {
		fmt.Println("Execution finished")
	}()
	file, err := os.Open("customers.txt")
	if err != nil {
		panic("The indicated file was not found or it is damaged")
	}
	fmt.Println(file)
}
