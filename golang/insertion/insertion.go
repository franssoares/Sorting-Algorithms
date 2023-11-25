package insertion

/**
## Estratégia do algoritmo de acordo com seu nome

- Ordenação por inserção
- Escolhe um elemento do vetor desordenado, e insere no vetor ordenado na posição correta
  (que mantenha o vetor ordenado)
- Se no SelectionSort, a ideia é selecionar a menor carta da mão esquerda e trazer para o
  fim da mão direita, no InsertionSort
  a ideia é selecionar a próxima carta da mão esquerda e **inserir de modo ordenado na mão
  direita**.
*/

type InsertionSort struct {
	slice_values *[]int
}

func (ins *InsertionSort) SortAscending(slice []int) {

	ins.slice_values = &slice

	n := len((*ins.slice_values))

	for i := 1; i < n; i++ {

		for j := i; j > 0; j-- {

			if (*ins.slice_values)[j-1] > (*ins.slice_values)[j] {

				(*ins.slice_values)[j], (*ins.slice_values)[j-1] = (*ins.slice_values)[j-1], (*ins.slice_values)[j]

			}
		}
	}
}
