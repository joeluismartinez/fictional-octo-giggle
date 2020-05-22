package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected %q but got %q.\nLoop executed %d times.", expected, repeated, len(repeated))
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
func ExampleRepeat() {
	repeat := Repeat("a")
	fmt.Println(repeat)
	// Output: aaaaa
}
