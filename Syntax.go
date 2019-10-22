package main

import "fmt"

func main() {
  var x float64
   x = 20.0
  var rate int =3
  fmt.Println("rate",rate)
   y := 42 
   fmt.Println(x)
   fmt.Println(y)
   fmt.Printf("x is of type %T\n", x)
   fmt.Println("hi")
   fmt.Printf("y is of type %T \n", y)
   fmt.Println("__________")
 var a, b, c = 3, 4, "foo"  //da satr 3la b3do
	
   fmt.Println(a)
   fmt.Println(b)
   fmt.Println(c)
   fmt.Printf("a is of type %T\n", a)
   fmt.Printf("b is of type %T\n", b)
   fmt.Printf("c is of type %T\n", c)
 var x1=10 
 var y1=20
  var f=x1+y1
  fmt.Println(f)
  fmt.Println("________________")
 if (x1>3) {
      fmt.Println("Yes")
  }
     num := 10
    if num % 2 == 0 { //checks if number is even
        fmt.Println("the number is even") 
    }  else {
        fmt.Println("the number is odd")
    }
  for i:=0;i<10;i++{
      fmt.Println(i)
  }
  switch rate{
      
     case 0:
     fmt.Println("good")
     break
     case 1:
     fmt.Println("very good")
     break
       case 2:
     fmt.Println("very good")
     break
       case 3:
     fmt.Println("very good")
      break
      case 4:
     fmt.Println("very good")
     default:
     fmt.Println("bye")
  }
  fmt.Println("______functions")
  mass()
    nam("merna")
   fmt.Println( add(2,3))
   fmt.Println("-----------arrays-----")
   var g [10] int
   g[0]=1
   g[2]=2
 }
 func mass(){
     
     fmt.Println("fuck you")
 }
 func nam (name string){
     
     fmt.Println(name)
 }
  func add (x int ,y int) int{
     
     return x+y
 }
//output
/*
rate 3
20
42
x is of type float64
hi
y is of type int 
__________
3
4
foo
a is of type int
b is of type int
c is of type string
30
________________
Yes
the number is even
0
1
2
3
4
5
6
7
8
9
very good
______functions
fuck you
merna
5
-----------arrays-----*/

