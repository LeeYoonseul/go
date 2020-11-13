package main

import "fmt"


func main(){
	array3 := [5]string{"a","b","c","d","e"}
	//var slice3 []string
	slice3 := make([]string,3)
	slice3 = array3[0:3]

	fmt.Println(slice3)
	array3[2] = "Queen"
	fmt.Println(slice3)
	fmt.Printf("%x\n", &slice3[2])
}
