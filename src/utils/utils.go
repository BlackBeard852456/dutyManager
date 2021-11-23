// Package utils Package qui contient toutes les petites fonctions utiles au programme
package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/schollz/progressbar/v3"
)

// ClearConsole Permet de nettoyer la console
func ClearConsole() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

// CreateProgressBar Permet de créer une progress bar
func CreateProgressBar(description string) {
	bar := progressbar.NewOptions(100,
		progressbar.OptionSetDescription(description))
	for i := 0; i < 100; i++ {
		bar.Add(1)
		time.Sleep(10 * time.Millisecond)
	}
}

// CheckError Permet de checker une erreur
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
