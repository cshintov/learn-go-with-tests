package main

import (
	"mocking/internal"
	"os"
)


func main() {
    sleeper := &internal.DefaultSleeper{}
	internal.Countdown(os.Stdout, sleeper)
}
