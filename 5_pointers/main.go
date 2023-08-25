package main

//Como em todas as linguagens da família C, tudo em Go é passado por valor.
//Passar um valor int para uma função faz uma cópia do int, e passar um valor de ponteiro faz uma cópia do ponteiro.

// * precedendo um tipo indica que é um ponteiro para esse tipo.
// & precedendo uma variável retorna o endereço de memória da variável.
// * precedendo uma variável retorna o valor da variável apontada pelo ponteiro.
func main() {
	var text string
	var pointer *string

	text = "Hello World"
	pointer = &text

	println(text)
	println(pointer)
	println(*pointer)

	person := Person{"John"}
	println(person.name)
	doSomething(&person)
	println(person.name)
}

type Person struct {
	name string
}

func doSomething(person *Person) {
	person.name = "Peter"
}
