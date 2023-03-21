package internal

import (
	"fmt"
	"io"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}

	fmt.Fprint(out, finalWord)
}

// func Countdown(out io.Writer, sleeper Sleeper) {
// 	for i := countdownStart; i > 0; i-- {
// 		sleeper.Sleep()
// 	}

// 	for i := countdownStart; i > 0; i-- {
// 		fmt.Fprintln(out, i)
// 	}

// 	fmt.Fprint(out, finalWord)
// }

// The count and the sleep is part of the behaviour of the Countdown.
// Countdown counts down from count and sleeps in between
// func Countdown(count int, sleeper Sleeper) string {
// 	out := ""
// 	for ; count > 0; count-- {
// 		out = addString(out, count)
// 		sleeper.Sleep()
// 	}

// 	return out + finalWord
// }
