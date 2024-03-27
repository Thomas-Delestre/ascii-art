package color

import (
	"strings"
)

// Etablit la configuration de l'option color
func ColorFlagDecrypter(couleurFlag string) string {
	var couleur string
	var letter string

	// Définition des lettres précises à colorier et mise en forme de la couleur
	if couleurFlag != "" {
		for _, k := range couleurFlag {
			if k != ' ' {
				couleur += strings.ToLower(string(k))
				letter += string(k)
			}
		}
	}

	// Définition de la couleur depuis l'hexa ou
	if couleur[0] == '#' {
		couleur = HexaToRGB(couleur)
	} else if strings.HasPrefix(couleur, "hsl") {
		couleur = HSLToRGB(couleur)
	} else if !strings.HasPrefix(couleur, "rgb(") && !strings.HasPrefix(couleur, "hsl") {
		couleur = ColorToRGB(couleur)
	}

	return couleur
}
