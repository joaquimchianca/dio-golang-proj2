package main

import "fmt"

func parte1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func parte2() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0:
			fmt.Println("Pin")
		case i%5 == 0:
			fmt.Println("Pan")
		default:
			fmt.Println(i)
		}
	}
}

func main() {
	parte2()
}
