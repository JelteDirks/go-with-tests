package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("hello with param", func(t *testing.T) {
		want := "Hello, Jelte!"
		have := Hello("Jelte")

		if want != have {
			t.Errorf("want=%v, have=%v\n", want, have)
		}
	})
    t.Run("should be world", func(t *testing.T) {
        want := "Hello, world!"
        have := Hello("")

        if want != have {
            report(want, have, t)
        }
    })
}

func report(want string, have string, t *testing.T) {
    t.Errorf("want=%v, have=%v\n", want, have)
}

