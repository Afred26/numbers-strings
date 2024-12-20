package parse_string

import "strings"

// ParseString erwartet einen String, der entweder eine Dezimal- oder eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseString(s string) int {
	bevor, after, _ := strings.Cut(s, " ")
	if bevor == "0x" {
		return ParseStringHexadecimal(after)
	}
	if bevor == "0" {
		return ParseStringDecimal(after)
	}

	return -1
}
