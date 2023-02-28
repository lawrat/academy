package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("entrez votre prenom : ")
	scanner.Scan()
	name := scanner.Text()

	fmt.Println(name, "\n", len(name), "\n", string(name[0]), "\n", string(name[len(name)-1]))
}
