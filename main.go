package main

import "fmt"

func main()  {
	var e float32
    e = 0
    for i := 0; i < 30; i++ {
        e = e + 1/float32(factorial(i));
    }
    fmt.Println("e = ",e)
}
func factorial(n int) int{
    var f int;
    var i int;

    if n == 0{
        f = 1;
    } else{
        f = 1
        for i = n; i >= 1; i-- {
            f = f * i;
        }
    }
    return f
}
