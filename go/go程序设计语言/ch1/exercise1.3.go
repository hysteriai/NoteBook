package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var sumString, addStringa = "", "a"
	timeStart := time.Now()
	for i := 0; i < 1000000; i++ {
		sumString += addStringa
	}
	fmt.Printf("%.5fs elapsed\n", time.Since(timeStart).Seconds())

	var stringSlice []string
	var addStringb string = "b"
	timeStartJion := time.Now()
	for i := 0; i < 1000000; i++ {
		strings.Join(stringSlice, addStringb)
	}
	fmt.Printf("%.5fs elapsed\n", time.Since(timeStartJion).Seconds())
}

/*
output
170.65084s elapsed
0.00840s elapsed
*/
