package tools

import (
	"log"
	"os"
	"strings"
)

// Transforme la bibliotheque standard et cherck la police entrée
func Biblio(input string) []string {
	// Définition du nom du fichier comprenant la police ascii
	police := "template/standard.txt"
	if input != "standard" {
		switch input {
		case "shadow":
			police = "template/shadow.txt"
		case "thinkertoy":
			police = "template/thinkertoy.txt"
		case "varsity":
			police = "template/varsity.txt"
		default:
			log.Fatalln("Police Introuvable")
		}
	}

	// Ouverture du fichier comprenant la police ascii
	in, _ := os.ReadFile(police)
	var str string
	for _, k := range in {
		if k != 13 {
			str += string(k)
		}
	}

	// Création d'un tableau de strings correspondant à la police de chaque caractères par ligne
	bib := strings.Split(str, "\n")
	return bib
}

// Verifie si l'argument est un nom de police
func IsAPolice(str string) bool {
	switch str {
	case "standard":
		return true
	case "shadow":
		return true
	case "thinkertoy":
		return true
	case "varsity":
		return true
	default:
		return false
	}
}
