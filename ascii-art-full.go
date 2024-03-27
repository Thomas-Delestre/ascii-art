package main

import (
	"fmt"
	"os"

	opt "ascii_art/fonctions/options"
)

func main() {
	arg := os.Args
	fmt.Println((opt.OptionChecker(arg)))
}
