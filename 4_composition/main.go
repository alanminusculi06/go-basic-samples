package main

//A composição é um método empregado para escrever segmentos de código reutilizáveis.
type Animal interface {
	speak()
}

type AnimalImpl struct {
	Name string
}

type Dog struct {
	AnimalImpl
}

type Cat struct {
	AnimalImpl
}

func (d Dog) speak() {
	println("au au")
}

func (c Cat) speak() {
	println("miau")
}

func main() {
	dog := Dog{AnimalImpl{"dog"}}
	cat := Cat{AnimalImpl{"cat"}}

	doSomething(dog, cat)
}

func doSomething(animal ...Animal) {
	for _, a := range animal {
		a.speak()
	}
}
