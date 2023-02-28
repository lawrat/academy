package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("quelle est votre nom svp : ")
	scanner.Scan()
	ars := scanner.Text()
	if ars != "" {
		fmt.Print("pouvez vous lire et ecrire, repondez par O pour oui et N pour NON : ")
		scanner.Scan()
	}
	ar := scanner.Text()

	for true {
		if ar == "O" {
			fmt.Print("quelle est votre age : ")
			scanner.Scan()
			age, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("echec de conversion")
			}
			if age >= 16 {
				fmt.Printf("felicitations %s vous etes apte a vous inscrire !\n", ars)
				os.Exit(0)
			} else {
				fmt.Printf("desole %s vous n'avez pas l'age requis\n", ars)
				os.Exit(0)
			}
		} else if ar == "o" || ar == "n" {
			fmt.Println("entrez errone, veuillez choisir O pour oui et N pour non")
			fmt.Print("reesayez : ")
			scanner.Scan()
			ar = scanner.Text()
			continue
		} else if ar == "N" {
			fmt.Println("desole savoir lire et ecrire sont requis")
			os.Exit(0)
		} else {
			fmt.Println("repondez par O ou N svp !")
			fmt.Print("resayez : ")
			scanner.Scan()
			ar = scanner.Text()
			continue
		}
	}
}
