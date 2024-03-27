package color

import (
	"fmt"
	"strings"

	tools "ascii_art/fonctions/tools"
)

// Convertit le hsl(x,x,x) en rgb(x,x,x)
func HSLToRGB(str string) string {
	R, G, B := RGBToInt(str)
	if !(R >= 0 && R <= 100 && G >= 0 && G <= 100 && B >= 0 && B <= 100) {
		return "rgb(255,255,255)"
	}
	return "rgb(" + tools.Itoa(R*255/100) + "," + tools.Itoa(G*255/100) + "," + tools.Itoa(B*255/100) + ")"
}

// Convertit l'hexadécimal #XXXXXX en rgb(x,x,x)
func HexaToRGB(color string) string {
	for _, k := range color {
		if k != '#' && !(k >= '0' && k <= '9') && !(k >= 'a' && k <= 'f') {
			fmt.Println(string(k))
			return "rgb(255,255,255)"
		}
	}
	res := "rgb(" + tools.HexaToDecim(color[1:3]) + "," + tools.HexaToDecim(color[3:5]) + "," + tools.HexaToDecim(color[5:]) + ")"
	return res
}

// Convertit la "couleur" en rgb(x,x,x)
func ColorToRGB(color string) string {
	if strings.HasPrefix(color, "rgb(") || strings.HasPrefix(color, "ansi(") || strings.HasPrefix(color, "#") {
		return color
	}
	switch color {
	case "red":
		return "rgb(255,0,0)"
	case "orange":
		return "rgb(255,165,0)"
	case "yellow":
		return "rgb(255,255,0)"
	case "green":
		return "rgb(0,128,0)"
	case "blue":
		return "rgb(0,0,255)"
	case "indigo":
		return "rgb(75,0,130)"
	case "purple":
		return "rgb(148,0,211)"
	default:
		return "rgb(255,255,255)"
	}
}

// Récupère les valeurs en int d'un rgb(x,x,x)
func RGBToInt(str string) (int, int, int) {
	str = strings.ReplaceAll(str, "%", "")
	var slice []int
	var nb int
	for _, k := range str[4:] {
		if k >= '0' && k <= '9' {
			nb *= 10
			nb += int(byte(k) - '0')
		} else {
			slice = append(slice, nb)
			nb = 0
		}
	}
	if len(slice) != 3 {
		return 255, 255, 255
	}
	return slice[0], slice[1], slice[2]
}

func ANSIDecoder(color string) string {
	if len(color) > 10 {
		x, y, z := RGBToInt(color[1:])
		return tools.Itoa(x) + ";" + tools.Itoa(y) + ";" + tools.Itoa(z)
	}
	var nb int
	for _, k := range color {
		if k >= '0' && k <= '9' {
			nb *= 10
			nb += int(byte(k) - '0')
		}
	}

	return tools.Itoa(nb)
}

// Formate la string pour y inclure la couleur souhaitée
func Colorize(str, couleur string) string {
	if strings.HasPrefix(couleur, "rgb") {
		var R, G, B int = RGBToInt(couleur)
		return "\x1b[38;5;" + tools.Itoa(16+((R/51)*36)+((G/51)*6)+(B/51)) + "m" + str + "\x1b[0m"
	} else {
		couleur = ANSIDecoder(couleur)
		return "\x1b[" + couleur + "m" + str + "\x1b[0m"
	}
}
