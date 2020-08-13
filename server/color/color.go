package color

import (
	"errors"
	"fmt"
	"math"
)

// Red, green and blue multipliers to be used on perceived brightness calculations
const RedPBMultiplier = 0.241
const GreenPBMultiplier = 0.691
const BluePBMultiplier = 0.068

type RGB struct {
	R, G, B float64
}

// Calculates perceived brightness on a scale of 0,255
func (c RGB) PerceivedBrightness255() float64 {
	return math.Sqrt(RedPBMultiplier*c.R*c.R + GreenPBMultiplier*c.G*c.G + BluePBMultiplier*c.B*c.B)
}

// Calculates perceived brightness on a scale of 0,100
func (c RGB) PerceivedBrightness() float64 {
	return c.PerceivedBrightness255() / 255 * 100
}

// Takes a string like '#123456' or 'ABCDEF' and returns an RGB between 0..255
func HTMLToRGB(in string) (RGB, error) {
	if in[0] == '#' {
		in = in[1:]
	}

	if len(in) != 6 {
		return RGB{}, errors.New("Invalid string length")
	}

	var r, g, b byte
	if n, err := fmt.Sscanf(in, "%2x%2x%2x", &r, &g, &b); err != nil || n != 3 {
		return RGB{}, err
	}

	return RGB{float64(r), float64(g), float64(b)}, nil
}

// HTMLToRGB but returns nil instead of error if an error occurs
func HTMLToRGBAlt(in string) *RGB {
	rgb, err := HTMLToRGB(in)
	if err != nil {
		return nil
	}
	return &rgb
}

// A nudge to make truncation round to nearest number instead of flooring
const delta = 1 / 512.0

func (c RGB) ToHTML(withHash bool) string {
	if withHash {
		return fmt.Sprintf("#%02x%02x%02x", byte((c.R + delta)), byte((c.G + delta)), byte((c.B + delta)))
	}
	return fmt.Sprintf("%02x%02x%02x", byte((c.R + delta)), byte((c.G + delta)), byte((c.B + delta)))
}

// HSB/HSB

type HSB struct {
	H, S, B float64
}

// ToHSB - Returns HSB as (0..360, 0..1, 0..1)
func (c RGB) ToHSB() HSB {
	var h, s, v float64

	r := c.R
	g := c.G
	b := c.B

	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)
	delta := max - min

	if max != 0.0 {
		s = delta / max
		v = max / 255
	}

	// hue
	if delta != 0 {
		if r == max {
			h = (g - b) / delta
		} else if g == max {
			h = 2.0 + ((b - r) / delta)
		} else {
			h = 4 + ((r - g) / delta)
		}
	}
	h = h * 60
	if h < 0 {
		h = h + 360
	}

	return HSB{h, s, v}
}

// ToRGB - convert HSB to RGB
func (c HSB) ToRGB() RGB {
	Hp := c.H / 60.0
	C := c.B * c.S
	X := C * (1.0 - math.Abs(math.Mod(Hp, 2.0)-1.0))

	m := c.B - C
	r, g, b := 0.0, 0.0, 0.0

	switch {
	case 0.0 <= Hp && Hp < 1.0:
		r = C
		g = X
	case 1.0 <= Hp && Hp < 2.0:
		r = X
		g = C
	case 2.0 <= Hp && Hp < 3.0:
		g = C
		b = X
	case 3.0 <= Hp && Hp < 4.0:
		g = X
		b = C
	case 4.0 <= Hp && Hp < 5.0:
		r = X
		b = C
	case 5.0 <= Hp && Hp <= 6.0:
		r = C
		b = X
	}

	return RGB{255 * (m + r), 255 * (m + g), 255 * (m + b)}
}

func (c HSB) ToHTML(withHash bool) string {
	return c.ToRGB().ToHTML(withHash)
}

// HSL
type HSL struct {
	H, S, L float64
}

// ToHSL - Returns HSL as (0..360, 0..1, 0..1)
func (c RGB) ToHSL() HSL {
	var h, s, l float64

	r := c.R / 255
	g := c.G / 255
	b := c.B / 255

	min := math.Min(math.Min(r, g), b)
	max := math.Max(math.Max(r, g), b)
	delta := max - min

	// Luminosity is the average of the max and min rgb color intensities.
	l = (max + min) / 2

	// saturation
	if delta == 0 {
		// it's gray
		return HSL{0, 0, l}
	}

	// it's not gray
	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	// hue
	if delta != 0 {
		if r == max {
			h = (g - b) / delta
		} else if g == max {
			h = 2.0 + ((b - r) / delta)
		} else {
			h = 4 + ((r - g) / delta)
		}
	}
	h = h * 60
	if h < 0 {
		h = h + 360
	}

	return HSL{h, s, l}
}

// ToRGB - convert HSL to RGB
func (c HSL) ToRGB() RGB {
	h := c.H
	s := c.S
	l := c.L

	if s == 0 {
		// it's gray
		return RGB{l * 255.0, l * 255.0, l * 255.0}
	}

	v := (1.0 - math.Abs(2.0*l-1.0)) * s
	x := v * (1.0 - math.Abs(math.Mod(h/60.0, 2.0)-1.0))
	m := l - v/2.0

	var r, g, b float64

	if 0 <= h && h < 60 {
		r = v
		g = x
		b = 0
	} else if 60 <= h && h < 120 {
		r = x
		g = v
		b = 0
	} else if 120 <= h && h < 180 {
		r = 0
		g = v
		b = x
	} else if 180 <= h && h < 240 {
		r = 0
		g = x
		b = v
	} else if 240 <= h && h < 300 {
		r = x
		g = 0
		b = v
	} else if 300 <= h && h < 360 {
		r = v
		g = 0
		b = x
	}
	r = math.Round((r + m) * 255)
	g = math.Round((g + m) * 255)
	b = math.Round((b + m) * 255)

	return RGB{r, g, b}
}

func (c HSL) ToHTML(withHash bool) string {
	return c.ToRGB().ToHTML(withHash)
}
