package interfaces

// Em Go, as interfaces não precisam ser explicitamente importadas como pacotes.
// As interfaces são parte da linguagem e estão disponíveis globalmente em qualquer arquivo Go no mesmo pacote.
// Nesse caso -> main/interfaces, para todo pacote em -> main/

type Sort interface {
	SortAscending(slice []int)
}
