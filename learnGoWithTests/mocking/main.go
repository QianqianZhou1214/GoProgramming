package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

//for main

const finalWord = "Go!"
const countdownStart = 3

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}

type Sleeper interface {
	Sleep()
}

//Spies are a kind of mock which can record how a dependency is used.

func Countdown(out io.Writer, sleeper Sleeper) {
	//Countdown receives an output target.
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, finalWord)
}
