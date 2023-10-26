package main

import "fmt"

func main() {

	//var tablica []string
	tablica := []string{"Ala", "kot", "Iza", "Ala", "dom"}
	var mapa = make(map[string]int, len(tablica))
	var index int

	for _, v := range tablica {
		mapa[v] = index
		index++
	}

	for v, _ := range mapa {
		fmt.Println(v)
	}
}
