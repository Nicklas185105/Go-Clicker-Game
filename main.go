package main

import (
	"clicker-game/frontend"
	"os"

	"github.com/dtylman/gowd"

	"fmt"
)

func main() {
	err := gowd.Run(frontend.NewApplication().Element)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
