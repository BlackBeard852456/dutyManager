// Paquet qui gère les devoirs
package duty

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
)

type Duty struct {
	id       int
	name     string
	entilted string
	matter   string
}

// Fonction qui permet de créer un nouveau devoir
func CreateNewDuty() Duty {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entre le nom du devoir : ")
	dutyName, _ := reader.ReadString('\n')
	fmt.Print("Entre l'intitulé du devoir : ")
	dutyEntilted, _ := reader.ReadString('\n')
	fmt.Print("Entre la matière du devoir : ")
	dutyMatter, _ := reader.ReadString('\n')
	return Duty{name: dutyName, entilted: dutyEntilted, matter: dutyMatter}
}

// Fonction qui permet d'ajouter le devoir dans la bdd
func AddDutyInTheDatabase(databaseConnection *sql.DB, newDuty Duty) {
	stmt, _ := databaseConnection.Prepare("INSERT INTO duty (id, name, entilted, matter) VALUES (?, ?, ?, ?)")
	stmt.Exec(nil, newDuty.name, newDuty.entilted, newDuty.matter)
	defer stmt.Close()
}
