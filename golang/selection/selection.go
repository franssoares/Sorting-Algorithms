package selection

/**
	=> min, max => O(n²)
  Estratégia do algoritmo de acordo com seu nome:
  -selecionar os menores valores e posicionar no final do vetor ordenado
    selecionamos o menor valor do vetor e trocamos de posição com o primeiro elemento
    selecionamos o segundo menor valor do vetor e trocamos de posição com o segundo elemento
  -repetindo esse procedimento conseguimos ordenar um vetor com a estratégia da seleção
*/

type SelectionSort struct {
	slice_values *[]int
}

func (s *SelectionSort) SortAscending(slice []int) {

	//atribuicao de cada de uma slice na outra, como usa-se &, entao o valor atribuido a s.slice_values eh o endereco de memorio do slice. referenciacao.
	s.slice_values = &slice //{4, 2, 7, 1, 5, 6, 5, 234, -2} o endereco do slice da main.go eh multivalorado, ou seja, o slice do parametro dessa funcao possui o mesmo endereco da main.go

	var min int

	for i := 0; i < len(*s.slice_values)-1; i++ {

		min = i

		for j := i + 1; j < len(*s.slice_values); j++ {

			if (*s.slice_values)[i] > (*s.slice_values)[j] {

				min = j // atualizar o index onde tem o valor menor da comparacao
			}
		}

		(*s.slice_values)[i], (*s.slice_values)[min] = (*s.slice_values)[min], (*s.slice_values)[i] //troca multipla de valores sem a necessidade de um terceira variavel temporaria
	}
}
