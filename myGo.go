package main

// import (
// 	"fmt"
// 	"time"
// )

// var test int = 2135

// // struct
// type Person struct {
// 	firstName string
// 	lastName  string
// 	favorites []string
// }

// // interface
// type GetInfo interface {
// 	personInfo() string
// }

// // method of struct
// func (p Person) personInfo() string {
// 	// fmt.Println(p.firstName)
// 	return p.firstName + p.lastName
// }

// // function
// func showInfo(g GetInfo) {
// 	fmt.Println("go:" + g.personInfo())
// }

// // generics function
// func generFunc[T any](v T) T {
// 	fmt.Println(v)
// 	return v
// }

// func testRoutine(str string) {
// 	i := 0
// 	for {
// 		fmt.Println(i)
// 		i++
// 		time.Sleep(1 * time.Second)
// 	}

// }

// func main() {
// 	num := []int{10, 20, 30, 40, 50}
// 	num = append(num, 2222)
// 	fmt.Println(num)
// 	var p = Person{
// 		firstName: "Boyuan",
// 		lastName:  "Huang",
// 		favorites: []string{"a", "b", "c"},
// 	}

// 	// p.printInfo()
// 	var x = 10
// 	var address *int = &x
// 	fmt.Println(address) // print address of variable x
// 	showInfo(p)

// 	go testRoutine("123")

// 	for i := 0; i < len(num); i++ {
// 		fmt.Println(num[i])
// 	}
// }
