// todo arquivo execultavel em go deve ter o pacote main para a compilacao funcionar
// Em Go, um programa executável deve sempre ter um pacote chamado main
package main

//o import eh usado para chamar bibliotecas ou pacotes interno ou externo, NAMESPACE assim como o uso do std::cin() em c++, em go se usa nome_do_pacote.metodo()
//colocar em import pelo caminho do diretorio, usar os metodos pela nome do pacote
//cada diretorio deve ter no maximo um pacote sendo usado
import (
	"fmt"
)

func main() {

	var sel selectionSort
	slice := []int{4, 2, 7, 1, 5, 6, 5, 234, -2}

	fmt.Println("Slice original antes:", slice)

	// Chamar o método com a slice
	sel.sortAscending(slice)

	fmt.Println("Slice original depois:", slice)
}
