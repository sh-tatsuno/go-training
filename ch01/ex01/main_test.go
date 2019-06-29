package main

import "testing"

func TestRun(t *testing.T) {
	cases := []struct {
		args     []string
		expected string
	}{
		{args: []string{"hoge", "go", "run", "main.go"}, expected: "hoge go run main.go"},
		{args: []string{}, expected: ""},
	}

	for _, c := range cases {
		actual := run(c.args)
		if actual != c.expected {
			t.Fatalf("mismatch result. actual: %v, expected: %v", actual, c.expected)
		}
	}
}
