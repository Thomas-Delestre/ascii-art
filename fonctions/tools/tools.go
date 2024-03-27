package tools

import "strings"

// Convertit l'hexadecimal en décimal
func HexaToDecim(arg string) string {
	hx := strings.ToUpper(RevStr(arg))
	var res int
	for i := range hx {
		var tmp int = PowOf16(i)
		switch {
		case hx[i] >= 'A' && hx[i] <= 'F':
			res += tmp * int(byte(hx[i])-'A'+10)
		default:
			res += tmp * int(byte(hx[i])-'0')
		}
	}
	decim := Itoa(res)
	return decim
}

// retourne une puissance de 16
func PowOf16(pow int) int {
	var res int = 1
	for pow >= 1 {
		res *= 16
		pow--
	}
	return res
}

// Itoa
func Itoa(nb int) string {
	if nb == 0 {
		return "0"
	}
	var str string
	for nb > 0 {
		str = string(byte(nb%10)+'0') + str
		nb /= 10
	}
	return str
}

// Reverse string
func RevStr(arg string) string {
	var rev string
	for i := len(arg) - 1; i >= 0; i-- {
		rev += string(arg[i])
	}
	return rev
}
