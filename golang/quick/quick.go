package quick

import (
	"math/rand"
	"time"
)

// Randomização do pivô
var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

type QuickSort struct {
	slice_values *[]int
}

func (qui *QuickSort) SortAscending(slice []int) {

	qui.slice_values = &slice

	n := len(slice)

	qui.QuickSortRec(slice, 0, n-1)

}

func (qui *QuickSort) QuickSortRec(slice []int, ini int, fim int) {
	if ini < fim {
		pivotIndex := qui.Partition(slice, ini, fim)
		qui.QuickSortRec(slice, ini, pivotIndex-1)
		qui.QuickSortRec(slice, pivotIndex+1, fim)
	}
}

func (qui *QuickSort) Partition(slice []int, ini int, fim int) int {

	// Generate a random number within the interval [int, fim]
	randomNum := globalRand.Intn(fim-ini+1) + ini

	// Troca o pivô aleatório com o último elemento
	aux := slice[fim]
	slice[fim] = slice[randomNum]
	slice[randomNum] = aux

	pivot := slice[fim]
	pIndex := ini

	for i := ini; i < fim; i++ {
		if slice[i] <= pivot {
			temp := slice[pIndex]
			slice[pIndex] = slice[i]
			slice[i] = temp

			pIndex++
		}
	}

	slice[fim] = slice[pIndex]
	slice[pIndex] = pivot
	return pIndex
}
