package scheduler

import (
	"log"

	"github.com/rickb777/date"
)

// Resident is a resident's full name and their preferred week on call
type Resident struct {
	Name          string
	PreferredWeek int
}

// Day is the day's number in a given block, the date, and the residents assigned to that Day
type Day struct {
	ID    int
	Date  date.Date
	Oak   *Resident
	Maple *Resident
}

// BlockDates returns a slice of days for a block given the block's start date
func BlockDates(start string) []Day {
	startDate, err := date.AutoParse(start)
	if err != nil {
		log.Fatal("Error parsing start date.")
	}

	var dates []Day
	for i := 0; i < 28; i++ {
		dates = append(dates, Day{
			ID:    i,
			Date:  startDate.Add(date.PeriodOfDays(i)),
			Oak:   nil,
			Maple: nil,
		})
	}

	return dates
}
