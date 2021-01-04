package scheduler

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

// Resident is a resident's full name and their preferred week on call
type Resident struct {
	Name          string
	PreferredWeek int
}

// Calls produces a list of days that a resident is on call for the provided block
func (r Resident) Calls(block []Day) []Day {
	callDays := make([]Day, 0)

	for _, day := range block {
		if *day.Oak == r {
			callDays = append(callDays, day)
		}
	}

	return callDays
}

// Residents parses a csv of resident names and preferred weeks on call
func Residents(path string) []Resident {
	var residents []Resident

	file, err := os.Open(path)

	if err != nil {
		log.Fatal("Error reading csv")
	}

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		name := record[0]
		preferredWeek, err := strconv.Atoi(record[1])
		if err != nil {
			log.Fatal(err)
		}
		residents = append(residents, Resident{Name: name, PreferredWeek: preferredWeek})
	}

	return residents
}
