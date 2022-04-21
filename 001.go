//https://projecteuler.net/problem=1
package main

import (    "fmt"	)

func main() {

	s3 := 0
    for i := 3; i < 1000; i += 3{
    	s3 += i
    	fmt.Println(i, s3)
    }

    s5 := 0
    for i := 5; i < 1000; i += 5{
    	s5 += i
    	fmt.Println(i, s5)
    }

    s15 := 0
    for i := 15; i < 1000; i += 15{
    	s15 += i
    	fmt.Println(i, s15)
    }

    fmt.Println(s3 + s5 - s15)
}