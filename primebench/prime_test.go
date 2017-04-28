package primebench

import "testing"

func BenchmarkPrimeInGoStyle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve()
	}
}

func TestPrimeInGoStyle(t *testing.T) {
	sieve()
}
