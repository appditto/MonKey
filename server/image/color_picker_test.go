package image

import (
	"testing"

	"github.com/appditto/monKey/server/color"
)

func TestGetColor(t *testing.T) {
	result, _ := GetColor("ff", "ee", "dd")
	expectedR := 0
	expectedG := 238
	expectedB := 33
	if int(result.R) != expectedR {
		t.Errorf("Expected %d got %f", expectedR, result.R)
	} else if int(result.G) != expectedG {
		t.Errorf("Expected %d got %f", expectedG, result.G)
	} else if int(result.B) != expectedB {
		t.Errorf("Expected %d got %f", expectedB, result.B)
	}
}

func TestShadowOpacityFur(t *testing.T) {
	clr := color.RGB{
		R: 0.0,
		G: 238.0,
		B: 33.0,
	}
	result := GetShadowOpacityFur(clr)
	expected := 0.13
	if result != expected {
		t.Errorf("Expected %f got %f", expected, result)
	}
}

func TestShadowOpacityFurDark(t *testing.T) {
	clr := color.RGB{
		R: 0.0,
		G: 238.0,
		B: 33.0,
	}
	result := GetShadowOpacityFurDark(clr)
	expected := 0.19
	if result != expected {
		t.Errorf("Expected %f got %f", expected, result)
	}
}

func TestShadowOpacityIris(t *testing.T) {
	clr := color.RGB{
		R: 0.0,
		G: 238.0,
		B: 33.0,
	}
	result := GetShadowOpacityIris(clr)
	expected := 0.1
	if result != expected {
		t.Errorf("Expected %f got %f", expected, result)
	}
}
