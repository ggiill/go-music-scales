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

Provides options for:

**Root Note and Mode -> Notes of scale**
```console
$ go run cmd/go-music-scales/main.go scale -root A -mode Minor
I am an A Minor scale and my notes are A, B, C, D, E, F, G, A
```

**Notes of scale -> All scales (Root Notes and Modes)**
```console
$ go run cmd/go-music-scales/main.go notes --list A B C D E F G
These scales are comprised of the notes [A B C D E F G]: A Aeolian, A Minor, B Locrian, C Ionian, C Major, D Dorian, E Phrygian, F Lydian, G Myxolydian
```
