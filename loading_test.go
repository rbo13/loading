package loading_test

import (
	"errors"
	"testing"
	"time"

	"github.com/whaangbuu/go-loading/loading"
)

func TestSpin(t *testing.T) {
	title := "Test spin"
	t.Run(title, func(t *testing.T) {
		showLoading(title)
	})

}

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

func TestSetSpeed(t *testing.T) {
	loader := loading.NewLoading("Setting speed")

	speedErr := loader.SetSpeed(200 * time.Millisecond)

	if speedErr == nil {
		t.Errorf("Loading instance is nil, due to: %v", speedErr)
	}

}

func TestSetLoader(t *testing.T) {
	loader := loading.NewLoading("Set Loaders")

	loaders := loader.SetLoaders([]string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"})

	if loaders == nil {
		t.Errorf("Loaders not set due to: %v", loaders)
	}
}

func showLoading(title string) {
	loader := loading.StartNew(title)
	defer loader.Stop()
	time.Sleep(2 * time.Second)
}
