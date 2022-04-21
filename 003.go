//Largest prime factor
package main

import (
    "fmt"
)

func main() {

	n := 600851475143
//    n := 13195

    for i :=2; i*i < n; i++ {
        if n%i == 0 {
            fmt.Println(n , "/", i, "=", n/i)
            for ii :=2; ii*ii < i; ii++{
                if (i) % ii == 0 {
                    fmt.Println("no", n, i, "/", ii, "=", i/ii)
                }
            }
        }
    }
}