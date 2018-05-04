package loading

import (
	"errors"
	"fmt"
	"io"
	"sync"
	"syscall"
	"time"

	"github.com/fatih/color"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	// 100ms per frame
	DEFAULT_FRAME_RATE = time.Millisecond * 100
)

var errColorNotFound = errors.New("color not found")

var Loaders = []string{"⣾", "⣽", "⣻", "⢿", "⡿", "⣟", "⣯", "⣷"}

// allowableColors is a collection of allowed colors.
var allowableColors = map[string]bool{
	"red":   true,
	"gree":  true,
	"blue":  true,
	"white": true, // The default color
}

var foregroundColorAttribute = map[string]color.Attribute{
	"red":   color.FgRed,
	"green": color.FgGreen,
	"blue":  color.FgBlue,
	"white": color.FgWhite,
}

// IsColorAllowed returns boolean if a color is allowed.
func IsColorAllowed(color string) bool {
	allowed := false
	if allowableColors[color] {
		allowed = true
	}
	return allowed
}

type Loading struct {
	sync.Mutex
	Title     string
	Charset   []string
	FrameRate time.Duration
	runChan   chan struct{}
	stopOnce  sync.Once
	Output    io.Writer
	NoTty     bool
	Color     func(a ...interface{}) string
}

func NewLoading(title string) *Loading {
	loading := &Loading{
		Title:     title,
		Charset:   Loaders,
		FrameRate: DEFAULT_FRAME_RATE,
		runChan:   make(chan struct{}),
		Color:     color.New(color.FgWhite).SprintFunc(),
	}
	if !terminal.IsTerminal(syscall.Stdout) {
		loading.NoTty = true
	}
	return loading
}

func StartNew(title string) *Loading {
	return NewLoading(title).Start()
}

// start loading
func (loading *Loading) Start() *Loading {
	go loading.writer()
	return loading
}

func (loading *Loading) SetColor(c string) error {
	if !IsColorAllowed(c) {
		return errColorNotFound
	}
	loading.Color = color.New(foregroundColorAttribute[c]).SprintFunc()
	//loading.Restart()
	return nil
}

// loading framerate speed
func (loading *Loading) SetSpeed(rate time.Duration) *Loading {
	loading.Lock()
	loading.FrameRate = rate
	loading.Unlock()
	return loading
}

// SetLoaders set the character loader of the loading
func (loading *Loading) SetLoaders(chars []string) *Loading {
	loading.Lock()
	loading.Charset = chars
	loading.Unlock()
	return loading
}

// Stop stops and clears the loading
func (loading *Loading) Stop() {
	loading.stopOnce.Do(func() {
		close(loading.runChan)
		loading.clear()
	})
}

func (loading *Loading) Restart() {
	loading.Stop()
	loading.Start()
}

// animates our loader
func (loading *Loading) animate() {
	var out string
	for i := 0; i < len(loading.Charset); i++ {
		out = loading.Color(loading.Charset[i]) + " " + loading.Title
		switch {
		case loading.Output != nil:
			fmt.Fprint(loading.Output, out)
		case !loading.NoTty:
			fmt.Print(out)
		}
		time.Sleep(loading.FrameRate)
		loading.clear()
	}
}

func (loading *Loading) writer() {
	loading.animate()
	for {
		select {
		case <-loading.runChan:
			return
		default:
			loading.animate()
		}
	}
}

func (loading *Loading) clear() {
	if !loading.NoTty {
		fmt.Printf("\033[2K")
		fmt.Println()
		fmt.Printf("\033[1A")
	}
}
