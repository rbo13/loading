package loading_test

import (
	"testing"

	"github.com/whaangbuu/go-loading/loading"
)

func TestStart(t *testing.T) {
	// arrange
	l := loading.NewLoading("Starting...")
	// act
	got := l.Start()
	want := l

	// assert
	if want != got {
		t.Errorf("Want '%v', but Got '%v'", want, got)
	}
	t.Log(got)
}

func TestIsColorAllowed(t *testing.T) {
	// arrange
	allowedColor := "blue"

	// act
	got := loading.IsColorAllowed(allowedColor)
	want := true

	// assert
	if want != got {
		t.Errorf("Want '%t', but Got '%t'", want, got)
	}

	t.Log(got)
}
