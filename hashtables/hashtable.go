package main

import "fmt"

func main() {
  voted := make(map [string] bool)
  voted["Tom"] = true
  voted["Mike"] = true
  for true{
      var input string
      fmt.Scanln(&input)
          if voted[input]{
              fmt.Print("kick him out\n")
          } else{
              fmt.Print("let him vote\n")
              voted[input] = true
      }
  }
}