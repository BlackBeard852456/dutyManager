// Paquet qui gère les devoirs
package duty

import (
	"bufio"
	"fmt"
	"os"
)

type Duty struct {
	id       int
	name     string
	entilted string
	matter   string
}

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
