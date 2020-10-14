package main

import "fmt"

func main()  {
	var e float32
	var infinito int
	e = 0

	fmt.Scanln(&infinito)

    for i := 0; i <= infinito; i++ {
		if i == 0 {
			e = e + 1
		}else{
			e = e + 1/float32(factorial(i));
		}
    }
    fmt.Println("e = ",e)
}
func factorial(n int) int{
    var f int;
    var i int;

    f = 1
    for i = n; i >= 1; i-- {
        f = f * i;
    }

    return f
}
