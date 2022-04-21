//Even Fibonacci numbers
package main

import (
    "fmt"
//    "hash/crc32"
)

func main() {

	s := 2
    f1 := 1
    f2 := 2
    f := 0

    for {
        f = f1 + f2
        f1 = f2
        f2 = f
        if f%2 == 0 {
            s = s + f
        }

        if f > 4000000 {
            break
        }

    	fmt.Println(f, s)
    }
}