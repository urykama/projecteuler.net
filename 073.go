
//Counting fractions in a range
//Problem 73

package main

import (
    "fmt"
)

func main() {
// a<b<c
// a+b+c=1000
// a*a + b*b = c*c

var i, n, d  uint32
var    mn[64] uint32
var    md[64] uint32
for d =8; d > 1; d-- {
    for n =1; n < d; n++ {
        i = d*7 + n
        mn[i] = n
        md[i] = d
//        fmt.Println ("n=", n, "d=", d)


        }
        
    }

j:=1
mn[j] = 1
md[j] = 8
for i = 1; i<64; i++ {
    if md[i] > 0 {
        fmt.Println ("i =", i, "n=", mn[i], "d=", md[i])
        fmt.Println ("i =", i, "mn[i] / md[i]=", (mn[i]*100) / md[i], "mn[j] / md[j]=", mn[j] / md[j])
        if ((mn[i]*100) / md[i]) > ((mn[j]*100) / md[j]) {
            j++
            mn[j] = mn[i]
            md[j] = md[i]
        fmt.Println ("j =", j, "n=", mn[j], "d=", md[j])
        }
    }    
}

for i = 1; i<64; i++ {
    if md[i] > 0 {
        fmt.Println ("i =", i, "n=", mn[i], "d=", md[i])
    }    
}

}
