package main

import "fmt"

const USDtoEUR = 0.89
const USDtoRUB = 80.85

func main() {
	EURtoRUB := USDtoRUB / USDtoEUR
	fmt.Printf("1 EUR = %.2f RUB\n", EURtoRUB)
	input1, input2, input3 := inputUser()
	fmt.Println(input1)
	fmt.Println(input2)
	fmt.Println(input3)
	fmt.Println(calculation(100.0, 10.0, 5.0))
}

func inputUser() (string, string, string) {
	var input1, input2, input3 string
	fmt.Print("Введите значение 1: ")
	fmt.Scan(&input1)
	fmt.Print("Введите значение 2: ")
	fmt.Scan(&input2)
	fmt.Print("Введите значение 3: ")
	fmt.Scan(&input3)
	return input1, input2, input3
}

func calculation(num float64, initial float64, target float64) float64 {
	return num + initial + target
}
