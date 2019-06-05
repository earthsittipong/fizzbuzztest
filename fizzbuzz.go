package main

import (
	"fmt"
)

func main() {
println("hello world")
//var i int = 100

/* for loop number 1-100 */
for i:=1;i<=100;i++{

if i%15 == 0 {
    fmt.Println("type of ",i,"=FizzBuzz")
}else if i%5 == 0{
    fmt.Println("type of ",i,"=Buzz")
}else if i%3 == 0 {
        fmt.Println("type of ",i,"=Fizz")
} 


 

}
}