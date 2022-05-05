package main

import "fmt"

type person struct{
	name string
	age int
	favFood []string
}

func main() {
	foodArray := []string {"abc","def"}
	choe := person{"jaehyun", 32, foodArray}

	fmt.Println(choe)
}