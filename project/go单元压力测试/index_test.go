package index1

import "testing"

func BenchmarkMethod1(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Method1(10)
	}
}
