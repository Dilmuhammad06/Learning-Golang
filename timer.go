package main

import (
	"fmt"
	"time"
)


func main(){

var cc int;
fmt.Println("How many seconds you want to count back: ");
fmt.Scanln(&cc);

for{
    if cc > 0{
        cc--;
        time.Sleep(time.Second);
        fmt.Println(cc, "seconds left");
 }  else {
        fmt.Println("It's time bro");
        break;
     }

   }
}
