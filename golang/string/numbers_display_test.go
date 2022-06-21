package utils_string

import (
	"testing"
)

func TestNumberDisplay(t *testing.T) {
	decimal := DisplayDecimal(10)
	number_binary := DisplayBinary(10)
	octal := DisplayOctal(10)
	hexadecimal := DisplayHexadecimal(10)

	t.Log("Decimal: ", decimal)
	t.Log("Binary: ", number_binary)
	t.Log("Octal: ", octal)
	t.Log("Hexadecimal: ", hexadecimal)
}
