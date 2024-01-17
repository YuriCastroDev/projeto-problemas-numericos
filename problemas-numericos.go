package main

import "fmt"

func main() {
	// Desafio parte 1
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("Número %v é divisível por 3\n", i)
		}
	}

	// Desafio parte 2
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Pin")
			continue
		} else if i%3 != 0 && i%5 == 0 {
			fmt.Println("Pan")
			continue
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("Pin Pan")
			continue
		}
		fmt.Println(i)
	}
}
