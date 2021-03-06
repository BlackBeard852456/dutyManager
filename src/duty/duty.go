// Package duty qui gère les devoirs
package duty

import (
	"amolixs/utils"
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

type Duty struct {
	id        int
	name      string
	entilted  string
	matter    string
	limitDate string
}

// CreateNewDuty Fonction qui permet de créer un nouveau devoir
func CreateNewDuty() Duty {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Entre le nom du devoir : ")
	scanner.Scan()
	dutyName := scanner.Text()
	fmt.Print("Entre l'intitulé du devoir : ")
	scanner.Scan()
	dutyEntilted := scanner.Text()
	dutyMatter := mattersManagement()
	fmt.Print("Entre la date limite de celui-ci : ")
	scanner.Scan()
	dutyLimitDate := scanner.Text()
	return Duty{name: dutyName, entilted: dutyEntilted, matter: dutyMatter, limitDate: dutyLimitDate}
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
	stmt, err := databaseConnection.Prepare("INSERT INTO duty (id, name, entilted, matter, limitDate) VALUES (?, ?, ?, ?, ?)")
	utils.CheckError(err)
	defer stmt.Close()
	utils.CreateProgressBar("Création du devoir")
	stmt.Exec(nil, newDuty.name, newDuty.entilted, newDuty.matter, newDuty.limitDate)
	color.Green("[+] Création du devoir terminé !")
	time.Sleep(1 * time.Second)
	utils.ClearConsole()
}

// Recherche un devoir dans la bdd
func SearchDutyInTheDatabase(databaseConnection *sql.DB, nameDuty string) []Duty {
	rows, err := databaseConnection.Query("SELECT * FROM duty WHERE name LIKE " + "'" + nameDuty + "%'")
	defer rows.Close()
	utils.CheckError(err)
	dutys := make([]Duty, 0)
	for rows.Next() {
		ourDuty := Duty{}
		err = rows.Scan(&ourDuty.id, &ourDuty.name, &ourDuty.entilted, &ourDuty.matter, &ourDuty.limitDate)
		utils.CheckError(err)
		dutys = append(dutys, ourDuty)
	}
	if len(dutys) == 0 {
		color.Red("[!] Aucun devoir trouvé !")
	} else {
		color.Green(fmt.Sprintf("[+] %d devoirs trouvés !", len(dutys)))
	}
	return dutys
}

// Permet de récupérer tous les devoirs
func GetAllDutys(databaseConnection *sql.DB) []Duty {
	rows, err := databaseConnection.Query("SELECT * FROM duty")
	defer rows.Close()
	utils.CheckError(err)
	dutys := make([]Duty, 0)
	for rows.Next() {
		ourDuty := Duty{}
		err = rows.Scan(&ourDuty.id, &ourDuty.name, &ourDuty.entilted, &ourDuty.matter, &ourDuty.limitDate)
		dutys = append(dutys, ourDuty)
	}
	return dutys
}

// Affiche des devoirs
func DisplayDutys(dutys []Duty) {
	for _, duty := range dutys {
		color.Yellow("#############################")
		color.Yellow(fmt.Sprintf("Id => %d", duty.id))
		color.Yellow(fmt.Sprintf("Nom => %s", duty.name))
		color.Yellow(fmt.Sprintf("Intitulé => %s", duty.entilted))
		color.Yellow(fmt.Sprintf("Matière => %s", duty.matter))
		color.Yellow(fmt.Sprintf("Date limite => %s", duty.limitDate))
		color.Yellow("#############################")
	}
}

// Ecrit tous les devoirs dans le ficher dutys.data
func WriteAllDutysInFile(nameFile string, allDutys []Duty) {
	if utils.VerifFileIsPresent(nameFile) {
		utils.ExecCommand("rm", []string{nameFile})
	}
	file, err := os.OpenFile(nameFile, os.O_WRONLY|os.O_CREATE, 0644)
	utils.CheckError(err)
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, duty := range allDutys {
		w.WriteString(fmt.Sprintf("ID => %d \n", duty.id))
		w.WriteString("NOM => " + duty.name + "\n")
		w.WriteString("INTITULE => " + duty.entilted + "\n")
		w.WriteString("MATIERE => " + duty.matter + "\n")
		w.WriteString("DATE LIMITE => " + duty.limitDate + "\n")
		w.WriteString("------------------------ \n")
	}
	w.Flush()
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
	time.Sleep(1 * time.Second)
	utils.ClearConsole()
	stmt.Close()
}

// Permet de récupérer un devoir grace a son identifiant
func GetDutyById(databaseConnection *sql.DB, idDuty string) Duty {
	rows, err := databaseConnection.Query("SELECT * FROM duty WHERE id = '" + idDuty + "'")
	utils.CheckError(err)
	defer rows.Close()
	dutyRetrieve := Duty{}
	for rows.Next() {
		rows.Scan(&dutyRetrieve.id, &dutyRetrieve.name, &dutyRetrieve.entilted, &dutyRetrieve.matter, &dutyRetrieve.limitDate)
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
	newMatter := mattersManagement()
	fmt.Printf("Date limite (actuel) : %s >", dutyToUpdate.limitDate)
	scanner.Scan()
	newLimitDate := scanner.Text()
	return Duty{id: dutyToUpdate.id, name: newName, entilted: newEntilted, matter: newMatter, limitDate: newLimitDate}
}

// Met à jour un devoir dans la bdd
func UpdateDutyInTheDatabase(databaseConnection *sql.DB, newDuty Duty) {
	stmt, err := databaseConnection.Prepare("UPDATE duty set name = ?, entilted = ?, matter = ?, limitDate = ? WHERE id = ?")
	utils.CheckError(err)
	defer stmt.Close()
	stmt.Exec(newDuty.name, newDuty.entilted, newDuty.matter, newDuty.limitDate, newDuty.id)
	color.Green("[+] Devoir mis à jour correctement !")
}
