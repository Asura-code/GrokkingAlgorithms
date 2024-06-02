package main

import "fmt"

func factorial(x int) int{
	if x == 1 || x == 0{
        return 1
    }else{
        return x * factorial(x-1)
    }
}

func main() {
  fmt.Println(factorial(5))
}