package main

import "fmt"

func main(){
    var n int 
    fmt.Scan(&n)
    if n%2 !=0{
        fmt.Println("Weird")
    }else if n%2 == 0 && n>=2 && n<=5{
        fmt.Println("Not Weird")
    }else if n%2 == 0 && n>=6 && n<=20{
        fmt.Println("Weird")
    }else if n%2 == 0 && n>20 {
        fmt.Println("Not Weird")
    }
}
