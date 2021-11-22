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

func CreateNewDuty() {
	reader := bufio.NewReader(os.Stdin)
	dutyName, _ := reader.ReadString('\n')
	fmt.Println(dutyName)
}
