package align

func AlignSpace(space int) string {
	var str string
	for i := 0; i < space; i++ {
		str += " "
	}
	return str
}
