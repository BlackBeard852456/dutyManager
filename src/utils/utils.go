// Package qui contient toutes les petites fonctions utiles au programme
package utils

import (
	"fmt"
	"os"
	"os/exec"
)

func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
