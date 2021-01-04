package scheduler

import (
	"log"

	"github.com/rickb777/date"
)

// BlockLength is the length of a residency block in days
const BlockLength int = 28

// Day is the day's number in a given block, the date, and the residents assigned to that Day
type Day struct {
	ID    int
	Date  date.Date
	Oak   *Resident
	Maple *Resident
}

// Filled returns true if both call slots have been filled
func (d Day) Filled() bool {
	return d.Oak != nil && d.Maple != nil
}

// BlockDates returns a slice of days for a block given the block's start date
func BlockDates(start string) []Day {
	startDate, err := date.AutoParse(start)
	if err != nil {
		log.Fatal("Error parsing start date.")
	}

	var dates []Day
	for i := 0; i < BlockLength; i++ {
		dates = append(dates, Day{
			ID:    i,
			Date:  startDate.Add(date.PeriodOfDays(i)),
			Oak:   nil,
			Maple: nil,
		})
	}

	return dates
}
