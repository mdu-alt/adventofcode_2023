package day02

func isNotAlphanumeric(r rune) bool {
	return !('0' <= r && r <= '9') && !('A' <= r && r <= 'Z') && !('a' <= r && r <= 'z')
}
