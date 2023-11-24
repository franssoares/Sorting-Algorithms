// Em Go, um programa executável deve sempre ter um pacote chamado main
package main

// O import eh usado para chamar bibliotecas ou pacotes interno ou externo, NAMESPACE assim como o uso do std::cin() em c++, em go se usa nome_do_pacote.metodo()
// Colocar em import pelo caminho do diretorio, usar os metodos pela nome do pacote
// Cada diretorio deve ter no maximo um pacote sendo usado
import (
	"fmt"
	"main/selection" // Perceba que o package main esta no repositorio recorrente, e selection eh o package acessado
)

func main() {

	var sel selection.SelectionSort // Namespace assim como o uso do std::cin() em c++
	slice := []int{4, 2, 7, 1, 5, 6, 5, 234, -2}

	fmt.Println("Slice original antes:", slice)

	// Chamar o método com a slice
	sel.SortAscending(slice)

	fmt.Println("Slice original depois:", slice)
}
