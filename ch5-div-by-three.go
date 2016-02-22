/*
Numbers from 1 to 100 evenly divisible by 3
*/
package main

import "fmt"

func main() {
    i := 1
    for i <= 100 {
        if i % 3 == 0 {
            fmt.Println(i, " is divisible by 3")
        }
        
        i=i+1 
    }
}