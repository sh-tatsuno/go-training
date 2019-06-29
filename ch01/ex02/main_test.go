package main

import "testing"

func TestRun(t *testing.T) {
	cases := []struct {
		args     []string
		expected string
	}{
		{args: []string{"hoge", "go", "run", "main.go"}, expected: "0 : hoge\n1 : go\n2 : run\n3 : main.go"},
		{args: []string{}, expected: ""},
	}

	for _, c := range cases {
		actual := run(c.args)
		if actual != c.expected {
			t.Fatalf("mismatch result. actual: %v, expected: %v", actual, c.expected)
		}
	}
}
