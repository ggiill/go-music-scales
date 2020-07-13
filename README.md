# go-music-scales

Go library that will provide the notes of a musical scale.

Has a (very basic) CLI component.

## Library: `scales`

```go
package main

import (
  scales "github.com/gturetsky/go-music-scales"
)

func main() {
	s := scales.NewScale("Eb", "Minor")
	s.Identify()
}
```

## CLI
*(Will flesh this out a bit more in the future...)*
```
$ go run cmd/go-music-scales/main.go Eb Minor
I am an Eb Minor scale and my notes are Eb, F, Gb, Ab, Bb, Cb, Db
```
