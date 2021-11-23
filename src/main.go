package main

import (
	"amolixs/duty"
	"amolixs/menu"
	"amolixs/utils"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
)

// Fonction qui permet d'afficher le logo
func printLogo() {
	utils.ClearConsole()
	fmt.Println("##########################")
	fmt.Println("Duty Manager")
	fmt.Println("##########################")
}

// Fonction qui permet de gérer le choix des options du menu
func handleMenu(db *sql.DB, choiceMenu int) {
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
		duty.DisplayDuty(currentDuty)
	case 4:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Entre l'id du devoir à supprimer : ")
		scanner.Scan()
		idDutyToDelete, _ := strconv.Atoi(scanner.Text())
		duty.DeleteDutyPerIdInTheDatabase(db, idDutyToDelete)
	}
}

// Fonction qui permet de détecter les erreurs
func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Fonction principal du programme
func main() {
	db, err := sql.Open("sqlite3", "db.db")
	checkErr(err)
	defer db.Close()
	printLogo()
	menu.PrintMenu()
	choiceMenu := menu.GetChoiceOption()
	utils.ClearConsole()
	handleMenu(db, choiceMenu)
}
