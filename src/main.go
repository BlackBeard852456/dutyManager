// Paquet principal
package main

import (
	"amolixs/menu"
	"amolixs/utils"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

// Function principal du programme
func main() {
	db, err := sql.Open("sqlite3", "db.db")
	utils.CheckError(err)
	defer db.Close()
	menu.PrintLogo()
	menu.PrintMenu()
	choiceMenu := menu.GetChoiceOption()
	utils.ClearConsole()
	menu.HandleMenu(db, choiceMenu)
}
