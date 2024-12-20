package parse_string

import "fmt"

func ExampleParseString() {

	fmt.Println(ParseString("0x 12AB4"))
	fmt.Println(ParseString("q"))
	fmt.Println(ParseString("0 6543"))

	//Output:
	//76468
	//-1
	//6543
}
