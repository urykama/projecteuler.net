
//Ordered fractions
//Problem 72
package main
import (   "fmt"   ;   "time"   )

func fractions (n, d, nn, dd int, c chan int) {
	out := 0
	nnn := n + nn //; fmt.Print (nnn, " / " )
	ddd := d + dd //; fmt.Println ("          ", ddd)
	if ddd > 8 { c <- out; return }
//	fmt.Println ( nnn, " / ", ddd, "out =", out)
	go fractions (n, d, nnn, ddd, out )
	out = out + c
	go fractions (nnn, ddd, nn, dd, out )
	out = out + c
	c <- out
return 
}



func main() {
	c := make(chan int, 10)
	t0 := time.Now();
	out := 0
	go fractions (1, 8, 1, 7, out)
	out = out + c
	
		fmt.Println ( "out ", out )
//	out = out + fractions (1, 5, 1, 4, out)
//		fmt.Println ( "out ", out )
//	out = out + fractions (1, 4, 2, 7, out)
//		fmt.Println ( "out ", out )
//	out = out + fractions (2, 7, 1, 3, out)
//		fmt.Println ( "out ", out )
//	out = out + fractions (1, 3, 3, 8, out)
//		fmt.Println ( "out ", out )
//		fmt.Println (( 37995444063+5427920629+7237227493+10132118373+3546241436+81780669928+178035794615+368736737200+7599088909+8684672867+21711682301) * 2 + 21)




t1 := time.Now();
fmt.Printf("Elapsed time: %v", t1.Sub(t0));
}

// cd C:\Users\User\Go\ProjectEuler
// go run g.go
// 5066084535
// 	0/1 	1/8 	 37995444063
//	1/8		1/7		  5427920629
//	1/7		1/6		  7237227493
//	1/6		1/5		 10132118373
//	1/5		1/4		 35462414361
//	1/4		2/7		 81780669928
//	2/7		1/3		178035794615
//	1/3		3/8		368736737200
//	3/8		2/5		  7599088909
//	2/5		3/7		  8684672867
//	3/7		1/2		 21711682301
//				   1461775195649