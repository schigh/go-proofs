package main

import (
	"testing"
)

func BenchmarkRun(b *testing.B) {
	b.Run("jsoniterator_compat", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			run()
		}
	})
}
