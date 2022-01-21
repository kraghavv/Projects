package main

import (
"bufio"
"fmt"
"os"
"strconv"
)

var res float64

var count int = 0

func main() {

scanner := bufio.NewScanner(os.Stdin)
for {
fmt.Print("Enter number / Leave blank and press enter for final result:")

scanner.Scan()

temp := scanner.Text()

if len(temp) != 0 {
if count == 0 {
count++
fmt.Print("Enter another number:")
fmt.Scanln(&res)
}
x, _ := strconv.ParseFloat(temp, 64)
fmt.Print("Type any one of + / * - operator:")
var opo string
fmt.Scanln(&opo)
switch opo {
case "+":
res += x
fmt.Println(res)
case "-":
res -= x
fmt.Println(res)
case "*":
res *= x
fmt.Println(res)

case "/":
res /= x
fmt.Println(res)
}
} else {
break
}
}
fmt.Println("res=", res)

