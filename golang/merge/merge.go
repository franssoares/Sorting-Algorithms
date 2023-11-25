package merge

type MergeSort struct {
	slice_values *[]int
}

// Inicio -> chama esse metodo para implementar o mergesort
func (mer *MergeSort) SortAscending(slice []int) {

	/*
	* vetor a ser ordenado, primeiro indice do vetor,
	* ultimo incice do vetor
	**/

	n := len(slice)

	mer.MergeSortRec(slice, 0, n-1)

}

// Divisao
func (mer *MergeSort) MergeSortRec(slice []int, left int, right int) {

	mer.slice_values = &slice

	// Se left for menor que right, isso significa que ha mais de um elemento
	if left < right {

		middle := left + (right-left)/2

		/*
		* Recursao:
		* A pilha de chamadas (call stack) segue a abordagem de
		* última entrada, primeira saída (Last In, First Out - LIFO)
		**/

		// Divisão da Metade Esquerda
		mer.MergeSortRec((*mer.slice_values), left, middle)

		// Divisão da Metade Direita
		mer.MergeSortRec((*mer.slice_values), middle+1, right)

		// Ordenando e mesclando
		mer.MergeRec((*mer.slice_values), left, middle, right)
	}
}

// Conquistando a morena
func (mer *MergeSort) MergeRec(slice []int, left int, middle int, right int) {

	/*
	* Para nao complicar a nossa vida não tomarei mer.slice_values = &slice
	* Ira ver slice
	**/

	n1 := middle - left + 1
	n2 := right - middle

	// Criação de slices temporárias
	leftSlice := make([]int, n1)
	rightSlice := make([]int, n2)

	// Copiar dados para as slices temporárias
	for i := 0; i < n1; i++ {
		leftSlice[i] = slice[left+i]
	}

	for j := 0; j < n2; j++ {
		rightSlice[j] = slice[middle+1+j]
	}

	// Índices iniciais dos subarrays
	i := 0
	j := 0

	// Índice inicial do subarray fundido
	k := left

	// Mescla os subarrays de volta no slice original
	for i < n1 && j < n2 {
		if leftSlice[i] <= rightSlice[j] {
			slice[k] = leftSlice[i]
			i++
		} else {
			slice[k] = rightSlice[j]
			j++
		}
		k++
	}

	// Copia os elementos restantes de leftSlice, se houver
	for i < n1 {
		slice[k] = leftSlice[i]
		i++
		k++
	}

	// Copia os elementos restantes de rightSlice, se houver
	for j < n2 {
		slice[k] = rightSlice[j]
		j++
		k++
	}
}
