package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"
)

func SlicesIndex[S ~[]E, E comparable](s S, v E) int {
	for i := range s {
		if s[i] == v {
			return i
		}
	}

	return -1
}

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val T
}

func (l *List[T]) Push(v T) {
	if l.tail == nil {
		l.head = &element[T]{val: v}
		l.tail = l.head
	} else {
		l.tail.next = &element[T]{val: v}
		l.tail = l.tail.next
	}
}

func (lst *List[T]) AllElements() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

func (lst *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e:= lst.head; e != nil; e=e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

	
func genFib() iter.Seq[int] {
    return func(yield func(int) bool) {
        a, b := 0, 1
        for {
            if !yield(a) {
                return
            }
            a, b = b, a+b
        }
    }
}

func Generics() {
	var s = []string{"foo", "bar", "zoo"}
    fmt.Println("index of zoo:", SlicesIndex(s, "zoo"))

	lst := List[int]{}
    lst.Push(10)
    lst.Push(13)
    lst.Push(23)

	fmt.Println("list:", lst.AllElements())

	for e := range lst.All() {
		fmt.Print(e, ",")
	}

	all := slices.Collect(lst.All())
    fmt.Println("all:", all)

	for part := range strings.SplitSeq("go-by-example", "-") {
        fmt.Printf("part: %s\n", part)
    }

	for n := range genFib() {
		
		if n >= 10 {
			break
		}
		fmt.Println(n)
	}
}