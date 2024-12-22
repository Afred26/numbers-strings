package parse_string

import "testing"

func TestParseStringHexadecimal(t *testing.T) {
	if ParseStringHexadecimal("0", 16) != 0 {
		t.Errorf("ParseStringHexadecimal(\"0\") == %d, expected 0", ParseStringHexadecimal("0", 16))
	}

	if ParseStringHexadecimal("1", 16) != 1 {
		t.Errorf("ParseStringHexadecimal(\"1\") == %d, expected 1", ParseStringHexadecimal("1", 16))
	}

	if ParseStringHexadecimal("A", 16) != 10 {
		t.Errorf("ParseStringHexadecimal(\"A\") == %d, expected 10", ParseStringHexadecimal("A", 16))
	}

	if ParseStringHexadecimal("456", 16) != 1110 {
		t.Errorf("ParseStringHexadecimal(\"456\") == %d, expected 1110", ParseStringHexadecimal("456", 16))
	}

	if ParseStringHexadecimal("38", 16) != 56 {
		t.Errorf("ParseStringHexadecimal(\"38\") == %d, expected 56", ParseStringHexadecimal("38", 16))
	}
}
