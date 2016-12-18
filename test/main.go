package main

import "fmt"

func main() {
	var public string
	var private string
	fmt.Println("Public API key:")
	fmt.Scanln(&public)
	fmt.Println("Private API key:")
	fmt.Scanln(&private)
}
