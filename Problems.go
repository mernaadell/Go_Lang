package main

import "fmt"
import "math"

func main() {
    var x int =0
 fmt.Println("1 leap year\n2 calculate sum_differencet\n3 sum_multiples\n4 binary_decimal\n5 plaindorm\n6 prime_factors\n7 calc_seconds\n8 binary_search")
 fmt.Println("Enter the code number ")
 fmt.Scanln(&x) 
 fmt.Println("your choice is ",x)
 
 switch x{
         case 1:
        leap_year()
         break;
        case 2:
        fmt.Println( "square_of_sum - sum_of_squares", sum_difference())
     break;
        case 3:
        fmt.Println(sum_multiples())
     break;
        case 4:
       fmt.Print(binary_decimal())
     break;
        case 5:
        plaindorm()
     break;
        case 6:
        prime_factors()
     break;
        case 7:
        calc_seconds()
     break;
        case 8:
         binary_search()
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
       var str string
  fmt.Println("Enter the string")
  fmt.Scanln(&str)
 //var str string;
// var i int
  var flag int;
 var len2 =len(str)
  	flag = 0;
  	for i:= 0; i < len2; i++{
		if(str[i] != str[len2 - i - 1]){
			flag = 1;
			break;	
		} 
	}
	if(flag == 0){
	fmt.Printf(" %s is a Palindrome String", str);
	}else{
		fmt.Printf("%s is Not a Palindrome String", str);
	}
    
 }
 func prime_factors(){
       var num, isPrime int

    /* Input a number from user */
    fmt.Println("Enter any number to print Prime factors: ");
     fmt.Scanln(&num)

   fmt.Println("All Prime Factors  \n", num)

    /* Find all Prime factors */
    for i:=2; i<=num; i++ {
        /* Check 'i' for factor of num */
        if(num%i==0){
            /* Check 'i' for Prime */
            isPrime = 1
            for j:=2; j<=i/2; j++ {
                if(i%j==0){
                    isPrime = 0
                    break
                }
            }

            /* If 'i' is Prime number and factor of num */
            if(isPrime==1){
                fmt.Printf("%d, ", i)
            }
        }
    }
 }
 func calc_seconds(){
var time int
 fmt.Scanln(&time)
var day = time / (24 * 3600)
time = time % (24 * 3600)
var hour = time / 3600
time %= 3600
var minutes = time / 60
time %= 60
var seconds = time
var years=day/365
day=day-(years*365)
fmt.Println("years:days:hours:min:sec-> ",years,day, hour, minutes, seconds)
 }
 func binary_search(){
      var c, first, last, middle, n, search int
      var array[100]int
 
  fmt.Printf("Enter number of elements\n")
 fmt.Scanln(&n)
 
   fmt.Printf("Enter %d integers\n", n);
 
   for c = 0; c < n; c++{
      
         fmt.Scanln(&array[c])
   }
 
   fmt.Printf("Enter value to find\n")
    fmt.Scanln(&search)
    first = 0
   last = n - 1
  middle = (first+last)/2
 
  for first <= last {
      if (array[middle] < search){
         first = middle + 1
          
      } else if (array[middle] == search) {
         fmt.Printf("%d found at location %d.\n", search, middle+1)
         break
      }else{
         last = middle - 1
      }
      middle = (first + last)/2
   }
   if (first > last){
      fmt.Printf("Not found! %d isn't present in the list.\n", search)}
 }
