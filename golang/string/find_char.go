package utils_string

func FindChar(char string, str string) int {
	for index, letter := range str {
		if string(letter) == char {
			return index
		}
	}
	return -1
}
