package parse_string

import (
	"strconv"
)

// ParseString erwartet einen String, der entweder eine Dezimal- oder eine Hexadezimalzahl repräsentiert, und liefert die zugehörige Zahl.
// Ist der String kein gültiger Wert, wird -1 zurückgegeben.
func ParseString(s string) int64 {
	if n, err := strconv.ParseInt(s, 0, 0); err == nil {
		return n
	} else {
		return -1
	}
}
