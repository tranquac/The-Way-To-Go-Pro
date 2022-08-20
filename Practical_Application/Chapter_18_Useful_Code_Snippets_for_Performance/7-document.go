package main

import (
	"bufio"
	"fmt"
	"os"
)

func one(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return
	}
	defer file.Close()
	iReader := bufio.NewReader(file)
	for {
		str, err := iReader.ReadString('\n')
		if err != nil {
			return // error or EOF
		}
		fmt.Printf("The input was: %s", str)
	}
}

// func cat(f *file.File) {
// 	const NBUF = 512
// 	var buf [NBUF]byte
// 	for {
// 		switch nr, er := f.Read(buf[:]); true {
// 		case nr < 0:
// 			fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n",
// 				f.String(), er.String())
// 			os.Exit(1)
// 		case nr == 0: // EOF
// 			return
// 		case nr > 0:
// 			if nw, ew := file.Stdout.Write(buf[0:nr]); nw != nr {
// 				fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n",
// 					f.String(), ew.String())
// 			}
// 		}
// 	}
// }

func main() {
	//(1) How to open a file and read:
	one("input.dat")
	//(2) How to read and write files by slicing:
}
