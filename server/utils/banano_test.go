package utils

import (
	"testing"
)

func TestGenerateAddress(t *testing.T) {
	generated := GenerateAddress()
	if !ValidateAddress(generated) {
		t.Errorf("Invalid address %s", generated)
	}
}

func TestValidateAddress(t *testing.T) {
	// Valid
	valid := "ban_1zyb1s96twbtycqwgh1o6wsnpsksgdoohokikgjqjaz63pxnju457pz8tm3r"
	if !ValidateAddress(valid) {
		t.Errorf("Valid address test failed %s", valid)
	}
	// Invalid
	invalid := "ban_1zyb1s96twbtycqwgh1o6wsnpsksgdoohokikgjqjaz63pxnju457pz8tm3ra"
	if ValidateAddress(invalid) {
		t.Errorf("Valid address returned true when should have been false %s", invalid)
	}
	invalid = "ban_1zyb1s96twbtycqwgh1o6wsnpsksgdoohokikgjqjaz63pxnju457pz8tm3rb"
	if ValidateAddress(invalid) {
		t.Errorf("Valid address returned true when should have been false %s", invalid)
	}
}
