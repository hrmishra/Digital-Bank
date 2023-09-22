package main

import (
	"fmt"

	"digitalbank/handler"
	"digitalbank/service"
)

func main() {
	fmt.Println("Enter Commands...")
	var inputFile string
	_, err := fmt.Scan(&inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	b := service.New()

	h := handler.NewHandler(b)
	h.BankOperations(inputFile)
}
