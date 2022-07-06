package main

import (
	"fmt"
	"strconv"
    "math"
)

func main(){
    var a,b float64
    aStr := "0.1"
    value,err := strconv.ParseFloat(aStr,64)
    if err != nil{
        fmt.Println(value,err)
    }
    a = 0.8 - value-value
    b = 0.6
    fmt.Println(a,b)
    a = math.Round(a*100) / 100
    fmt.Println(a,b)
}
