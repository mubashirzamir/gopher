package main

import (
	"os"
	"time"

	"github.com/mubashirzamir/gopher/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
