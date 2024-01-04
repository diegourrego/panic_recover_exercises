package exercise_2

import (
	"fmt"
	"os"
)

// Imprimir datos

// PrintData prints the data of the file
func PrintData() {
	file, err := os.Open("./exercise_2/customers.txt")

	defer func() {
		err := file.Close()
		if err != nil {
			panic("The file could not be closed")
		}
		fmt.Println("Execution finished")
	}()

	if err != nil {
		panic("The indicated file was not found or it is damaged")
	}

	fmt.Println(file)

}
