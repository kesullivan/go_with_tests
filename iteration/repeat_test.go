package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 6)
	expected := "aaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

// `go test -bench=.` to run benchmark
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

// go test -v to see this test run
func ExampleRepeat() {
	repeat := Repeat("b", 9)
	fmt.Println(repeat)
	// Output: bbbbbbbbb
}
