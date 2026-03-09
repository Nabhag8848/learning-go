package main

import (
	"fmt"
	"maps"
)

func Maps() {
	m := make(map[string]int)

	m["k1"] = 1
	m["k2"] = 2

	fmt.Println("map:", m)

	v1 := m["k1"]
    fmt.Println("v1:", v1)

    v3 := m["k3"]
    fmt.Println("v3:", v3)

	_, prs := m["k3"]
    fmt.Println("is k3 exist:", prs)

	fmt.Println("len:", len(m))

	delete(m, "k1") // if it exists
	fmt.Println("map:", m)

	clear(m)
    fmt.Println("map:", m)

	n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}

	checkEquality(n ,n2)

	for k, v := range n2 {
		fmt.Printf("%s -> %d", k, v)
		fmt.Println()
	}
}


func checkEquality(n, n2 map[string]int){
	fmt.Println(maps.Equal(n, n2))
}