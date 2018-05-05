package loading_test

import (
	"errors"
	"testing"

	"github.com/whaangbuu/go-loading/loading"
)

func TestStart(t *testing.T) {
	l := loading.NewLoading("Starting...")
	got := l.Start()
	want := l

	if want != got {
		t.Errorf("Want '%v', but Got '%v'", want, got)
	}
}

func TestIsColorAllowed(t *testing.T) {
	allowedColor := "blue"

	got := loading.IsColorAllowed(allowedColor)
	want := true

	if want != got {
		t.Errorf("Want '%t', but Got '%t'", want, got)
	}

}

func TestStartNew(t *testing.T) {
	got := loading.StartNew("Starting new")

	if got == nil {
		t.Errorf("NIL is not allowed")
	}
}

func TestSetColorErr(t *testing.T) {
	l := loading.NewLoading("Setting color")
	want := errors.New("color not found")
	got := l.SetColor("black")

	if want == got {
		t.Errorf("Want '%v', but Got'%v'", want, got)
	}
}

func TestSetColor(t *testing.T) {
	loader := loading.NewLoading("Setting valid color")

	colorErr := loader.SetColor("blue")

	if colorErr != nil {
		t.Errorf("ERROR DUE TO: %v", colorErr)
	}
}
