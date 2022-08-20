package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)                                            //2022-08-06 20:10:20.1423497 +0700 +07 m=+0.002105201
	fmt.Printf("%02d.%02d.%4d", t.Day(), t.Month(), t.Year()) //06.08.2022
	t = time.Now().UTC()
	fmt.Println(t) //06.08.20222022-08-06 13:11:12.717435 +0000 UTC

	//calculate time
	//  var week = 60 * 60 * 24 * 7 * 1e9 //must be nanosecond
	// week_from_now := t.Add(week)
	// fmt.Println(week_from_now)
	//format time
	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04:05"))
	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
