
//Counting fractions in a range
//Problem 71
package main
import (   "fmt"   ;   "time"   )

func main() {
	t0 := time.Now();
var n, d int
n = 2
d = 5
for i := 0; i < 300000; i++ {
    n = n + 3
    d = d + 7
//    fmt.Println ( 3*1000000/7 )
    
    if ( d >= 1000000) {
    	fmt.Println ( "i=", i, "n=", n, "d=", d, n*1000000/d )
        break
    }
}

fmt.Println ( 3*1000/7 )
fmt.Println ( n-3, d-7 )
fmt.Println ( (1000000-1)/7  )
fmt.Println ( (1000000-1)/7*3  )
fmt.Println ( (1000000-1)/7*7  )
t1 := time.Now();
fmt.Printf("Elapsed time: %v", t1.Sub(t0));
}
