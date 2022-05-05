package main

import (
  "fmt"
)

func fib(n int) int {
  if n <= 1 {
    return n
  } else {
    return fib(n-1) + fib(n-2)
  }
}

func main() {
  var n int = 0

  for true {
    fmt.Printf("F(%d) = F(%d)-1 + F(%d)-2\nF(%d) = %d\n", n, n, n, n, fib(n))
    n++
  }

}
