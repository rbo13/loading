package loading_test

import (
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

func TestSetColor(t *testing.T) {
	//
}
