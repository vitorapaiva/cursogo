package exercicio

import "fmt"

var x int
var y string
var z bool

var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true

func Exercicio1() {
  fmt.Println("Exercicio 1")
  x:=42
  y:="James Bond"
  z:=true
  fmt.Println(x, y, z)
  fmt.Println(x)
  fmt.Println(y)
  fmt.Println(z)
}

func Exercicio2() {
  fmt.Println("Exercicio 2")
  fmt.Println(x, y, z)
  //valor zero
}

func Exercicio3() {
  fmt.Println("Exercicio 3")
  s := fmt.Sprintf("%v, %v, %v", x3, y3, z3)
  fmt.Println(s)
}

func Exercicio4() {
  fmt.Println("Exercicio 4")
  
}

func Exercicio5() {
  fmt.Println("Exercicio 5")
}

func Exercicio6() {
  fmt.Println("Exercicio 6")
}