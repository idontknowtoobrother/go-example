package interfaces

import "fmt"

// Speaker interface
type Speaker interface {
	Speak() string
}

// Dog struct
type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type Person struct {
	Name string
}

// Person's implementation of the Speaker interface
func (p Person) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func MakeSound(s Speaker) {
	fmt.Println(s.Speak())
}
