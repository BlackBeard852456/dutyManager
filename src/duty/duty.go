// Paquet qui gère les devoirs
package duty

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
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
	dutyMatter := mattersManagement()
	return Duty{name: dutyName, entilted: dutyEntilted, matter: dutyMatter}
}

// Fonction qui gère le choix de la matière
func mattersManagement() string {
	scanner := bufio.NewScanner(os.Stdin)
	matters := [5]string{"Francais", "COMMU", "Bureautique", "Programmation", "Anglais"}
	for i, matter := range matters {
		fmt.Println(i+1, "-", matter)
	}
	fmt.Print("Choix matière : ")
	scanner.Scan()
	choiceMatter, _ := strconv.Atoi(scanner.Text())
	return matters[choiceMatter-1]
}

// Fonction qui permet d'ajouter le devoir dans la bdd
func AddDutyInTheDatabase(databaseConnection *sql.DB, newDuty Duty) {
	stmt, _ := databaseConnection.Prepare("INSERT INTO duty (id, name, entilted, matter) VALUES (?, ?, ?, ?)")
	stmt.Exec(nil, newDuty.name, newDuty.entilted, newDuty.matter)
	defer stmt.Close()
}

// Recherche un devoir dans la bdd
func SearchDutyInTheDatabase(databaseConnection *sql.DB, nameDuty string) []Duty {
	rows, err := databaseConnection.Query("SELECT * FROM duty WHERE name LIKE " + "'" + nameDuty + "%'")
	defer rows.Close()

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	dutys := make([]Duty, 0)

	for rows.Next() {
		ourDuty := Duty{}
		err = rows.Scan(&ourDuty.id, &ourDuty.name, &ourDuty.entilted, &ourDuty.matter)
		if err != nil {
			log.Fatal(err)
		}
		dutys = append(dutys, ourDuty)
	}
	return dutys
}
