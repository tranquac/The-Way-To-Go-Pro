package main

var a = "Q"

func main() {
	n()
	m()
	n()
}

func n() {
	print(a)
}

func m() {
	a = "O"
	print(a)
}