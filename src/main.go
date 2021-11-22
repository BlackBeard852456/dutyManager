package main

import (
	"amolixs/duty"
	"amolixs/menu"
	"database/sql"
	"fmt"
	"log"

	"github.com/dixonwille/wmenu/v5"
	_ "github.com/mattn/go-sqlite3"
)

// Fonction qui permet d'afficher le logo
func printLogo() {
	fmt.Println("##########################")
	fmt.Println("Duty Manager")
	fmt.Println("##########################")
}

// Fonction qui permet de gérer le choix des options du menu
func handleMenu(db *sql.DB, opts []wmenu.Opt) {
	switch opts[0].Value {
	case 0:
		newDuty := duty.CreateNewDuty()
		duty.AddDutyInTheDatabase(db, newDuty)
		break
	case 1:
		fmt.Println("Recherche d'un devoir")
	case 2:
		fmt.Println("Mise à jour d'un devoir")
	case 3:
		fmt.Println("Supression d'un devoir")
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
	fmt.Println(choiceMenu)
}
