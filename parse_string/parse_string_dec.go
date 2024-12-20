package parse_string

import "strconv"

// ParseStringDecimal erwartet einen String, der eine Dezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseStringDecimal(s string) int {
	if p, err := strconv.Atoi(s); err == nil {
		return p
	}
	return -1
}
