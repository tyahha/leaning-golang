package main

import (
	"os"

	"github.com/tenntenn/greeting"
)

func main() {
	println(os.Args)
	println(greeting.Do())
}
