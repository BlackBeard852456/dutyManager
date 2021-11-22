// Paquet qui gère le menu programme
package menu

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Affiche le menu
func PrintMenu() {
	fmt.Println("1-Ajouter un devoir")
	fmt.Println("2-Trouvez un devoir existant")
	fmt.Println("3-Mettre à jour un devoir existant")
	fmt.Println("4-Supprimer un devoir existant")
}

// Récupére le choix de l'optionA
func GetChoiceOption() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">> ")
	scanner.Scan()
	choiceInput, _ := strconv.Atoi(scanner.Text())
	return choiceInput
}
