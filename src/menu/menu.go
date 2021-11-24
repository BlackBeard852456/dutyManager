// Package menu Paquet qui gère le menu programme
package menu

import (
	"amolixs/duty"
	"amolixs/utils"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

// PrintLogo Fonction qui permet d'afficher le logo
func PrintLogo() {
	fmt.Println("##########################")
	fmt.Println("Duty Manager (by amolixs)")
	fmt.Println("##########################")
}

// PrintMenu Affiche le menu
func PrintMenu() {
	fmt.Println("1-Ajouter un devoir")
	fmt.Println("2-Trouvez un devoir existant")
	fmt.Println("3-Mettre à jour un devoir existant")
	fmt.Println("4-Supprimer un devoir existant")
	fmt.Println("5-Afficher tous les devoirs")
	fmt.Println("6-Quitter programme")
}

// GetChoiceOption Récupére le choix de l'optionA
func GetChoiceOption() int {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(">> ")
	scanner.Scan()
	choiceInput, _ := strconv.Atoi(scanner.Text())
	return choiceInput
}

// HandleMenu Fonction qui permet de gérer le choix des options du menu
func HandleMenu(db *sql.DB, choiceMenu int) {
	switch choiceMenu {
	case 1:
		newDuty := duty.CreateNewDuty()
		duty.AddDutyInTheDatabase(db, newDuty)
	case 2:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entrez des(un) devoir(s) à rechercher : ")
		scanner.Scan()
		searchInput := scanner.Text()
		var dutys []duty.Duty = duty.SearchDutyInTheDatabase(db, searchInput)
		duty.DisplayDutys(dutys)
	case 3:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entre l'id du devoir à mettre à jour : ")
		scanner.Scan()
		dutyIdToUpdate := scanner.Text()
		currentDuty := duty.GetDutyById(db, dutyIdToUpdate)
		newDuty := duty.UpdateDuty(currentDuty)
		duty.UpdateDutyInTheDatabase(db, newDuty)
	case 4:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entre l'id du devoir à supprimer : ")
		scanner.Scan()
		idDutyToDelete, err := strconv.Atoi(scanner.Text())
		utils.CheckError(err)
		duty.DeleteDutyPerIdInTheDatabase(db, idDutyToDelete)
		utils.ClearConsole()
	case 5:
		nameFile := "./dutys.data"
		allDutys := duty.GetAllDutys(db)
		duty.WriteAllDutysInFile(nameFile, allDutys)
		utils.ExecCommand("less", []string{nameFile})
	case 6:
		fmt.Println("Merci d'avoir utiliser DutyManager !")
		os.Exit(3)
	}
}
