package main

import "fmt"

func main() {
	service, err := InitializMyService()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(service.Greeter.Message)
}
