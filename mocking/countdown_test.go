package main


import (
	"testing"
	"bytes"
	"fmt"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int 
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

const finalWord = "Go!"
const countDownStart = 3

func TestCountdown(t *testing.T) {
    buffer := &bytes.Buffer{}
    spySleeper := &SpySleeper{}

    Countdown(buffer, spySleeper)

    got := buffer.String()
    want := `3
    2
    1
    Go!`

    if got != want {
        t.Errorf("got %q want %q", got, want)
    }

    if spySleeper.Calls != 4 {
    	t.Errorf("Not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
    }
}

func Countdown(out *bytes.Buffer, sleeper Sleeper){
		for i := countdownStart; i>0; i-- {
			sleeper.Sleep()
			fmt.Fprintln(out, i)
		}


		sleeper.Sleep()
	    fmt.Fprint(out, finalWord)
}