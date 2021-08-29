package my

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	got := Add(1, 2)
	expected := 3
	if got != expected {
		t.Errorf("Did not get expected result. Got: %v, wanted: %v", got, expected)

	}
}

func BenchmarkFormat(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Format()
	}
}

func BenchmarkConcat(b *testing.B) {
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		Concat()
	}
}
