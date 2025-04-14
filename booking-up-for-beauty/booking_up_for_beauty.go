package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05" // Correct layout

	t, err := time.Parse(layout, date)
	if err != nil {
		return time.Time{}
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("ERROR %v \n", err)
		return false
	}
	t1 := time.Now()
	return parsedDate.Sub(t1).Minutes() <= 0
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// Thursday, July 25, 2019 13:45:00
	layout := "Monday, January 2, 2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("ERROR %v \n", err)
		return false
	}
	return parsedDate.Hour() >= 12 && parsedDate.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Printf("ERROR %v \n", err)
		return ""
	}
	return "You have an appointment on " + parsedDate.Format("Monday, January 2, 2006,") + " at " + parsedDate.Format("15:04.")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)

}
