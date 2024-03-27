package align

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func TerminalWidth() int {
	cmd := exec.Command("tput", "cols")
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Impossible de récupérer la largeur du terminal:", err)
		return 0
	}

	columns := strings.TrimSpace(string(output))
	width, _ := strconv.Atoi(columns)
	return width
}
