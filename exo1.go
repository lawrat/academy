package main

import (
	"fmt"
	"os"
)

func main() {

	arg := os.Args[1]
	alphabet := "ABCDEFJHIJKLMNOPQRSTUVWXYZ"
	alpha := "abcdefjhijklmnopqrstuvwxyz"

	for _, v := range arg {
		for i, val := range alphabet {
			if v == val {
				fmt.Printf("%s\n", alphabet[i:])
			}
		}
		for i, val1 := range alpha {
			if v == val1 {
				fmt.Printf("%s\n", alpha[i:])
			}
		}
	}
}
