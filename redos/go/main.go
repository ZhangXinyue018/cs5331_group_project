package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	Normal()
	Attack()
}

func Normal() {
	fmt.Println("----------- start normal regex -----------")
	time1 := time.Now()
	regexp.MatchString("(a+)+$", "aaa")
	time2 := time.Now()
	diff := time2.Sub(time1).Nanoseconds()
	fmt.Printf("it takes %d nanoseconds for normal regex \n", diff)
}

func Attack() {
	fmt.Println("----------- start attack regex -----------")
	time1 := time.Now()
	regexp.MatchString("(a+)+$", "aaaaaaaaaaaaaaaaaaaaax")
	time2 := time.Now()
	diff := time2.Sub(time1).Nanoseconds()
	fmt.Printf("it takes %d nanoseconds for attack regex in \n", diff)
}
