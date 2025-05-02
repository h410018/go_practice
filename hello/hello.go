package main

import (
	"fmt"
	"log"

	"boyuan.com/greetings"
)

// type MyClass struct {
// 	X, Y int
// }

// func (m *MyClass) testFunc() { // m *MyClass, m MyClass 的區別
// 	copy_m := m
// 	copy_m.X += 100
// 	fmt.Println(m.X, m.Y)
// }

func main() {
	// c := MyClass{2, 1}
	// (&c).testFunc()
	log.SetPrefix("greetings: ")
	log.SetFlags(log.Ldate | log.Ltime)

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// Get a greeting message and print it.
	// messages, err := greetings.Hello("黃柏源")
	messages, err := greetings.Hellos(names)
	if err != nil {
		fmt.Println("錯誤訊息: ", messages)
		log.Fatal("err 訊息: ", err)
	}
	fmt.Println(messages)

}
