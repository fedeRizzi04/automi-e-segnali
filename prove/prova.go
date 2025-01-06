package main

import "fmt"

type Punto struct {
	x int
	y int
}

func main() {

	var m map[Punto]string = make(map[Punto]string)

	m[Punto{1, 2}] = "ciao"
	m[Punto{2, 3}] = "ciao2"

	s := m[Punto{1, 2}]
	fmt.Println(s)

	s = m[Punto{1, 11}]
	fmt.Println(s)
	if s == "" {
		fmt.Println("non c'è")
	}

	P := Punto{1, 2}
	p2 := Punto{1, 2}
	if P == p2 {
		fmt.Println("uguali")
	} else {
		fmt.Println("diversi")
	}

}
