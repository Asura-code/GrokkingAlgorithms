package main

import "fmt"

func fibonacci(x int) int{
    if x == 1{
        return 0
    }
    if x == 2{
        return 1
    }else{
        return fibonacci(x-1) + fibonacci(x-2)
    }
}

func main() {
  fmt.Println(fibonacci(12))
}