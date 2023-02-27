package main

import (
	"bufio"
	"fmt"
	"os"
)

func ptrLen(ptr string) int {
	astr := []rune(ptr)
	a := 1
	for index := range astr {
		a = index
	}
	return a + 1
}

func firstRune(str string) string {
	a := []rune(str)
	return string(a[0])
}

func lastRune(s string) string {
	c := []rune(s)
	return string(c[len(s)-1])
}

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("ENTREZ VOTRE PRENOM SVP : ")
	scanner.Scan()

	ar := scanner.Text()
	for true {
		if ar == "" {
			fmt.Println("VEUILLEZ VERIFIER VOTRE SAISIE, VOUS AVEZ ENTREZ UN ENSEMBLE VIDE")
			os.Exit(0)
		} else {
			fmt.Println(ptrLen(ar))
			fmt.Println(firstRune(ar))
			fmt.Println(lastRune(ar))
			return
		}
	}

}
