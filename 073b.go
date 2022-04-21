//go run gorut.go
//Counting fractions in a range
//Problem 73

package main

import (
    "fmt"
)

func main() {

var max uint64 = 12000

const maxx  = 14400000
var j, i, n, d  uint64 
var    mn[maxx] uint64
var    md[maxx] uint64
//var mmn[maxx] uint32
//var mmd[maxx] uint32

i = 1
for d = max; d > 1; d-- {
    for n = 1; n < d; n++ {
        if (n*1200000/d) > 400000{
            if (n*1200000/d) < 600000{
                mn[i] = n
                md[i] = d
//                fmt.Println ("i =", i, "n=", n, "d=", d, mn[i], md[i])
                i++
            }
        }
    }
        
}  
im:=i-1

fmt.Println ("im =", im, "n=", mn[im], "d=", md[im])


//for i = im; i>0; i-- { fmt.Println ("i =", i, "mn=", mn[i], "md=", md[i]) }

j=0
var it  uint64 = 1
var itt uint64 = 120000000
for j < 901 {
    itt = 120000000
    for i = im; i>0; i-- {
//        fmt.Println ("1)  i =", i, "mn[i]=", mn[i], "md[i]=", md[i], "it=", it, "itt=", itt, (mn[i]*12000000/md[i]))
        if (mn[i]*120000000/md[i]) < itt {
            if (mn[i]*120000000/md[i]) > it {
                itt=(mn[i]*120000000/md[i])
//                fmt.Println ("2)  i =", i, "n=", n, "d=", d, mn[i], md[i])
            }
//            fmt.Println ("3)  i =", i, "mn[i]=", mn[i], "md[i]=", md[i], "it=",it, "itt=", itt)
        }
//        fmt.Println ("4)  j =", j, "mn[i]=", mn[i], "md[i]=", md[i], "it=",it, "itt=", itt) 
    }
    if it == 120000000 { break }
    it = itt
    j++
}
fmt.Println("im=", im, "j=", j)
//for i = im; i>0; i-- { fmt.Println ("mn=", mn[i], "md=", md[i], "i =", i, "j=", j) }

}
