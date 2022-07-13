package main

import (
      "fmt"
)

func SieveOfEratosthenes(n int) []int {
      integers := make([]bool, n+1) //create slice to store if number is prime
      for i := 2; i < n+1; i++ {
              integers[i] = true
      }

      for p := 2; p*p <= n; p++ {
             if integers[p] == true {
                     for i := p * 2; i <= n; i += p {
                              integers[i] = false
                      }
              }
      }

      // return all prime numbers <= n
      var primes []int
      for p := 2; p <= n; p++ {
              if integers[p] == true {
                      primes = append(primes, p)
              }
      }

      return primes
}

func main() {
      fmt.Println(SieveOfEratosthenes(10))
}
