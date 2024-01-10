package main

import (
	"fmt"
)

func Heddy() {
	var a int = 1
	c := "Hello world"
	var d = 3
	fmt.Println(c)
	fmt.Println(d)
	//-------------
	if a != 1 {
		fmt.Println("Germain est un Oméga BAKA!!!!!!")
	} else {
		fmt.Println("Bastien est comme Germain")
	}

	for {
		if a != 10 {
			fmt.Println("Not good value")
			a = a + 1
		} else {
			fmt.Println("Good Value")
			break
		}
	}
}
