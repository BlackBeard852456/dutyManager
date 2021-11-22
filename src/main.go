package main

import (
	"amolixs/duty"
	"amolixs/menu"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

// Fonction qui permet d'afficher le logo
func printLogo() {
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
		var dutys []duty.Duty = duty.SearchDutyInTheDatabase(db, "b")
		duty.DisplayDutys(dutys)
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
	handleMenu(db, choiceMenu)
}
