package parse_string

import (
	"math"
	"slices"
	"strings"
)

// ParseStringDecimal erwartet einen String, der eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringHexadecimal(s string, b int) int {
	var n int
	var m int
	value := strings.Split(s, "")
	slices.Reverse(value)
	for i, ef := range value {
		m = ParseDigit(ef)
		if m == -1 {
			return -1
		}
		n += int(math.Pow(float64(b), float64(i))) * ParseDigit(ef)
	}

	return n
}
