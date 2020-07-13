package main

import (
	"os"

	scales "github.com/gturetsky/go-music-scales"
)

func main() {
	args := os.Args
	s := scales.NewScale(args[1], args[2])
	s.Identify()
}
