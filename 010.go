//go run gorut.go
//Summation of primes
//Problem 010
//The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.
//Find the sum of all the primes below two million.

package main

import ("fmt"
        "math"
        "time")

func main() {
t0 := time.Now();
x:=0
zz:=10
ii:=0

for ii = 3; ii < 4000000; ii = ii + 2 {
//    if ii % 2 == 0 {
//            continue
//        }
    if ii % 3 == 0 {
            continue
        }
    if ii % 10 == 5 {
        continue
    }
    x = int (math.Sqrt (float64(ii))+1)
    for i := 2; i < x; i++{
        if ii%i > 0 {
            if i == x-1 {
                zz = zz + ii
//                fmt.Println ( "ii=", ii, "i=", i, "fuck", ii%i)
            }
        } else {
                break
            }
    }
}
fmt.Println (ii, "Сумма Простых равна", zz)
t1 := time.Now();
fmt.Printf("Elapsed time: %v", t1.Sub(t0));
}