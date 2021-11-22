package main

import (
	"fmt"
	"log"

	"github.com/dixonwille/wmenu/v5"
)

// Fonction qui permet d'afficher le logo
func printLogo() {
	fmt.Println("##########################")
	fmt.Println("Duty Manager")
	fmt.Println("##########################")
}

// Fonction qui permet de gérer le choix des options du menu
func handleMenu(opts []wmenu.Opt) {
	switch opts[0].Value {
	case 0:
		fmt.Println("Ajout d'un nouveau devoir")
	case 1:
		fmt.Println("Recherche d'un devoir")
	case 2:
		fmt.Println("Mise à jour d'un devoir")
	case 3:
		fmt.Println("Supression d'un devoir")
	}
}

// Fonction principal du programme
func main() {
	printLogo()
	menu := wmenu.NewMenu(">> ")
	menu.Action(func(opts []wmenu.Opt) error { handleMenu(opts); return nil })
	menu.Option("Ajouter un devoir", 0, true, nil)
	menu.Option("Trouvez un devoir", 1, false, nil)
	menu.Option("Mettre à jour un devoir", 2, false, nil)
	menu.Option("Supprimer un devoir", 3, false, nil)
	err := menu.Run()
	if err != nil {
		log.Fatal(err)
	}
}
