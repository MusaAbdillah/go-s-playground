package mocking

import (
	"bytes"
	"fmt"
	"testing"
)

// Struct
type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

type SpyCountdownOperations struct {
	Calls []string
}

// Singleton function
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Constant
const (
	finalWord      = "Go!"
	countDownStart = 3
	write          = "write"
	sleep          = "sleep"
)

// Main Testing function
func TestCountdown(t *testing.T) {
	t.Run("Prints 3 to Go! ", func(t *testing.T) {
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
	})

	// t.Run("Sleep before every print!", func(t *testing.T) {
	// 	spySleepPrinter := &SpyCountdownOperations{}
	// 	// Countdown(spySleepPrinter, spySleepPrinter)

	// 	want := []string{
	// 		sleep,
	// 		write,
	// 		sleep,
	// 		write,
	// 		sleep,
	// 		write,
	// 		sleep,
	// 		write,
	// 	}

	// 	if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
	// 		t.Errorf("wanted calls %v got %v", want, spySleepPrinter)
	// 	}
	// })
}

func Countdown(out *bytes.Buffer, sleeper Sleeper) {

	for i := countDownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
