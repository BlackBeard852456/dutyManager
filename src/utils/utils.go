// Package qui contient toutes les petites fonctions utiles au programme
package utils

import (
	"fmt"
	"os/exec"
)

func ClearConsole() {
	cmd := exec.Command("clear")
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}
