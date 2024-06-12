package main

import (
	"fmt"
	"os"
)

// debug
func log(msg ...interface{}) {
	fmt.Fprintln(os.Stderr, msg...)

}
