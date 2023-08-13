package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/tormoder/fit"
	"github.com/vlekakis/TTUtils/proc"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

const (
	LAP_SWIM   = 17
	OPEN_WATER = 18
)

func readFitFile(filePath string) {

	dat, err := os.ReadFile(filePath)
	check(err)

	fit, err := fit.Decode(bytes.NewReader(dat))
	check(err)

	// Get the actual activity
	activity, err := fit.Activity()
	check(err)

	// Print the sport of the first Session message
	fmt.Println("Sessions Length: ", len(activity.Sessions))

	for _, session := range activity.Sessions {
		switch session.SubSport {
		case OPEN_WATER:
			proc.ProcessOpenWaterSession(*session)
			proc.ProcessActivityLaps(*activity)
		default:
			fmt.Println("Unknown SubSport: ", session.SubSport)
		}
	}

}

func main() {
	fitFile := flag.String("fitFile", "", "fit file to analyze")
	flag.Parse()

	if *fitFile == "" {
		fmt.Println("fitFile is required")
		os.Exit(1)
	}

	readFitFile(*fitFile)
}
