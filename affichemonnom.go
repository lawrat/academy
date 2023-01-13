package affichage

import (
	"fmt"
	"os"
)

func main() {

	//ecrire un programme qui affiche mon nom

	e := os.Args[1]
	fmt.Println(e)
}

// output :
// lawratou
