package scales_test

import (
	"testing"

	scales "github.com/gturetsky/go-music-scales"
)

func TestAllNotes(t *testing.T) {
	cof := scales.CircleOfFifths()
	modes := scales.CreateModes()
	for _, root := range cof {
		for mode, _ := range modes {
			scale, err := scales.NewScale(root, mode)
			if err != nil {
				t.Error(err)
			}
			notes := scale.GetNotes()
			t.Logf("%s %s: %s\n", root, mode, notes)
			if len(notes) != 7 {
				t.Errorf("Error: %s %s scale did not return 7 notes", root, mode)
			}
		}
	}
}
