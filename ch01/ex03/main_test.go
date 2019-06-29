package main

import "testing"

func TestEchoSlow(t *testing.T) {
	cases := []struct {
		args     []string
		expected string
	}{
		{args: []string{"hoge", "go", "run", "main.go"}, expected: "hoge go run main.go"},
		{args: []string{}, expected: ""},
	}

	for _, c := range cases {
		actual := echoSlow(c.args)
		if actual != c.expected {
			t.Fatalf("mismatch result. actual: %v, expected: %v", actual, c.expected)
		}
	}
}

func TestEchoFast(t *testing.T) {
	cases := []struct {
		args     []string
		expected string
	}{
		{args: []string{"hoge", "go", "run", "main.go"}, expected: "hoge go run main.go"},
		{args: []string{}, expected: ""},
	}

	for _, c := range cases {
		actual := echoFast(c.args)
		if actual != c.expected {
			t.Fatalf("mismatch result. actual: %v, expected: %v", actual, c.expected)
		}
	}
}

var benchArgs = []string{"hoge", "go", "run", "main.go"}

func BenchmarkEchoSlow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoSlow(benchArgs)
	}
}

func BenchmarkEchoFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echoFast(benchArgs)
	}
}
