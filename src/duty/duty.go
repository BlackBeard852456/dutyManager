// Package duty qui gère les devoirs
package duty

import (
	"amolixs/utils"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/fatih/color"
)

type Duty struct {
	id       int
	name     string
	entilted string
	matter   string
}

// CreateNewDuty Fonction qui permet de créer un nouveau devoir
func CreateNewDuty() Duty {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Entre le nom du devoir : ")
	dutyName, _ := reader.ReadString('\n')
	fmt.Print("Entre l'intitulé du devoir : ")
	dutyEntilted, _ := reader.ReadString('\n')
	dutyMatter := mattersManagement()
	return Duty{name: dutyName, entilted: dutyEntilted, matter: dutyMatter}
}

// Function qui gère le choix de la matière
func mattersManagement() string {
	scanner := bufio.NewScanner(os.Stdin)
	matters := [5]string{"Francais", "COMMU", "Bureautique", "Programmation", "Anglais"}
	for i, matter := range matters {
		fmt.Println(i+1, "-", matter)
	}
	fmt.Print("Choix matière : ")
	scanner.Scan()
	choiceMatter, err := strconv.Atoi(scanner.Text())
	utils.CheckError(err)
	return matters[choiceMatter-1]
}

// AddDutyInTheDatabase Fonction qui permet d'ajouter le devoir dans la bdd
func AddDutyInTheDatabase(databaseConnection *sql.DB, newDuty Duty) {
	stmt, err := databaseConnection.Prepare("INSERT INTO duty (id, name, entilted, matter) VALUES (?, ?, ?, ?)")
	utils.CheckError(err)
	defer stmt.Close()
	utils.CreateProgressBar("Création du devoir")
	stmt.Exec(nil, newDuty.name, newDuty.entilted, newDuty.matter)
	color.Green("[+] Création du devoir terminé !")
}

// Recherche un devoir dans la bdd
func SearchDutyInTheDatabase(databaseConnection *sql.DB, nameDuty string) []Duty {
	rows, err := databaseConnection.Query("SELECT * FROM duty WHERE name LIKE " + "'" + nameDuty + "%'")
	defer rows.Close()
	utils.CheckError(err)
	dutys := make([]Duty, 0)
	for rows.Next() {
		ourDuty := Duty{}
		err = rows.Scan(&ourDuty.id, &ourDuty.name, &ourDuty.entilted, &ourDuty.matter)
		utils.CheckError(err)
		dutys = append(dutys, ourDuty)
	}
	if len(dutys) == 0 {
		color.Red("[!] Aucun devoir trouvé !")
	}
	return dutys
}

// Affiche des devoirs
func DisplayDutys(dutys []Duty) {
	for _, duty := range dutys {
		fmt.Println("#############################")
		fmt.Println("Id => ", duty.id)
		fmt.Println("Nom => ", duty.name)
		fmt.Println("Intitulé => ", duty.entilted)
		fmt.Println("Matière => ", duty.matter)
	}
}

// Affiche un devoir
func DisplayDuty(duty Duty) {
	fmt.Println("#############################")
	fmt.Println("Id => ", duty.id)
	fmt.Println("Nom => ", duty.name)
	fmt.Println("Intitulé => ", duty.entilted)
	fmt.Println("Matière => ", duty.matter)
}

// Supprime un devoir dans la bdd grace a son identifiant
func DeleteDutyPerIdInTheDatabase(databaseConnection *sql.DB, idDuty int) {
	stmt, err := databaseConnection.Prepare("DELETE FROM duty WHERE id = ?")
	utils.CheckError(err)
	_, err = stmt.Exec(idDuty)
	utils.CheckError(err)
	utils.CreateProgressBar("Suppression du devoir")
	color.Green("[+] Suppression du devoir terminé !")
	defer stmt.Close()
}

// Permet de récupérer un devoir grace a son identifiant
func GetDutyById(databaseConnection *sql.DB, idDuty string) Duty {
	rows, err := databaseConnection.Query("SELECT * FROM duty WHERE id = '" + idDuty + "'")
	utils.CheckError(err)
	defer rows.Close()
	dutyRetrieve := Duty{}
	for rows.Next() {
		rows.Scan(&dutyRetrieve.id, &dutyRetrieve.name, &dutyRetrieve.entilted, &dutyRetrieve.matter)
	}
	return dutyRetrieve
}

// Permet de mettre à jour un devoir existant
func UpdateDuty(dutyToUpdate Duty) Duty {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Nom (actuel) : %s >", dutyToUpdate.name)
	scanner.Scan()
	newName := scanner.Text()
	fmt.Printf("Intitulé (actuel) : %s >", dutyToUpdate.entilted)
	scanner.Scan()
	newEntilted := scanner.Text()
	fmt.Printf("Matière (actuel) : %s >", dutyToUpdate.matter)
	scanner.Scan()
	newMatter := scanner.Text()
	return Duty{id: dutyToUpdate.id, name: newName, entilted: newEntilted, matter: newMatter}
}

// Met à jour un devoir dans la bdd
func UpdateDutyInTheDatabase(databaseConnection *sql.DB, newDuty Duty) {
	stmt, err := databaseConnection.Prepare("UPDATE duty set name = ?, entilted = ?, matter = ? WHERE id = ?")
	utils.CheckError(err)
	defer stmt.Close()
	stmt.Exec(newDuty.name, newDuty.entilted, newDuty.matter, newDuty.id)
	color.Green("[+] Devoir mis à jour correctement !")
}
