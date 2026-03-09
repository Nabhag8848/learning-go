package main

import (
	"fmt"
	"time"
)

func ControlConstruct() {
	i := 1 
	for i < 3 {
		i = i + 1
		i++
	}
	fmt.Println(i)

	for j := 0; j < 3;j++ {
		fmt.Println(j)
	}

	for k := range 3 {
		if k % 2 == 0 {
			continue
		}
		fmt.Println("range", k)
	}

	for {
		break;
	}


	if num := 9; num < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }

	animals := []string{
		"peacock",
		"elephant",
	}

	for i := range len(animals) {
		switch animals[i] {
			case "peacock":
				fmt.Println("peacock")
			case "elephant":
				fmt.Println("elephant")
		}
	}


    switch time.Now().Weekday() {
    case time.Saturday, time.Sunday:
        fmt.Println("It's the weekend")
    default:
        fmt.Println("It's a weekday")
    }

    t := time.Now()
    switch {
    case t.Hour() < 12:
        fmt.Println("It's before noon")
    default:
        fmt.Println("It's after noon")
    }

    whatAmI := func(i interface{}) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
	
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")
	
}