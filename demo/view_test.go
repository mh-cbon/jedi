package main

import (
	"testing"
)

func TestView(t *testing.T) {
	t.Run("can tell it is a view", func(t *testing.T) {
		s := JSampleViewSetup()
		if s.IsView() == false {
			t.Fatal("It must be true")
		}
	})
	t.Run("can tell it is a table", func(t *testing.T) {
		s := JProductSetup()
		if s.IsView() {
			t.Fatal("It must be false")
		}
	})
}
