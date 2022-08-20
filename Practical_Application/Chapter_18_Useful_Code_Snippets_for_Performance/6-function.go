package main

import "log"

func protect(g func()) {
	defer func() {
		log.Println("done")
		// Println executes normally even if there is a panic
		if x := recover(); x != nil {
			log.Printf("run time panic: %v", x)
		}
	}()
	log.Println("start")
	g()
}

func main() {
	//How to use the built-in function recover()to terminate panic()procedure (refer to Section 13.3 ):
}
