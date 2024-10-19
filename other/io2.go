package main

import (
	"fmt"
	"strings"
)

func main() {
	r := strings.NewReader("I dont know what")
	buf := make([]byte, 4)
	for {
		n, err := r.Read(buf)
		fmt.Println(buf[:n], err)
		//reads 4 characters at a time in each iteration
		if err != nil {
			fmt.Println("break")
			break
		}
	}
}
