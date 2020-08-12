// Package scales provides functionality to get:
//
// - The notes of a scale given a root and mode
//
// - The scales corresponding to a given set of notes
package scales

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

const (
	// NoteLetters are the valid letters for notes
	NoteLetters string = "ABCDEFG"
	// StepHalf is an enum for a half step
	StepHalf int = 1
	// StepWhole is an enum for a whole step
	StepWhole int = 2
)

var circleOfFifths = []string{
	"C", "B#", // enharmonic
	"G",
	"D",
	"A",
	"E", "Fb", // enharmonic
	"B", "Cb", // enharmonic
	"F#", "Gb", // enharmonic
	"C#", "Db", // enharmonic
	"Ab", "G#", // enharmonic
	"Eb", "D#", // enharmonic
	"Bb", "A#", // enharmonic
	"F", "E#"} // enharmonic

// CircleOfFifths is a constructor for the slice representing the
// circle of fifths.
func CircleOfFifths() []string {
	return circleOfFifths
}

// Modes is a constructor for the main modes and additional scales.
func Modes() map[string][]int {
	scaleModes := []string{"Ionian", "Dorian", "Phrygian", "Lydian", "Myxolydian", "Aeolian", "Locrian"}
	modes := map[string][]int{}
	modeSteps := []int{StepWhole, StepWhole, StepHalf, StepWhole, StepWhole, StepWhole, StepHalf}
	for modeIndex, modeName := range scaleModes {
		modes[modeName] = append(modeSteps[modeIndex:], modeSteps[:modeIndex]...)
	}
	// Include Major & Minor (which are aliases of two of the modes above):
	modes["Major"] = modes["Ionian"]
	modes["Minor"] = modes["Aeolian"]
	// Weird scales:
	// TODO: Add more
	modes["Harmonic Minor"] = []int{StepWhole, StepHalf, StepWhole, StepWhole, StepHalf, StepWhole + StepHalf, StepHalf}
	modes["Melodic Minor"] = []int{StepWhole, StepHalf, StepWhole, StepWhole, StepWhole, StepWhole, StepHalf}
	modes["Phrygian Dominant"] = []int{StepHalf, StepWhole + StepHalf, StepHalf, StepWhole, StepHalf, StepWhole, StepWhole}
	return modes
}

var notes = map[int][]string{
	// Notes and their funky enharmonic equivalents
	0:  []string{"B#", "C", "Dbb"},
	1:  []string{"C#", "Db"},
	2:  []string{"C##", "D", "Ebb"},
	3:  []string{"D#", "Eb"},
	4:  []string{"D##", "E", "Fb"},
	5:  []string{"E#", "F", "Gbb"},
	6:  []string{"E##", "F#", "Gb"},
	7:  []string{"F##", "G", "Abb"},
	8:  []string{"G#", "Ab"},
	9:  []string{"G##", "A", "Bbb"},
	10: []string{"A#", "Bb", "Cbb"},
	11: []string{"A##", "B", "Cb"},
}

// Notes is a constructor for the 12 notes and their enharmonic equivalents.
func Notes() map[int][]string {
	return notes
}

// Scale is a struct representing a musical scale. Scale objects should be
// initialized using the NewScale constructor.
type Scale struct {
	Root  string
	Mode  string
	Modes map[string][]int
	Notes map[int][]string
}

// Identify prints the name of the scale and its notes.
func (s *Scale) Identify() {
	var qualifier string
	switch string(s.Root[0]) {
	case "A", "E":
		qualifier = "an"
	default:
		qualifier = "a"
	}
	fmt.Printf("I am %s %s %s scale and my notes are %s\n",
		qualifier, s.Root, s.Mode, strings.Join(s.GetNotes(), ", "))
}

func (s *Scale) getNoteOptions() [][]string {
	var startingNoteIndex int
	// Find the []string with the starting note of the scale:
	for noteIndex, noteNotes := range s.Notes {
		for _, noteAlias := range noteNotes {
			if noteAlias == s.Root {
				// We found it!
				startingNoteIndex = noteIndex
			}
		}
	}
	noteList := [][]string{s.Notes[startingNoteIndex]}
	noteOffset := startingNoteIndex
	modeSteps := s.Modes[s.Mode]
	for index, offset := range modeSteps {
		if index < 6 {
			noteOffset += offset
			noteOffset = noteOffset % 12 // Mod for values >= 12 (12 tones)
			noteList = append(noteList, s.Notes[noteOffset])
		}
	}
	return noteList
}

