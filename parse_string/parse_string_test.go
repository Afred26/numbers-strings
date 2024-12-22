package parse_string

import "fmt"

func ExampleParseString() {

	fmt.Println(ParseString("0x12AB4"))
	fmt.Println(ParseString("="))
	fmt.Println(ParseString("6543"))
	//fmt.Println(ParseString("06543"))
	fmt.Println(ParseString("0b1100110"))

	//Output:
	//76468
	//-1
	//6543
	//102
}
