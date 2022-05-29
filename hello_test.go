package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("hello with param", func(t *testing.T) {
		want := "Hello, Jelte!"
		have := Hello("Jelte")
        
        assertEqual(want, have, t)
	})
    t.Run("should be world", func(t *testing.T) {
        want := "Hello, world!"
        have := Hello("")

        assertEqual(want, have, t)
    })
}

func assertEqual(want string, have string, t testing.TB) {
    t.Helper()
    if want != have {
        t.Errorf("want=%v, have=%v\n", want, have)
    }
}

