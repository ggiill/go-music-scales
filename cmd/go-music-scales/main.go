package main

import (
	"fmt"
	"os"

	scales "github.com/gturetsky/go-music-scales"
)

func main() {
	args := os.Args
	s, err := scales.NewScale(args[1], args[2])
	if err != nil {
		fmt.Println(err)
	} else {
		s.Identify()
	}
}
