package natural

// NumberString3Digits erwartet eine Zahl 0 <= n <= 999 und liefert den zugehörigen String.
func NumberString3Digits(n int) string {
	num := digit(n)
	//var name []string{}
	if n < 13 {
		eins_zwölf(n)
	}
	p := num[0] + num[1]
	if p < 13 {
		return eins_zwölf(p)
	}
	return ""
}

//i.O.
func digit(n int) []int {
	num := []int{}
	var d int
	for n != 0 {
		d = n % 10
		n /= 10
		num = append(num, d)
	}
	//slices.Reverse(num)
	return num

}

// i.O.
func eins_zwölf(n int) string {
	Zahl := []string{"null", "eins", "zwei", "drei", "vier", "fünf", "sechs", "sieben", "acht", "neun", "zehn", "elf", "zwölf"}
	for i, ef := range Zahl {
		if n == i {
			return ef
		}
	}
	return ""
}
