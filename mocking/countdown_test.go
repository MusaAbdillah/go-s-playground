package mocking

import (
	"bytes"
	"fmt"
	"testing"
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

const (
	finalWord      = "Go!"
	countDownStart = 3
)

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

	fmt.Println(spySleeper.Calls)
	if spySleeper.Calls != 4 {
		t.Errorf("Not enough calls to sleeper, want 4 got %d", spySleeper.Calls)
	}
}

func Countdown(out *bytes.Buffer, sleeper Sleeper) {
	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
