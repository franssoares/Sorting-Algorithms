<h1 align="center"><img src="https://technology.riotgames.com/sites/default/files/articles/116/golangheader.png"/></h1>

# Uso do Golang
Para executar o programa a partir do terminal, utilize o seguinte comando

#### Executar o Programa

```bash
go run main
```

#### Testes e Benchmarks

Para realizar testes unitários, utilize o seguinte comando:

```bash
go test -v
```

Para executar benchmarks, use o seguinte comando:

```bash
go test -bench .
```

Este comando executará benchmarks para medir o desempenho dos algoritmos. No entanto, se você deseja executar apenas os benchmarks sem executar os testes unitários, você pode usar o seguinte comando:

```bash
go test -bench . -run XXX
```

O XXX é apenas um padrão que não corresponderá a nenhum teste, garantindo que apenas os benchmarks sejam executados.

> **obs:** Certifique-se de estar no diretório correto onde o código do algoritmo está localizado ao executar esses comandos.
