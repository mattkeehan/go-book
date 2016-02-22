/*
Print numbers from 1 to 100, replace divisible-by-3's with fizz, 5's with buzz, 3 AND 5 with fizzbuzz
*/
package main

import "fmt"

func main() {
    i := 1
    for i <= 100 { 
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("fizzbuzz")
        } else if i % 3 == 0 {
            fmt.Println("fizz")
        } else if i % 5 == 0 {
            fmt.Println("buzz")
        } else {
            fmt.Println(i)
        }
        
        i=i+1 
    }
}