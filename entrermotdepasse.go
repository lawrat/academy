package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// ecrire un programme qui demande a l'utilisateur d'entree son mot de passe pour 3 essaie

	scanner := bufio.NewScanner(os.Stdin)

	code := 2023

	for i := 0; i < 3; i++ {
		fmt.Println("entrez votre mot de passe : ")
		scanner.Scan()
		p := scanner.Text()

		motp, _ := strconv.Atoi(p)

		if motp == code {
			fmt.Println("GOOD")
			break
		} else {
			fmt.Println("ERREUR MOT DE PASSE BLOQUE !!!")
		}
	}

}
