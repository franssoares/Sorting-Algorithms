package bubble

/**
Estratégia do algoritmo de acordo com seu nome
-fazer os maiores valores do vetor flutuarem para o final do vetor;
-assim como bolhas de ar embaixo d'água flutuam mais rapidamente para a superfície de acordo com seu tamanho, os maiores valores do vetor devem "flutuar" mais rapidamente para o fim do vetor;
-isso é tudo que precisamos lembrar para começar a estudar o bubblesort.
*/

type Bubblesort struct {
	slice_values *[]int
}

func (b *Bubblesort) SortAscending(slice []int) {

	b.slice_values = &slice

	n := len(*b.slice_values)

	for i := 0; i < n-1; i++ {

		swapped := false

		for j := i + 1; j < n; j++ {

			if (*b.slice_values)[i] > (*b.slice_values)[j] {

				swapped = true
				(*b.slice_values)[i], (*b.slice_values)[j] = (*b.slice_values)[j], (*b.slice_values)[i]
			}
		}

		// Se o vetor esta ordenado entao pare a iteracao -> Pode-se alterar para um _for swapped{} (aprimoramento)
		// n-- // Uma ~bolha~ sobe a cada varredura, entao o vetor diminui a cada varredura (otimizado)
		if !swapped {
			break
		}
	}
}
