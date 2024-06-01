package main

import "fmt"

func selection_sort(a []int) []int{
    lenght := len(a)
    for i := 0; i< lenght; i++{
        min := i
        for j:= i+1; j < lenght; j++{
            if a[j] <= a[min]{
                min = j
            }
        }
        a[i], a[min] = a[min], a[i]
    }
    return a
}

func main() {
  m := []int{1,3,2,4,8,6,5,1,8}
  fmt.Print(selection_sort(m))
}