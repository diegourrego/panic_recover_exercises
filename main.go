package main

import (
	"fmt"
	"panic_recover_exercises/exercise_3"
)

func main() {
	defer func() {
		fmt.Println("End of execution")
		if r := recover(); r != nil {
			fmt.Println("Several errors were detected at runtime")
		}
	}()

	//exercise_1.OpenFile()
	//exercise_2.PrintData()

	exercise_3.RegisterClient(exercise_3.New(0, "./someWhere/myFile.txt", "Jennyfer S.",
		"3229876543", "Av cl fake 34 #12-34"))
}
