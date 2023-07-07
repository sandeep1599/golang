package main

import (
	"fmt"


)

func Split(sum int) ( x, y int){
 
x = sum - 3
y = sum - x
return
}



/*
func add(x int , y int) int {
  return x+y
}


func swap(x ,y string) (string,string){
  return y,x
}

*/


func main (){
/*  a:= 65
fmt.Println (a)

 b, c, d, _, f := go run main.go

fmt.Println(b,c,d,f)


b,c,d,e,f := 33,87,65,77,44

fmt.Println(b,c,d,f)


var g int 
fmt.Println(g)
g = 11
fmt.Println(g)

  var k  int= 98

fmt.Printf("%T",k)
fmt.Println(k)


adams := 54

fmt.Printf( "The %b is a binary of 54 \n", adams)
fmt.Printf("The %x is a hexadecimal of 54 n\n" , adams)

a,b,c,d,e,f := 0,1,2,3,4,5


fmt.Printf("%v \t %b \t %#x\n",a,a,a)
fmt.Printf("%v \t %b \t %#x\n",b,b,b)
fmt.Printf("%v \t %b \t %#x\n",c,c,c)
fmt.Printf("%v \t %b \t %#x\n",d,d,d)
fmt.Printf("%v \t %b \t %#x\n",e,e,e)
fmt.Printf("%v \t %b \t %#x\n",f,f,f)



 x  := 42
 v := 42.0

 fmt.Printf("%v of type %T \n", x,x)
 fmt.Printf("%v of type %T \n", v,v)


 var m float32 =43.7

 v = float64(m)
  fmt.Printf( "%v of type %T\n",v,v)

  

fmt.Println("Welcome to the playground")
fmt.Println("The time is", time.Now())



fmt.Println(add (6532767 , 8763462786))

x, y :=swap( "Hello!" , "Sandeep")
fmt.Println(x, y)

*/

fmt.Println(Split(20))

}