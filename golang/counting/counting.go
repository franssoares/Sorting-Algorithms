package counting

import "fmt"

type Counting struct {
	slice_values *[]int
}

func (cou *Counting) SortAscending(slice []int) {
	cou.CountingSort(slice)
}

func (cou *Counting) CountingSort(slice []int) {
	// Passagem
	cou.slice_values = &slice

	// Capturando min e max
	min, max := findMinMax((*cou.slice_values))
	fmt.Println("Índice min e max: ", min, max)

	// Criando slice auxiliar
	offset := -min // "offset" representa o deslocamento necessário para lidar com valores negativos
	intervalo := max - min + 1
	aux_slice := make([]int, intervalo)
	fmt.Println("Slice auxiliar criada: ", aux_slice)

	// Iterando o vetor auxiliar
	elementsFreq(slice, aux_slice, offset)
	fmt.Println("Slice auxiliar contada: ", aux_slice)

	// Soma acumulativa
	commulativeSum(aux_slice)
	fmt.Println("Slice auxiliar acumulado: ", aux_slice)

	// Ordenando o slice original
	sort(slice, aux_slice, offset)
	fmt.Println("Slice ordenado: ", slice)
}

// Funcoes auxiliares
func findMinMax(slice []int) (int, int) {

	n := len(slice)
	fmt.Println("Tamanho da slice original: ", n)

	minSlice := slice[0]
	maxSlice := slice[0]

	// Encontrando os maiores e menores valores da slice
	for i := 0; i < n; i++ {

		if minSlice > slice[i] {

			minSlice = slice[i]

		}
		if maxSlice < slice[i] {

			maxSlice = slice[i]

		}
	}

	fmt.Println("Valores min e max: ", minSlice, maxSlice)

	return minSlice, maxSlice
}

func elementsFreq(slice []int, aux_slice []int, offset int) {

	n := len(slice)

	for i := 0; i < n; i++ {

		aux_slice[slice[i]+offset]++

	}
}

func commulativeSum(aux_slice []int) {

	n := len(aux_slice)

	for i := 0; i < n-1; i++ {

		aux_slice[i+1] = aux_slice[i+1] + aux_slice[i]

	}
}

func sort(slice []int, auxSlice []int, offset int) {

	n := len(slice)
	fmt.Println(n)

	sortSlice := make([]int, n)
	fmt.Println("sortSlice criada: ", sortSlice)

	for i := n - 1; i >= 0; i-- {

		sortSlice[auxSlice[slice[i]+offset]-1] = slice[i]
		auxSlice[slice[i]+offset]--

	}

	// Copiar os elementos ordenados de volta para o slice original
	for i := 0; i < n; i++ {

		slice[i] = sortSlice[i]

	}
}
