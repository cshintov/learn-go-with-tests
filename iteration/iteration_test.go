package iteration

import "testing"
import "fmt"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func ExampleRepeat() {
    fmt.Println(Repeat("b", 1))
    // Output: b
}

func ExampleRepeat_times0() {
    fmt.Println(Repeat("b", 0))
    // Output: 
}

func ExampleRepeat_times5() {
    fmt.Println(Repeat("b", 5))
    // Output: bbbbb
}

func ExampleRepeat_times4() {
    fmt.Println(Repeat("b", 4))
    // Output: bbbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}
