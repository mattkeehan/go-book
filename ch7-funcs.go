package main

import ("fmt")

func average(xs []float64) float64 {
     total := 0.0
     for _, v := range xs {
           total += v
}
     return total / float64(len(xs))
}

func factorial(x uint) uint {
     if x == 0 {
         return 1 
     }
     
    return x * factorial(x-1)
}

func main() {
    xs := []float64{0,10}
    fmt.Println(average(xs))
    
    fmt.Println()
    //fac := 2
    fmt.Println(factorial(4))
}