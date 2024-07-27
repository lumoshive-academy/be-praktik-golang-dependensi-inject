package main

import "fmt"

func main() {
	service, err := InitializMyService("Lumoshive")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(service.Greeter.Message)
}
