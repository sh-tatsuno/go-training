package tmpconv

import (
	"testing"
)

func TestTmpconv(t *testing.T) {
	var actual, expected string

	c := AbsoluteZeroC
	actual = c.String()
	expected = "-273.15^C"
	if actual != expected {
		t.Fatalf("mismatch result. actual: %v, expected: %v", actual, expected)
	}

	f := CToF(c)
	actual = f.String()
	expected = "-459.66999999999996^F"
	if actual != expected {
		t.Fatalf("mismatch result. actual: %v, expected: %v", actual, expected)
	}
	fc := FToC(f)
	actual = fc.String()
	expected = "-273.15^C"
	if actual != expected {
		t.Fatalf("mismatch result. actual: %v, expected: %v", actual, expected)
	}

	k := CToK(c)
	actual = k.String()
	expected = "0^K"
	if actual != expected {
		t.Fatalf("mismatch result. actual: %v, expected: %v", actual, expected)
	}
	kc := KToC(k)
	actual = kc.String()
	expected = "-273.15^C"
	if actual != expected {
		t.Fatalf("mismatch result. actual: %v, expected: %v", actual, expected)
	}
}
