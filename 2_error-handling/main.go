package main

import (
	"errors"
	"fmt"
)

// "O acoplamento de exceções a uma estrutura de controle, resulta em um código complicado
// e também induz programadores a tratar erros comuns, como excepcionais".
// O Go também possui algumas funções integradas para sinalizar e se recuperar de
// condições verdadeiramente excepcionais, como Defer(), Panic() e Recover().
func main() {
	_, err := doSomething()
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func doSomething() (string, error) {
	//...
	return "", errors.New("some error")
}
