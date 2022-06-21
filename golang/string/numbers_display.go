package utils_string

import "fmt"

func DisplayDecimal(number int) string {
	return fmt.Sprintf("%d", number)
}

func DisplayBinary(number int) string {
	return fmt.Sprintf("%b", number)
}

func DisplayOctal(number int) string {
	return fmt.Sprintf("%O", number)
}

func DisplayHexadecimal(number int) string {
	return fmt.Sprintf("%x", number)
}
