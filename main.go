package main
import "fmt";
func main(){
fmt.Printf("I am a Go Developer\n");
//Unary Operator in Go/Golang Programming Language
 x :=10;
 y:=20;
 z:=45;
 x++;
 fmt.Printf("Increment=%v\n",x)
 x--;
 fmt.Printf("Decrement=%v\n",x);
 //Relational Operator
  if(x>y && x>z){
	  fmt.Printf( "%v\n",x);
  }else if(y>x && y>z){
	  fmt.Printf("%v\n",y);
  }else{
     fmt.Printf("%v\n",z);
  }
}