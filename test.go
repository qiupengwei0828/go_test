package main

import (
   "fmt"
)
/****
  *go  语言测试实践
  */
func main() {
   var num int =1;
   fmt.Println(num);
   var num1,num2 int=1,2;
   fmt.Println(num1);
   fmt.Println(num2);
   var(
      num3 int =4
      num4 int =5
    )
   fmt.Println(num3);
   fmt.Println(num4);

   var  str1 string ="abcd";
   fmt.Println(str1)

   var numbers=[3]int{1,2,3}
   fmt.Println(numbers)
   var length=len(numbers)
   fmt.Println(length)

}