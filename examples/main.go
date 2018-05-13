package main

import (
	"fmt"
	"time"

	"github.com/whaangbuu/go-loading/loading"
)

func main() {
	fmt.Println("Waiting to hatch...")
	loading := loading.StartNew("Pls wait :)")

	loading.SetColor("green")
	// Change the character set
	loading.SetLoaders([]string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠧", "⠇", "⠏"})

	// Other cool loaders
	// loading.SetLoaders([]string{"←", "↖", "↑", "↗", "→", "↘", "↓", "↙"})
	time.Sleep(3 * time.Second)
	loading.Stop()
}
