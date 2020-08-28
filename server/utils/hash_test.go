package utils

import "testing"

func TestSha256(t *testing.T) {
	testInput := "7fc9064e4d713af2afc73c1527334b665972eb57d65093a378a3e40dbb48ec43"
	hashed := Sha256(testInput, "123456789")
	expected := "fad674ab79c5615a0eb6af3fe763ea892c3bbb589268a2791cbbef9a71a51039"
	if hashed != expected {
		t.Errorf("Expected %s got %s", expected, hashed)
	}
	// Try different seed
	hashed = Sha256(testInput, "987654321")
	expected = "1b51a15f5f37e1a3e5674f445ec0730436cd3292f1cd3a2752307c75d3bb6a1b"
	if hashed != expected {
		t.Errorf("Expected %s got %s", expected, hashed)
	}
}
