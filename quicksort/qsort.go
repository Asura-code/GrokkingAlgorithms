package main

import (
    "fmt"
    )

func glue(a []int,b []int, c int) []int{
    a = append(a,c)
    for i := 0; i< len(b); i++{
        a = append(a,b[i])
    }
    return a
}

func qsort(a []int) []int{
    less := []int{}
    bigger := []int{}
    if len(a) < 2{
        return a
    }else{
        pivot := int(a[0])
        for i:=1; i<len(a); i++{
            if a[i]<= pivot{
                less = append(less, a[i]) 
            }else if a[i]> pivot{
                bigger = append(bigger, a[i])
            }
        }
        return glue(qsort(less),qsort(bigger),pivot)
    }
}

func main() {
    a := []int{10,-8,10,123,4,56,19,0,-8,3}
    fmt.Print(qsort(a))
}