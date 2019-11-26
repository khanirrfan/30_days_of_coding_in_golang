package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    //Enter your code here. Read input from STDIN. Print output to STDOUT
    var str string
    scanner := bufio.NewScanner(os.Stdin)

    if scanner.Scan() {
        str = scanner.Text()
    }
    fmt.Println("Hello, World.")
    fmt.Println(str)
    // os.Stdout("Hello, World. \n", str)
}
