package main

import "fmt"

func main(){
	var x = 10
	var z = 5
	if x == z{
		fmt.Println("X Sama dengan Z")
	} else if x > z{
		fmt.Println("X Besar dari Z")
	} else {
		fmt.Println("X Kecil dari Z")
	}
	
}