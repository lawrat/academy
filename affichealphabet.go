package affichage

import "fmt"

func main() {

	//ecrire un programme qui affiche l'alphabet de a à z et de z à a

	for i := 'a'; i <= 'z'; i++ {
		fmt.Print(string(i))
	}
	fmt.Println()

	for j := 'z'; j >= 'a'; j-- {
		fmt.Print(string(j))
	}
	fmt.Println()
}

// output :
// abcdefghijklmnopqrstuvwxyz
// zyxwvutsrqponmlkjihgfedcba
