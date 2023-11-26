package counting

import (
	"testing"
)

// go test -bench . (com o teste)
// go test -bench . -run XXX (sem o teste)
func BenchmarkSortAscending(b *testing.B) {
	var cou Counting

	// Caso de teste
	input := []int{5, 2, 9, 1, 5, 6, 3, 8, -1, 0}

	// Execute a função de ordenação no loop do benchmark
	for i := 0; i < b.N; i++ {
		// Crie uma cópia da entrada para garantir que cada iteração use os mesmos dados iniciais
		inputCopy := make([]int, len(input))
		copy(inputCopy, input)

		cou.SortAscending(inputCopy)
	}
}
