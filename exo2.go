package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("POUVEZ VOUS LIRE ET ECRIRE ? REPONDEZ PAR O pour oui et N pour non : ")
	scanner.Scan()
	ar := scanner.Text()

	for true {
		if ar == "O" || ar == "o" {
			scanner := bufio.NewScanner(os.Stdin)
			fmt.Print("QUELLE EST VOTRE AGE : ")
			scanner.Scan()
			age, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("LA CONVERSION A ECHOUER")
			}
			if age >= 16 {
				fmt.Println("FELICITATIONS VOUS ETES APTE A FAIRE LE TEST D'ENTREE !")
				os.Exit(0)
			} else {
				fmt.Println("DOMAGE VOUS ETES INAPTE !")
				os.Exit(0)
			}
		} else if ar == "N" || ar == "n" {
			fmt.Println("DESOLE SAVOIR LIRE ET ECRIRE SONT REQUIS !")
			os.Exit(0)
		} else {
			fmt.Println("REPONDEZ PAR O OU N SVP")
			fmt.Print("REESSAYEZ SVP : ")
			scanner.Scan()
			ar = scanner.Text()
			continue
		}
	}
}
