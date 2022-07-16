package booking

import (
	"fmt"
	"time"
)

const (
	//"7/13/2020 20:32:00"
	timeFormat1 = "1/2/2006 15:04:05"
	//("July 25, 2019 13:45:00")
	timeFormat2 = "January 2, 2006 15:04:05"
	//("Thursday, July 25, 2019 13:45:00")
	timeFormat3 = "Monday, January 2, 2006 15:04:05"

	timeFormat4 = "Monday, January 2, 2006, at 15:04."
)

// Schedule returns a time.Time from a string containing a date
//"7/13/2020 20:32:00"
func Schedule(date string) time.Time {
	time, err := time.Parse(timeFormat1, date)
	if err != nil {
		fmt.Println(err.Error())
	}
	return time
}

// HasPassed returns whether a date has passed
//"July 25, 2019 13:45:00"
func HasPassed(date string) bool {
	dt := time.Now()
	time, err := time.Parse(timeFormat2, date)
	if err != nil {
		fmt.Println(err.Error())
	}
	return time.Before(dt)
}

// IsAfternoonAppointment returns whether a time is in the afternoon
// ("Thursday, July 25, 2019 13:45:00")
func IsAfternoonAppointment(date string) bool {
	time, err := time.Parse(timeFormat3, date)
	if err != nil {
		fmt.Println(err.Error())
	}
	return time.Hour() >= 12 && time.Hour() < 18
		
}

// Description returns a formatted string of the appointment time
//("7/25/2019 13:45:00")
func Description(date string) string {
	time, err := time.Parse(timeFormat1, date)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	return fmt.Sprintf("You have an appointment on %s",time.Format(timeFormat4))
}

// AnniversaryDate returns a Time with this year's anniversary
func AnniversaryDate() time.Time {
	AnniversaryDate := time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
	return AnniversaryDate
}
