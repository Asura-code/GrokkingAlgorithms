package main

import (
    "fmt"
    "strconv"
    )

func fromAtoB(a int, b int) string{
    if a >= b || b - a == 1{
        return "empty"
    }
    if b - a == 2{
        return strconv.Itoa(b-1)
    }else{
        return fromAtoB(a,b-1) + "," + strconv.Itoa(b-1)
    }
}

func main() {
  fmt.Println(fromAtoB(19,33))
}