package scales_test

import (
	"testing"

	scales "github.com/ggiill/go-music-scales"
)

func TestAllNotes(t *testing.T) {
	cof := scales.CircleOfFifths()
	modes := scales.Modes()
	for _, root := range cof {
		for mode, _ := range modes {
			scale, err := scales.NewScale(root, mode)
			if err != nil {
				t.Error(err)
			}
			notes := scale.GetNotes()
			t.Logf("%s %s: %s\n", root, mode, notes)
			if len(notes) != 8 {
				t.Errorf("Error: %s %s scale did not return 8 notes", root, mode)
			}
		}
	}
}
