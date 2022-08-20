package main

func main() {
	min, max := MinMax(10, 8)
	println("Min's value is", min)
	println("Max's value is", max)
}

func MinMax(a,b int) (min, max int) {
	if a < b {
		return a,b
	} else {
		return b,a
	}
}