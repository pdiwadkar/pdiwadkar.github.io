package main

import (
	"fmt"
	"time"
)
func main(){
 cn := make(chan int)
 go EvenNumberGenerators(cn)
 for {
	fmt.Println(<-cn)
 }

}
func EvenNumberGenerators(cn chan int){
	ctr := 2
	for {
		cn <- ctr
		ctr +=2
		time.Sleep(time.Second*1)
	}	
}