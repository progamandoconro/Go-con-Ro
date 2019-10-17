package main

import ("fmt"
        "time"
)

func Fibonacci(n uint) uint64 {
        if n <= 1 {
                return uint64(n)
        }

        var n2, n1 uint64 = 0, 1

        for i := uint(2); i < n; i++ {
                n2, n1 = n1, n1+n2
        }

        return n2 + n1
}

func main (){
        var nth uint
        fmt.Println("Nth Fibonacci ?")
        fmt.Scan(&nth)
        time.Sleep(time.Second * 1 )

        for i:= uint(1); i<nth;i++{

        fmt.Println(Fibonacci(i))
        time.Sleep(500 * time.Millisecond)
        }

}



