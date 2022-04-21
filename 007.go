//go run gorut.go
//10001st prime
//Problem 007

package main

import ("fmt")

func main() {

z:=0
zz:=0
ii:=0

for j:=1; j< 105000; j++ {
    ii ++
    z = 0
//    fmt.Println (z,ii, zz)
    for i := 2; i < ii; i++{
//        fmt.Println ( ii, i, "ii%i=", ii%i)
        if ii%i > 0 {
//            fmt.Println ( ii, i, "ii%i=", ii%i)
            z = z + 1
        }
    }
    
        
    if z == ii-2{
        zz = zz + 1
        fmt.Println (zz, "Простое", ii)
    }
    
    if zz == 10001{
        break
    }
    
    
}
fmt.Println (zz, "Простое", ii)
}