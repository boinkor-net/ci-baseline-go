package main

import "log"

func main() {
	s := []string{"a", "b", "c"}
	a := [3]string(s)
	log.Printf("hi %v", a)
}
