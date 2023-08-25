package main

// Interfaces contém funções
// Todos os tipos de dados que implementan os métodos de uma interface satisfazem essa interface implicitamente.
type Animal interface {
	speak()
}

type Dog struct {
	name string
}

func (d Dog) speak() {
	println("au au")
}

func main() {
	dog := Dog{"dog"}
	doSomething(dog)
}

func doSomething(animal Animal) {
	animal.speak()
}
