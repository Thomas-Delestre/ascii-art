package reverse

import (
	"fmt"
	"os"
	"strings"

	color "ascii_art/fonctions/options/color"
	tools "ascii_art/fonctions/tools"
)

// Fonction qui return la police de l'ascii-art et la chaine de caratère qui le constitue
func Reverse(fileName, colorFlag string) (string, string) {
	dataTab := DataTabGetter(fileName) // Récupération du tableau
	var letters string                 // resultat du reverse
	var policeFound string             // police de l'ascii-art à déterminer

	letters += LetterDefiner(dataTab, "standard") // Récupération du resultat si standard
	policeFound = "standard"

	if letters == "" || HasOnlySpaces(letters) && 6*len(letters) != len(dataTab[0]) { // Récupération du resultat si thinkertoy
		letters += LetterDefiner(dataTab, "thinkertoy")
		policeFound = "thinkertoy"
	}
	if letters == "" || HasOnlySpaces(letters) && 6*len(letters) != len(dataTab[0]) { // Récupération du resultat si shadow
		letters += LetterDefiner(dataTab, "shadow")
		policeFound = "shadow"
	}
	if letters == "" || HasOnlySpaces(letters) && 6*len(letters) != len(dataTab[0]) { // Récupération du resultat si varsity
		letters += LetterDefiner(dataTab, "varsity")
		policeFound = "varsity"
	}

	// Gestion de la couleur de la réponse
	if colorFlag != "" {
		couleur := color.ColorFlagDecrypter(colorFlag)
		letters = color.Colorize(letters, couleur)
	}

	return policeFound, letters
}

// Ouverture du fichier contenant l'ascii-art à convertir puis return le tableau de string où chaque case représente une ligne de l'ascii-art à convertir
func DataTabGetter(fileName string) []string {
	// Ouverture du fichier
	dataToReverse, err := os.ReadFile("fonctions/options/reverse/filesToReverse/" + fileName)
	if err != nil {
		fmt.Println("Error while reading", fileName)
		os.Exit(1)
	}

	// Création de la string contenant tout l'ascii-art à convertir
	var result string
	for _, k := range dataToReverse {
		result += string(k)
	}

	// Création du tableau
	data := strings.Split(result, "\n")
	return data
}

// Fonction qui définit la chaine de caractères qui constitue l'ascii-art
func LetterDefiner(dataTab []string, police string) string {
	bib := tools.Biblio(police)                // Tableau de la police
	var letters string                         // Resultat de la convertion
	var remainingLetters int = len(dataTab[0]) // Longueur de l'ascii-art restant à convertir

	for remainingLetters > 0 { // Tant qu'il reste des lettres à décoder

		// String en ascii-art à tester pour chaque lettre, on ajoute un caractère ascii jusqu'à ce que la string corresponde puis est réinitialisée
		var asciiletter string

		for _, k := range dataTab[0][len(dataTab[0])-remainingLetters:] { // Range la première ligne du tableau sans les caractères traités

			var checker bool = false // Si ce booleen est true, c'est que j'ai trouvé la lettre

			asciiletter += string(k)

			if len(asciiletter) >= 4 { // Aucun caractère dessiné en ascii-art ne mesure moins de 4 caractères

				for asciiBibLine, asciiArtLine := range bib { // Parcours de la bibliothèque

					// La vérification ne se fera que s'il s'agit d'une première ligne de caractère ascii
					if asciiBibLine == 0 || (asciiBibLine-1)%9 == 0 {
						// Si la première ligne correspond, on verifie les suivantes
						if asciiArtLine == asciiletter {
							if IsLetter(bib[asciiBibLine:asciiBibLine+7], dataTab, remainingLetters, len(asciiletter)) {
								// Avec l'index de la première ligne du caractère ascii de la bibliotheque, on détermine la lettre
								letters += string(byte((asciiBibLine-1)/9 + 32))
								checker = true
							}
						}
					}
				}
				// La lettre est trouvée, on arrête enchaine avec les caratères suivants
				if checker {
					break
				}
			}
		}

		// On détermine les caractères restants à convertir
		remainingLetters -= len(asciiletter)
	}

	return letters
}

// Fonction qui détermine si les lignes suivantes correspondent à la lettre testée et dont la première ligne correspond
func IsLetter(bibLetter, dataTab []string, remainingLetters, asciiLen int) bool {
	for i := 0; i < 7; i++ {
		if bibLetter[i] != dataTab[i][len(dataTab[0])-remainingLetters:(len(dataTab[0])-remainingLetters)+asciiLen] {
			return false
		}
	}
	return true
}

// Fonction qui détermine si la string ne contient que des espaces
func HasOnlySpaces(str string) bool {
	for _, k := range str {
		if k != ' ' {
			return false
		}
	}
	return true
}
