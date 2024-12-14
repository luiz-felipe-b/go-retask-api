package iteration

import (
	"fmt"
    "strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 7)
	expected := "aaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func TestClone(t *testing.T) {
    repeated := Repeat("a", 7)
    cloned := strings.Clone(repeated)
    if repeated != cloned {
        t.Errorf("expected %q but got %q", repeated, cloned)
    }
}

func ExampleRepeat() {
    repeated := Repeat("a", 7)
    fmt.Println(repeated)
    // Output: aaaaaaa
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}
