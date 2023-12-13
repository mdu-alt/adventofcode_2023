package day03

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func findNumber(i int, line []byte) (int, int) {
	l, r := i, i
	for l >= 0 && isDigit(line[l]) {
		l--
	}
	for r < len(line) && isDigit(line[r]) {
		r++
	}
	return l + 1, r
}
