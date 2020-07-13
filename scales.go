package scales

import (
	"fmt"
	"os"
	"strings"
)

const (
	NoteLetters   string = "ABCDEFG"
	NoteHalf      int    = 1
	NoteWhole     int    = 2
	NoteWholeHalf int    = 3
)

type Scale struct {
	root  string
	mode  string
	modes map[string][]int
	notes map[int][]string
}

func (s *Scale) Identify() {
	switch string(s.root[0]) {
	case "A", "E":
		fmt.Print("I am an ", s.root, " ", s.mode, " scale ")
	default:
		fmt.Print("I am a ", s.root, " ", s.mode, " scale ")
	}
	fmt.Println("and my notes are", strings.Join(s.GetNotes(), ", "))
}

func (s *Scale) getNoteOptions() [][]string {
	var startingNoteIndex int
	for noteIndex, noteNotes := range s.notes {
		for _, noteAlias := range noteNotes {
			if noteAlias == s.root {
				startingNoteIndex = noteIndex
			}
		}
	}
	noteList := [][]string{s.notes[startingNoteIndex]}
	noteOffset := startingNoteIndex
	for index, offset := range s.modes[s.mode] {
		if index < 6 {
			noteOffset += offset
			noteOffset = noteOffset % 12
			noteList = append(noteList, s.notes[noteOffset])
		}
	}
	return noteList
}

func (s *Scale) GetNotes() []string {
	rootNoteLetter := string(s.root[0])
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

func NewScale(root string, mode string) *Scale {
	s := new(Scale)
	s.root = root
	s.mode = mode
	s.modes = map[string][]int{}
	modes := []string{"Ionian", "Dorian", "Phrygian", "Lydian", "Myxolydian", "Aeolian", "Lydian"}
	modeSteps := []int{NoteWhole, NoteWhole, NoteHalf, NoteWhole, NoteWhole, NoteWhole, NoteHalf}
	for modeIndex, modeName := range modes {
		s.modes[modeName] = append(modeSteps[modeIndex:], modeSteps[:modeIndex]...)
	}
	// Gotta include Major & Minor (which are aliases of the modes above):
	s.modes["Major"] = s.modes["Ionian"]
	s.modes["Minor"] = s.modes["Aeolian"]
	// Weird scales:
	s.modes["Harmonic Minor"] = []int{NoteWhole, NoteHalf, NoteWhole, NoteWhole, NoteHalf, NoteWholeHalf, NoteHalf}
	s.notes = map[int][]string{
		0:  []string{"B#", "C"},
		1:  []string{"C#", "Db"},
		2:  []string{"D"},
		3:  []string{"D#", "Eb"},
		4:  []string{"E", "Fb"},
		5:  []string{"E#", "F"},
		6:  []string{"F#", "Gb"},
		7:  []string{"G"},
		8:  []string{"G#", "Ab"},
		9:  []string{"A"},
		10: []string{"A#", "Bb"},
		11: []string{"B", "Cb"},
	}
	return s
}

func scales() {
	args := os.Args
	s := NewScale(args[1], args[2])
	s.Identify()
}
