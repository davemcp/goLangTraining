package dog

import (
	"fmt"
	"testing"
)

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(12)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(12)
	}
}

func ExampleYears() {
	fmt.Println(Years(12))
	// Output:
	// 84
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(12))
	// Output:
	// 84
}

func TestYears(t *testing.T) {
	x := Years(10)
	if x != 70 {
		t.Error("got", x, "expected", 70)
	}
}

func TestYearsTwo(t *testing.T) {
	x := YearsTwo(10)
	if x != 70 {
		t.Error("got", x, "expected", 70)
	}
}