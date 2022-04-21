//Special Pythagorean triplet
//Problem 9
package main

import (
    "fmt"
)

func main() {
// a<b<c
// a+b+c=1000
// a*a + b*b = c*c

var a, b, c  uint32


for a =1; a < 334; a++ {
    for b =a+1; b < 500; b++ {
        c = 1000 - a - b
        if b < c {
            if a*a + b*b == c*c {
                fmt.Println ("a=", a, "b=", b, "c=", c, "a*b*c=", a*b*c)
            }
        }
    }
}
}