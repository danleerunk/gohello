// Echo3 repeats command-line arguments brah.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for idx, arg := range os.Args {
		fmt.Println(strconv.Itoa(idx) + ": " + arg)
	}
}