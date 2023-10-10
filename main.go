package main

import (
	"fmt"
)

func main() {
    aray := []int{2,3,4,5,6}
    aray = append(aray, 10,29)
    fmt.Println(aray)
}


func variadic (num ...int) int{
    total := 0
    for i := 0; i < len(num); i++ {
        total += num[i]
    }
    return total
}
