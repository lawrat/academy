package affichage

import (
	"fmt"
	"os"
	"strconv"
)

func categorie() {
	// ecrire un programme qui demande l'age et ensuite ramène la catégorie de la personne
	age, _ := strconv.Atoi(os.Args[1])

	switch {
	case age < 19:
		fmt.Println("categorie : MINEUR")
	case age < 50:
		fmt.Println("categorie : MAJEUR")
	default:
		fmt.Println("categorie : SENIOR")
	}
}

// output :
// categorie : MINEUR
