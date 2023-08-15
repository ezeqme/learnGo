package main

import "fmt"

const espanhol = "espanhol"
const frances = "frances"
const esperanto = "esperanto"
const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const prefixoOlaEsperanto = "Saluton, "

// quando você programa em Go há um pacote main definido com uma função principal main dentro dele
// os pacotes são maneiras de agrupar códigos Go relacionados
// a palavra reservada func é usada para definir uma função com um nome e conteúdo
// ao usar import "fmt" estamos importando um pacote que contém a função Println que será usada para imprimir um valor na tela
func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	return prefixoSaudacao(idioma) + nome
}

func prefixoSaudacao(idioma string) (prefixo string) {
	switch idioma {
	case frances:
		prefixo = prefixoOlaFrances
	case espanhol:
		prefixo = prefixoOlaEspanhol
	case esperanto:
		prefixo = prefixoOlaEsperanto
	default:
		prefixo = prefixoOlaPortugues
	}
	return
}

func main() {
	fmt.Println(Ola("Chris", ""))
}
