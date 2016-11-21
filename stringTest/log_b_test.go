package stringTest

import (
	"testing"
)

func BenchmarkLogByStitchString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogByStitchString()
	}
}

func BenchmarkLogByMultiString2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogByMultiString()
	}
}

func BenchmarkLogByFormatString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogByFormatString()
	}
}
