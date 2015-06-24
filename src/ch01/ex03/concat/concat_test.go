package concat_test

import (
	"ch01/ex03/concat"
	"strings"
	"testing"
)

func TestConcatWithOneElement(t *testing.T) {
	data := []string{"Hello"}

	result := concat.Concat(data)
	if result != "Hello" {
		t.Errorf("Result is '%s', want 'Hello'", result)
	}
}

func TestConcatWithTreeElements(t *testing.T) {
	data := []string{"Hello", "World", "!"}

	result := concat.Concat(data)
	if result != "Hello World !" {
		t.Errorf("Result is '%s', want 'Hello World !'", result)
	}
}

func BenchmarkConcat(b *testing.B) {
	data := strings.Split("Go is an open source programming langu age that makes it easy to bui ld simple, reliable, and efficient software", " ")

	for i := 0; i < b.N; i++ {
		concat.Concat(data)
	}
}

func BenchmarkJoin(b *testing.B) {
	data := strings.Split("Go is an open source programming langu age that makes it easy to bui ld simple, reliable, and efficient software", " ")

	for i := 0; i < b.N; i++ {
		strings.Join(data, " ")
	}
}
