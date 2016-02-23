package main

import ("fmt")

//0, 1, 1, 2, 3, 5, 8, 13, 21, 34
func fibonacci(n uint) uint {
    if n == 0 {
        return 0
    }
    if n == 1 {
        return 1
    }
     
    return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
    fmt.Println(fibonacci(0))
    fmt.Println(fibonacci(1))
    fmt.Println(fibonacci(2))
    fmt.Println(fibonacci(3))
    fmt.Println(fibonacci(4))
    fmt.Println(fibonacci(5))
    fmt.Println(fibonacci(6))
    fmt.Println(fibonacci(7))
    fmt.Println(fibonacci(8))
    fmt.Println(fibonacci(9))
}