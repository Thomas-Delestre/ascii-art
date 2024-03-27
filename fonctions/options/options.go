package options

import (
	"flag"
	"fmt"
	"os"
	"strings"

	fonc "ascii_art/fonctions"
	rev "ascii_art/fonctions/options/reverse"
)

// Fonction qui détermine s'il y a une option dans la commande
func OptionChecker(arg []string) string {
	// Création des options et analyse
	help := flag.Bool("help", false, "Usage: go run . --help\nUse help option for more informations\n")
	reverse := flag.String("reverse", "", "Usage :\ngo run . [OPTION reverse] [OPTION align] [OPTION color] [OPTION color Option]\nEX: go run . --reverse=<fileName.txt> --color=red\n")
	output := flag.String("output", "", "Usage: go run . [OPTION output] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> \"texte\" template\n")
	color := flag.String("color", "", "Usage: go run . [OPTION color] [OPTION color Option] [STRING] [BANNER]\nEX: go run . --color=<colorChoice> <(optional)letter(s) you want to color> \"texte\" template\n")
	align := flag.String("align", "", "Usage: go run . [OPTION align] [STRING] [BANNER]\nEX: go run . --align=<alignChoice> \"texte\" template\nAlign option : \"justify\", \"left\", \"center\" and \"right\"\n")
	flag.Parse()

	// Option help
	if *help {
		fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . \"text\" \"template\"\n\n")
		fmt.Print("Options : \n--output : Save the result of ascii-art in the file <fileName.txt> without printing it in the terminal.\nCannot be used with others options.\n\n" + "--align : \nChoose the way the result is aligned in the terminal.\nchoices : \"justify\", \"left\", \"center\" and \"right\"\n\n" + "--color : \nChoose the color of the result in the terminal.\n\nColor encoding formats :\nRGB : go run . --color=\"rgb(x,y,z)\" \"text\"\nx,y and z are numbers from 0 to 255.\nx = red, y = green, z = blue\n\nHSL : go run . --color=\"hsl(x%,y%,z%)\" \"text\"\nx, y and z are numbers between 0 and 100.\nx = red, y = green, z = blue\n\nHexadecimal : go run . --color=\"#XXYYZZ\" \"text\"\nx,y and z are numbers between 0-9 and A-F.\nx = red, y = green, z = blue\n\nANSI Extended: go run . --color=\"ansi(38;5;XYZ)\" \"text\"\nXYZ is a number between 0 and 255.\nXYZ is the ID of the color in the extended ANSI color encoding format.\n\nANSI Standard: go run . --color=\"ansi(XY)\" \"text\"\nXY is the ID of the color in the standard ANSI color encoding format.\nANSI ID :\n \n\n")
		flag.PrintDefaults()
		os.Exit(0)
	}

	// Gestion Output
	if *output != "" && (*color != "" || *align != "" || *reverse != "") {
		fmt.Println("Error : output cannot be used with other options ! ")
		os.Exit(1)
	}

	// Gestion Reverse
	if *reverse != "" && *align != "" {
		fmt.Println("Error : reverse cannot be used with other options ! ")
		os.Exit(1)
	}

	// Gestion format align
	*align = strings.ToLower(*align)

	if *align != "" && *align != "center" && *align != "left" && *align != "right" && *align != "justify" {
		flag.Usage()
		os.Exit(1)
	}

	// Gestion format color
	if len(flag.Args()) > 2 && *color == "" {
		flag.Usage()
		os.Exit(1)
	}

	// Gestion format/nombre d'option
	var sameOptionColor int
	var sameOptionalign int
	var sameOptionOutput int
	var option bool
	var flarg []string = flag.Args()
	for i, k := range arg {
		if strings.HasPrefix(k, "--color=") {
			sameOptionColor++
		}
		if strings.HasPrefix(k, "--align=") {
			sameOptionalign++
		}
		if strings.HasPrefix(k, "--output=") {
			sameOptionOutput++
		}
		if strings.Contains(k, "--") && (!strings.HasPrefix(k, "--color=") && !strings.HasPrefix(k, "--align=") && !strings.HasPrefix(k, "--output=") || sameOptionColor > 1 || sameOptionalign > 1 || sameOptionOutput > 1) {
			flarg = arg[i:]
		}
		if sameOptionColor == 1 {
			flarg = arg[i+1:]
			break
		}
	}
	if len(os.Args[1:]) != len(flag.Args()) {
		for _, k := range os.Args[1:(len(os.Args) - len(flag.Args()))] {
			if !strings.HasPrefix(k, "--") {
				flag.Usage()
				os.Exit(1)
			}
		}
		option = true
	}

	// Generation du resultat
	var result string

	// Gestion de l'option reverse
	if *reverse != "" {
		var police string
		police, result = rev.Reverse(*reverse, *color)
		fmt.Println(police)
	} else {
		// Récupération du résultat d'ascii-art
		result = fonc.Conditions(flarg, *color, *align, option)
	}

	// Option output
	if *output != "" {
		file, _ := os.Create("fonctions/options/output/" + *output)
		file.WriteString(result)
		return "File " + *output + " saved"
	}
	return result
}
