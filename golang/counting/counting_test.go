package counting

import (
	"fmt"
	"reflect"
	"testing"
)

// O comando eh => go test -run TestSortAscending, go test -run bubble_test.go
// Ver mais conteudo -> go test -v
func TestSortAscending(t *testing.T) {

	var cou Counting

	// Imprime o método de ordenação
	fmt.Println("[Counting Sort]")

	// Caso de teste
	input := []int{-5, 2, 9, 1, 5, 6, 3, 8, 1, 0}
	expected := []int{-5, 0, 1, 1, 2, 3, 5, 6, 8, 9}

	// Imprime o vetor antes a ordenação
	fmt.Println("Vetor antes da ordenação:", input)

	// Chama o método de ordenação
	cou.SortAscending(input)

	// Verifica se a fatia foi ordenada corretamente
	if !reflect.DeepEqual(input, expected) {
		t.Errorf("Resultado incorreto. Esperado %v, obtido %v", expected, input)
	}

	// Imprime o vetor após a ordenação
	fmt.Println("Vetor após a ordenação:", input)
}
