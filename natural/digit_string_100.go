package natural

// DigitString100 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Hunderter-Stelle einer Zahl >= 21 vorkommen würde.
//
// Beispiel: Für 3 soll der String "dreihundert" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
func DigitString100(digit int) string {
	name := []string{"", "ein", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}
	for i, ef := range name {
		if i == digit {
			if i > 0 {
				ef += "hundert"
			}
			return ef

		}
	}
	return ""
}
