// basic_concurrency
package main

import (
	"fmt"
	"time"
)

// you need to put main fuction on sleep so go routines can run otherwise it will exit automatically
func main() {

	godur, _ := time.ParseDuration("10ms")

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Hello")
			//making it sleep gives a chance to run other go routine so the switching betwwen go routine is possible (see the output)
			time.Sleep(godur)
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Go")
			time.Sleep(godur)
		}
	}()

	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
}