// GetNotes returns a slice of the Scale's notes (based on the Scale's root and mode).
func (s *Scale) GetNotes() []string {
	// TODO: Add support for scales with !=8 notes (bebop, pentatonic, etc.)
	rootNoteLetter := string(s.Root[0])
	rnli := strings.Index(NoteLetters, rootNoteLetter) // rnli = root note letter index
	// Build the scale (without accidentals):
	scaleLetters := strings.Split(NoteLetters[rnli:]+NoteLetters[:rnli], "")
	scale := []string{}
	noteOptions := s.getNoteOptions()
	// Iterate through the noteOptions and find the appropriate note (option):
	for letterIndex, letterValue := range scaleLetters {
		noteOptionLetters := noteOptions[letterIndex]
		for _, noValue := range noteOptionLetters {
			if strings.HasPrefix(noValue, letterValue) {
				scale = append(scale, noValue)
			}
		}
	}
	scale = append(scale, scale[0]) // Add the octave
	return scale
}

// NewScale is a constructor function to create a Scale object with defaults.
func NewScale(root string, mode string) (*Scale, error) {
	isInCOF := false
	cof := CircleOfFifths()
	for _, c := range cof {
		if root == c {
			isInCOF = true
			break
		}
	}
	if !isInCOF {
		return nil, fmt.Errorf("Error: Root note %s not in Circle of Fifths", root)
	}
	s := new(Scale)
	s.Root = root
	s.Mode = mode
	s.Modes = Modes()
	isLegitMode := false
	for modesMode := range s.Modes {
		if modesMode == mode {
			isLegitMode = true
			break
		}
	}
	if !isLegitMode {
		return nil, fmt.Errorf("Error: Mode %s is not a supported mode", mode)
	}
	s.Notes = Notes()
	return s, nil
}

// VerifyScaleLetters checks the input notes to make sure they correspond
// to ABCDEFG.
func VerifyScaleLetters(notes []string) bool {
	notes = DedupeNotes(notes)
	sort.Strings(notes)
	for noteIndex, noteValue := range notes {
		noteValueLetter := noteValue[0]
		if noteValueLetter != NoteLetters[noteIndex] {
			return false
		}
	}
	return true
}

// ScaleNoteMatch checks the input slice of notes against a given Scale to see if they match.
func ScaleNoteMatch(notes []string, root string, mode string) (bool, error) {
	if !VerifyScaleLetters(notes) {
		return false, fmt.Errorf("Error: %s: Not a valid scale", notes)
	}
	scale, err := NewScale(root, mode)
	if err != nil {
		return false, fmt.Errorf("Error: %s %s: Not a valid root and mode", root, mode)
	}
LOOP:
	for _, scaleNote := range scale.GetNotes() {
		for _, note := range notes {
			if scaleNote == note {
				continue LOOP
			}
		}
		return false, nil
	}
	return true, nil
}

// GetScalesFromNotes iterates over all modes to find the modes that match the input notes.
func GetScalesFromNotes(notes []string) ([]string, error) {
	if !VerifyScaleLetters(notes) {
		return nil, fmt.Errorf("Error: %s: Not a valid set of scale notes", notes)
	}
	if len(notes) < 7 {
		return nil, fmt.Errorf("Error: %s: Need at least 7 notes", notes)
	}
	var wg sync.WaitGroup
	ch := make(chan string, 1)
	scalesFromNotes := []string{}
	cof := CircleOfFifths()
	modes := Modes()
	wg.Add(len(cof) * len(modes))
	for _, root := range cof {
		for mode := range modes {
			go func(notes []string, root string, mode string) {
				scaleNotes, err := ScaleNoteMatch(notes, root, mode)
				if err != nil {
					wg.Done()
					return
				}
				if scaleNotes {
					scaleName := root + " " + mode
					ch <- scaleName
					return
				}
				wg.Done()
				return
			}(notes, root, mode)
		}
	}
	go func() {
		for val := range ch {
			scalesFromNotes = append(scalesFromNotes, val)
			wg.Done()
		}
	}()
	wg.Wait()
	sort.Strings(scalesFromNotes)
	return scalesFromNotes, nil
}

// DedupeNotes dedupes notes (dedupes repeated notes and octaves).
func DedupeNotes(notes []string) []string {
	deduped := []string{}
	dedupedMap := make(map[string]int)
	for _, val := range notes {
		dedupedMap[val]++
	}
	for val := range dedupedMap {
		deduped = append(deduped, val)
	}
	return deduped
}
