package dog

import "testing"

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
