package fonctions

import (
	"fmt"
	"os"
	"strings"

	align "ascii_art/fonctions/options/align"
	color "ascii_art/fonctions/options/color"
	tools "ascii_art/fonctions/tools"
)

// Fonction de conditions
func Conditions(flarg []string, couleur, align string, option bool) string {
	if len(flarg) < 1 || len(flarg) > 3 {
		return "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> \"texte\" template\nMore information : go run . --help\n"
	} else if len(flarg) == 1 && flarg[0] == "\\n" || len(flarg) == 2 && flarg[1] == "\\n" || len(flarg) == 3 && flarg[2] == "\\n" {
		return "\n"
	} else if len(flarg[0]) == 0 {
		return ""
	}

	// Définition de la police d'écriture
	police := "standard"
	if len(flarg) == 3 || len(flarg) == 2 && tools.IsAPolice(flarg[len(flarg)-1]) {
		police = flarg[len(flarg)-1]
	}

	return Ascii_Return(flarg, police, couleur, align, option)
}

// Fonction qui transforme l'argument en ascii art
func Ascii_Return(flarg []string, police string, couleurFlag, alignFlag string, option bool) string {
	var arg string = flarg[len(flarg)-1]

	if len(flarg) == 2 {
		if tools.IsAPolice(flarg[len(flarg)-1]) {
			arg = flarg[len(flarg)-2]
		} else {
			arg = flarg[len(flarg)-1]
		}
	} else if len(flarg) == 3 {
		arg = flarg[len(flarg)-2]
	}
	arr := strings.Split(arg, "\\n")
	var couleur string
	if couleurFlag != "" {
		couleur = color.ColorFlagDecrypter(couleurFlag)
	}
	var letters string

	if len(flarg) == 3 || len(flarg) == 2 && !tools.IsAPolice(flarg[len(flarg)-1]) {
		letters = flarg[0]
	}

	bib := tools.Biblio(police)

	var ascii string
	for i := 0; i < len(arr); i++ {
		if arr[i] != "" {
			ascii += Ascii_Write(arr[i], bib, couleur, letters, alignFlag)
		}
		if arr[i] == "" {
			ascii += "\n"
		}
	}
	return ascii
}

// Fonction qui écrit une ligne de caractère en ascii art en une string
func Ascii_Write(str string, bib []string, couleur, letters, alignFlag string) string {
	// Ascii-Art :
	var ascii string // String à return

	// Color :
	var colorized bool // Booleen qui determine si le caractère doit être en couleur
	var colorCount int // Nombre de caratères à colorier

	// Align
	var asciiStrLen int // Longueur d'une ligne des caractères ascii à print
	var spaceAll int    // Longueur des espaces à combler dans l'option align
	var soloSpace int   // Justify : Partie Principale de la longueur de chaque espace de --align="justify"

	// Etablissement des espaces de justify
	for _, k := range str {
		asciiStrLen += len(bib[(k-32)*9+1])
	}
	spaceAll = align.TerminalWidth() - asciiStrLen
	if alignFlag != "" {
		fmt.Println("Largeur du terminal :", align.TerminalWidth(), "Largeur de l'ASCII-ART à générer", asciiStrLen)
	}

	if align.TerminalWidth() < asciiStrLen {
		fmt.Println("Terminal trop petit :", align.TerminalWidth(), "<", asciiStrLen)
		os.Exit(0)
	}

	switch alignFlag {
	case "center":
		soloSpace = spaceAll / 2
	case "justify":
		soloSpace = spaceAll / (len(str) - 1)
	default:
		soloSpace = spaceAll
	}

	// Ecriture de la string à print
	for i := 0; i <= 7; i++ {
		var restingSpace int = spaceAll
		var nbDespacesResttants int = spaceAll - (soloSpace * (len(str) - 1))
		var count int

		for ii, k := range str {

			// Gestion du nombre de lettres à colorier
			if colorCount == 0 {
				colorized = false
			}

			// Gestion de la couleur si aucune lettre à colorier n'est précisée
			if couleur != "" && len(letters) == 0 {
				colorized = true
				colorCount = len(str)
			}

			// CHOIX 1 :
			// Gestion de couleur pour une chaine de caractère à colorier quand l'ensemble des lettres qui se suivent sont présentes
			if len(letters) > 0 && str[ii] == letters[0] && couleur != "" {
				colorized = true
				for iii := range letters {
					if letters[iii] != str[ii+iii] {
						colorized = false
						break
					} else {
						colorCount = len(letters)
					}
				}
			}

			// OU

			// CHOIX 2 :
			// Gestion de couleur pour une chaine de caractères à colorier individuellement
			// if len(letters) > 0 {
			// 	for iii := range letters {
			// 		if letters[iii] == str[ii] {
			// 			colorized = true
			// 			colorCount = 1
			// 		}
			// 	}
			// }

			// Gestion right/ center
			if (alignFlag == "right" || alignFlag == "center") && ii == 0 {
				ascii += align.AlignSpace(soloSpace)
				count += soloSpace
			}
			// Gestion justify tous sauf le dernier caractère
			if ii > 0 && ii <= len(str)-2 && alignFlag == "justify" {
				if nbDespacesResttants > 0 && restingSpace-soloSpace > 1 {
					ascii += align.AlignSpace(soloSpace + 1)
					restingSpace = restingSpace - (soloSpace + 1)
					count += soloSpace + 1
				} else {
					ascii += align.AlignSpace(soloSpace)
					restingSpace -= soloSpace
					count += soloSpace

				}
				nbDespacesResttants--
			}

			// Gestion justify dernier caractère
			if alignFlag == "justify" && ii == len(str)-1 {
				ascii += align.AlignSpace(restingSpace)
				count += restingSpace
			}
			// Implémentation des caractères ascii avec/sans couleur
			if colorized && colorCount > 0 {
				ascii += color.Colorize(bib[(k-' ')*9+1+rune(i)], couleur)
				colorCount--
			} else {
				ascii += bib[(k-32)*9+1+rune(i)]
			}
			// Gestion center / left
			if ii == len(str)-1 {
				if alignFlag == "left" {
					ascii += align.AlignSpace(soloSpace)
					count += soloSpace
				} else if alignFlag == "center" {
					ascii += align.AlignSpace(align.TerminalWidth() - (asciiStrLen + soloSpace))
					count += align.TerminalWidth() - (asciiStrLen + soloSpace)
				}
			}
			// Impression du reste ( doit être 0)
			if i == 0 && ii == len(str)-1 && alignFlag != "" {
				fmt.Println("Largeur restante du terminal", (align.TerminalWidth() - (asciiStrLen + count)))
			}
		}
		ascii += "\n"
	}
	return ascii
}
