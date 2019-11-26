package main


import (
  "fmt"
  "os"
  "bufio"
  "strconv"
)

func main() {
  	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".
  
    var i uint64 = 4
    var d float64 = 4.0
    var s string = "HackerRank "

    scanner := bufio.NewScanner(os.Stdin)
    // Declare second integer, double, and String variables.
    var i2 int; var d2 float64; var s2 string
    // Read and save an integer, double, and String to your variables.
    if scanner.Scan(){
        i2, _ = strconv.Atoi(scanner.Text())
    }    
    if scanner.Scan(){
        d2, _ = strconv.ParseFloat(scanner.Text(), 64)
    }
    if scanner.Scan(){
        s2 = scanner.Text()
    }
    // fmt.Println(i,d,s,i2,d2,s2)
    // Print the sum of both integer variables on a new line.
    fmt.Println(int(i)+i2)
    // Print the sum of the double variables on a new line.
    fmt.Println(float64(d+float64(d2)))
    // Concatenate and print the String variables on a new line
    fmt.Println(s, s2)
    // The 's' variable above should be printed first.

}
