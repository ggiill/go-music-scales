package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	scales "github.com/gturetsky/go-music-scales"
)

func main() {

	subcommandScale := "scale"
	scaleToNotesCommand := flag.NewFlagSet(subcommandScale, flag.ExitOnError)
	scaleToNotesRootPtr := scaleToNotesCommand.String("root", "", "Root note (required)")
	scaleToNotesModePtr := scaleToNotesCommand.String("mode", "", "Mode (required)")

	subcommandNotes := "notes"
	notesToScalesCommand := flag.NewFlagSet(subcommandNotes, flag.ExitOnError)
	notesToScalesListPtr := notesToScalesCommand.String("list", "", "List of notes (required)")

	if len(os.Args) == 1 {
		fmt.Println("Must provide at least 1 subcommand [scale, notes]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case subcommandScale:
		scaleToNotesCommand.Parse(os.Args[2:])
	case subcommandNotes:
		notesToScalesCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	switch {
	case scaleToNotesCommand.Parsed():
		if *scaleToNotesRootPtr == "" {
			scaleToNotesCommand.PrintDefaults()
			os.Exit(1)
		}
		s, err := scales.NewScale(*scaleToNotesRootPtr, *scaleToNotesModePtr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			s.Identify()
		}
	case notesToScalesCommand.Parsed():
		if *notesToScalesListPtr == "" {
			notesToScalesCommand.PrintDefaults()
			os.Exit(1)
		}
		notesToScalesNotes := append([]string{*notesToScalesListPtr}, notesToScalesCommand.Args()...)
		scales, err := scales.GetScalesFromNotes(notesToScalesNotes)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		} else {
			scalesFormatted := strings.Join(scales, ", ")
			fmt.Println(scalesFormatted)
		}
	}
}
