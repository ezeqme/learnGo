package main

import "fmt"

const prefixoOlaPortugues = "Olá, "

// quando você programa em Go há um pacote main definido com uma função principal main dentro dele
// os pacotes são maneiras de agrupar códigos Go relacionados
// a palavra reservada func é usada para definir uma função com um nome e conteúdo
// ao usar import "fmt" estamos importando um pacote que contém a função Println que será usada para imprimir um valor na tela
func Ola(nome string) string {
	if nome == "" {
		nome = "mundo"
	}
	return prefixoOlaPortugues + nome
}

func main() {
	fmt.Println(Ola("Chris"))
}
