package main

import (
	"mocking/internal"
	"os"
)


func main() {
	internal.Countdown(os.Stdout)
}
