// Echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg // this is not efficient because each time around the loop, a new string is initialized
		sep = " "
	}
	fmt.Println(s)
}
