package natural

// DigitString1 erwartet eine Ziffer als int.
// Die Funktion gibt den zugehörigen String zurück, wie er üblicherweise
// an der Einer-Stelle einer Zahl >= 21 vorkommen würde.
// Außerdem wird bei Ziffern != 0 das Wort "und" angehängt.
//
// Beispiel: Für 3 soll der String "dreiund" geliefert werden.
//
// Anmerkung:
// Dies ist eine Hilfsfunktion, die genutzt werden soll,
// um den Gesamt-String einer Zahl zusammenzusetzen.
// Diese Funktion muss nur für den Normalfall (Zahlen >= 21) funktionieren.
func DigitString1(digit int) string {
	name := []string{"", "ein", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun"}
	for i, ef := range name {
		if i == digit {
			if i > 0 {
				ef += "und"
			}
			return ef

		}
	}
	return ""
}
