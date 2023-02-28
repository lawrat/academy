package main

import (
	"bufio"
	"fmt"
	"os"
)

func astLen(ast string) int {
	past := []rune(ast)
	a := 0
	for index := range past {
		a = index
	}
	return a + 1
}
func firstRune(s string) string {
	are := []rune(s)
	return string(are[0])
}
func lastRune(ptr string) string {
	res := []rune(ptr)
	return string(res[len(ptr)-1])
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("entrez votre prenom : ")
	scanner.Scan()

	ar := scanner.Text()

	for true {
		if ar == "" {
			fmt.Println("vous n'avez rien entrez !")
			fmt.Print("ressayer : ")
			scanner.Scan()
			ar = scanner.Text()
			continue
		} else {
			fmt.Println(astLen(ar))
			fmt.Println(firstRune(ar))
			fmt.Println(lastRune(ar))
			os.Exit(0)
		}
	}
}
