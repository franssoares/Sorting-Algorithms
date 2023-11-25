
# Por dentro do assunto (Call Stack)

A **Call Stack**, ou pilha de chamadas, é uma estrutura de dados na qual as informações relacionadas a chamadas de função são armazenadas durante a execução de um programa. A pilha de chamadas segue a abordagem **Last In, First Out (LIFO)**, o que significa que a última chamada que entra na pilha é a primeira a ser retirada.

Na pilha de chamadas, cada chamada de função cria um novo bloco de memória chamado "frame de pilha" ou "quadro de pilha". Este frame contém informações como:

**1. Parâmetros da função:** Valores passados para a função.
**2. Variáveis locais:** Variáveis declaradas dentro da função.
**3. Endereço de retorno:** O endereço da próxima instrução a ser executada após o término da função.
**4. Outras informações de controle:** Informações adicionais necessárias para o controle do fluxo de execução.

<h1 align="center"><img src="https://res.cloudinary.com/media-resource/image/upload/v1563577182/babscraig.com/call_stack_u2vvsi.jpg"/></h1>

-----

# Fluxo de Execução (MergeSort)

Aqui está um esboço do fluxo de execução durante a recursão como um exemplo, as funções são:

- > MergeSortRec(arr, left, right)
- > MergeRec(arr, left, middle, right)

### Chamada Inicial

- > MergeSortRec(arr, 0, 6)
- > A lista completa é arr = {38, 27, 43, 3, 9, 82, 10}.

**1. Lado esquerdo, Divisão em Metades (Primeira Chamada Recursiva)**

- > MergeSortRec(arr, 0, 3)
    - > Calcula middle como 1.
    - > Chamada recursiva para MergeSortRec(arr, 0, 1).
    - > Chamada recursiva para MergeSortRec(arr, 2, 3).

**2. Divisão da Metade Esquerda (Primeira Chamada Recursiva)**

- > MergeSortRec(arr, 0, 1)
    - > Calcula middle como 0.
    - > Chamada recursiva para MergeSortRec(arr, 0, 0).
    - > Chamada recursiva para MergeSortRec(arr, 1, 1).

**3. Ordenação da Metade Esquerda (Primeira Chamada Recursiva)**

- > MergeSortRec(arr, 0, 0) - Já que tem apenas um elemento, retorna imediatamente.
- > MergeSortRec(arr, 1, 1) - Já que tem apenas um elemento, retorna imediatamente.

**4. Vai para a terceira função, Mesclagem da Metade Esquerda (Primeira Chamada Recursiva)**

- > MergeRec(arr, 0, 0, 1) - Mescla a metade esquerda ordenada.

**5. Divisão da Metade Direita (Primeira Chamada Recursiva)**

- > MergeSortRec(arr, 2, 3)
    - > Calcula middle como 2.
    - > Chamada recursiva para MergeSortRec(arr, 2, 2).
    - > Chamada recursiva para MergeSortRec(arr, 3, 3).

**6. Ordenação da Metade Direita (Primeira Chamada Recursiva)**

- > MergeSortRec(arr, 2, 2) - Já que tem apenas um elemento, retorna imediatamente.
- > MergeSortRec(arr, 3, 3) - Já que tem apenas um elemento, retorna imediatamente.

**7. Vai para a terceira função, Mesclagem da Metade Direita (Primeira Chamada Recursiva)**

- > MergeRec(arr, 2, 2, 3) - Mescla a metade direita ordenada.

**8. Mesclagem Final (Primeira Chamada Recursiva)**

- > MergeRec(arr, 0, 1, 3) - Mescla as duas metades ordenadas.

**1. Lado direito, Divisão em Metades (Segunda Chamada Recursiva)**

- > MergeSortRec(arr, 4, 6)
    - > Calcula middle como 5.
    - > Chamada recursiva para MergeSortRec(arr, 4, 5).
    - > Chamada recursiva para MergeSortRec(arr, 6, 6).

**2. Ordenação da Metade Esquerda (Segunda Chamada Recursiva)**

- > MergeSortRec(arr, 4, 5) - Ambas as partes têm apenas um elemento, retornam imediatamente.

**3. Ordenação da Metade Direita (Segunda Chamada Recursiva)**

- > MergeSortRec(arr, 6, 6) - Já que tem apenas um elemento, retorna imediatamente.

**4. Vai para a terceira função, Mesclagem da Metade Direita (Segunda Chamada Recursiva)**

- > MergeRec(arr, 4, 5, 6) - Mescla a metade direita ordenada.

**5. Mesclagem Final (Segunda Chamada Recursiva)**

- > MergeRec(arr, 4, 5, 6) - Mescla as duas metades ordenadas.

**6. Mesclagem Final (Chamada Inicial)**

- > MergeRec(arr, 0, 3, 6) - Mescla as duas metades ordenadas.

Ao final desse processo, a lista inicial {38, 27, 43, 3, 9, 82, 10} está ordenada. Cada chamada recursiva divide a lista, ordena cada parte individualmente e, em seguida, mescla as partes ordenadas. Esse processo continua até que toda a lista esteja ordenada.

<h1 align="center"><img src="https://miro.medium.com/v2/resize:fit:720/format:webp/1*7Kox4Bll0Ddvb0td1tiXsg.png"/></h1>

Os arrays em azul são do método de divisão, enquanto os arrays em amarelo são contruídos pelo método da conquista.
