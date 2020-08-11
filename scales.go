package scales

import (
	"fmt"
	"strings"
)

const (
	NoteLetters   string = "ABCDEFG"
	NoteHalf      int    = 1
	NoteWhole     int    = 2
	NoteWholeHalf int    = 3
)

var circleOfFifths = []string{"C", "G", "D", "A", "B", "Cb", "F#", "Gb", "C#", "Db", "Ab", "Eb", "Bb", "F"}

func CircleOfFifths() []string {
	return circleOfFifths
}

var scaleModes = []string{"Ionian", "Dorian", "Phrygian", "Lydian", "Myxolydian", "Aeolian", "Locrian"}

func ScaleModes() []string {
	return scaleModes
}

func CreateModes() map[string][]int {
	scaleModes := ScaleModes()
	modes := map[string][]int{}
	modeSteps := []int{NoteWhole, NoteWhole, NoteHalf, NoteWhole, NoteWhole, NoteWhole, NoteHalf}
	for modeIndex, modeName := range scaleModes {
		modes[modeName] = append(modeSteps[modeIndex:], modeSteps[:modeIndex]...)
	}
	// Gotta include Major & Minor (which are aliases of the modes above):
	modes["Major"] = modes["Ionian"]
	modes["Minor"] = modes["Aeolian"]
	// Weird scales:
	modes["Harmonic Minor"] = []int{NoteWhole, NoteHalf, NoteWhole, NoteWhole, NoteHalf, NoteWholeHalf, NoteHalf}
	return modes
}

type Scale struct {
	Root  string
	Mode  string
	Modes map[string][]int
	Notes map[int][]string
}

func (s *Scale) Identify() {
	var qualifier string
	switch string(s.Root[0]) {
	case "A", "E":
		qualifier = "an"
	default:
		qualifier = "a"
	}
	fmt.Printf("I am %s %s %s scale and my notes are %s\n", qualifier, s.Root, s.Mode, strings.Join(s.GetNotes(), ", "))
}

func (s *Scale) getNoteOptions() [][]string {
	var startingNoteIndex int
	for noteIndex, noteNotes := range s.Notes {
		for _, noteAlias := range noteNotes {
			if noteAlias == s.Root {
				startingNoteIndex = noteIndex
			}
		}
	}
	noteList := [][]string{s.Notes[startingNoteIndex]}
	noteOffset := startingNoteIndex
	for index, offset := range s.Modes[s.Mode] {
		if index < 6 {
			noteOffset += offset
			noteOffset = noteOffset % 12
			noteList = append(noteList, s.Notes[noteOffset])
		}
	}
	return noteList
}

func (s *Scale) GetNotes() []string {
	rootNoteLetter := string(s.Root[0])
	rootNoteLetterIndex := strings.Index(NoteLetters, rootNoteLetter)
	scaleLetters := strings.Split(NoteLetters[rootNoteLetterIndex:]+NoteLetters[:rootNoteLetterIndex], "")
	scale := []string{}
	noteOptions := s.getNoteOptions()
	for letterIndex, letterValue := range scaleLetters {
		noteOptionLetters := noteOptions[letterIndex]
		for _, noValue := range noteOptionLetters {
			if strings.HasPrefix(noValue, letterValue) {
				scale = append(scale, noValue)
			}
		}
	}
	return scale
}

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
	s.Modes = CreateModes()
	s.Notes = map[int][]string{
		0:  []string{"B#", "C", "Dbb"},
		1:  []string{"C#", "Db"},
		2:  []string{"D", "Ebb"},
		3:  []string{"D#", "Eb"},
		4:  []string{"E", "Fb"},
		5:  []string{"E#", "F", "Gbb"},
		6:  []string{"F#", "Gb"},
		7:  []string{"F##", "G", "Abb"},
		8:  []string{"G#", "Ab"},
		9:  []string{"A", "Bbb"},
		10: []string{"A#", "Bb"},
		11: []string{"B", "Cb"},
	}
	return s, nil
}
