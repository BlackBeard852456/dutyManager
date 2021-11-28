// Paquet principal
package main

import (
	"amolixs/menu"
	"amolixs/utils"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

// Boucle principal du programme
func mainLoop(db *sql.DB) {
	for {
		menu.PrintLogo()
		menu.PrintMenu()
		choiceMenu := menu.GetChoiceOption()
		menu.HandleMenu(db, choiceMenu)
	}
}

// Function principal du programme
func main() {
	pathDatabase := fmt.Sprintf("%s/.dutyManager/db.db", os.Getenv("HOME"))
	db, err := sql.Open("sqlite3", pathDatabase)
	utils.CheckError(err)
	mainLoop(db)
	defer db.Close()
}
