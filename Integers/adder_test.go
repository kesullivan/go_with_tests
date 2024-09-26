package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Expected %d but got %d", expected, sum)
	}
}

// Functions prepended with "Example" in the test file create testable examples of our functions
// This is important to include so you have real examples of code usage that will break if things are changes
// the // Output: line is a special line that will be sure this function will run with go test
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
