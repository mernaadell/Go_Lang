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
package main

import "fmt"
import "math"

func main() {
    var x int =0
 fmt.Println("1 leap year\n 2 calculate the moment\n")
 fmt.Println("Enter the code number ")
 fmt.Scanln(&x) 
 fmt.Println("your choice is ",x)
 
 switch x{
         case 1:
        leap_year()
         break;
        case 2:
        fmt.Println(  sum_difference())
     break;
        case 3:
        fmt.Println(sum_multiples())
     break;
        case 4:
       fmt.Print(binary_decimal())
     break;
        case 5:
     break;
        case 6:
     break;
        case 7:
     break;
        case 8:
     break;
     
 }
 }
 func leap_year(){
     var year int
     fmt.Println("Enter the year")
     fmt.Scanln(&year)
     if(year%4==0){
         if(year%100==0){
             if(year%400==0){
                 fmt.Println(year,"is a Leap Year")
             } else{
                  fmt.Println(year,"isn't a Leap Year")
             }
         } else{
              fmt.Println(year,"is a Leap Year")
         }
     }else{
          fmt.Println(year,"isn't a Leap Year")
     }

}
 func sum_difference() int {
       var n int
     fmt.Println("Enter the Number")
     fmt.Scanln(&n)
   var sum_of_squares int = 0
   var square_of_sum int = 0
   for i:=1;i<=n;i++ {
        sum_of_squares +=i * i
        square_of_sum += i
   } 
      square_of_sum = int(math.Pow(float64(square_of_sum), 2)) 
    return square_of_sum - sum_of_squares
   
     
 }
 func sum_multiples() int{
       var n int
       var a int
     fmt.Println("Enter the Number")
     fmt.Scanln(&a)
     fmt.Println("to Range")
     fmt.Scanln(&n)
    //  var m=n/a //number of multiplies
    //  var sum = m * (m + 1) / 2; 
    //  var ans =a*sum
    var i=2
    var sum=0
    var c=0
    for c<n {
        sum+=i*a
        c=i*a
        i++
       // fmt.Println(sum)
    }
     return sum
 }
 func binary_decimal()int { 
     var n int
  fmt.Println("Enter the Number in binary")
  fmt.Scanln(&n)
   var num = n 
    var dec_value = 0
    var base = 1 
  
   var temp = num 
     for temp>0 { 
        var last_digit = temp % 10
        temp = temp / 10
  
        dec_value += last_digit * base
  
  
        base = base * 2
      //  fmt.Println(dec_value)
        
    } 
    return dec_value
     
 }
 func plaindorm(){
       var str ="mena"
  fmt.Println("Enter the string")
  //fmt.Scanln(&str)
 //var str string;
  var i int
  var flag int;
//  var len =len(str)
  	flag = 0;
  	for i:= 0; i < 5; i++{
		if(str[i] != str[5 - i - 1]){
			flag = 1;
			break;	
		} 
	}
	if(flag == 0){
	fmt.Printf("\n %s is a Palindrome String", str);
	}else{
		fmt.Printf("\n %s is Not a Palindrome String", str);
	}
    
 }
