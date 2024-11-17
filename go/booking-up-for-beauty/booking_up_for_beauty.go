package booking

import (
	"fmt"
	"time"
)
//Constants taken from https://go.dev/src/time/format.go
/* const (
	_                        = iota
	stdLongMonth             = iota + stdNeedDate  // "January"
	stdMonth                                       // "Jan"
	stdNumMonth                                    // "1"
	stdZeroMonth                                   // "01"
	stdLongWeekDay                                 // "Monday"
	stdWeekDay                                     // "Mon"
	stdDay                                         // "2"
	stdUnderDay                                    // "_2"
	stdZeroDay                                     // "02"
	stdUnderYearDay                                // "__2"
	stdZeroYearDay                                 // "002"
	stdHour                  = iota + stdNeedClock // "15"
	stdHour12                                      // "3"
	stdZeroHour12                                  // "03"
	stdMinute                                      // "4"
	stdZeroMinute                                  // "04"
	stdSecond                                      // "5"
	stdZeroSecond                                  // "05"
	stdLongYear              = iota + stdNeedDate  // "2006"
	stdYear                                        // "06"
	stdPM                    = iota + stdNeedClock // "PM"
	stdpm                                          // "pm"
	stdTZ                    = iota                // "MST"
	stdISO8601TZ                                   // "Z0700"  // prints Z for UTC
	stdISO8601SecondsTZ                            // "Z070000"
	stdISO8601ShortTZ                              // "Z07"
	stdISO8601ColonTZ                              // "Z07:00" // prints Z for UTC
	stdISO8601ColonSecondsTZ                       // "Z07:00:00"
	stdNumTZ                                       // "-0700"  // always numeric
	stdNumSecondsTz                                // "-070000"
	stdNumShortTZ                                  // "-07"    // always numeric
	stdNumColonTZ                                  // "-07:00" // always numeric
	stdNumColonSecondsTZ                           // "-07:00:00"
	stdFracSecond0                                 // ".0", ".00", ... , trailing zeros included
	stdFracSecond9                                 // ".9", ".99", ..., trailing zeros omitted

	stdNeedDate       = 1 << 8             // need month, day, year
	stdNeedClock      = 2 << 8             // need hour, minute, second
	stdArgShift       = 16                 // extra argument in high bits, above low stdArgShift
	stdSeparatorShift = 28                 // extra argument in high 4 bits for fractional second separators
	stdMask           = 1<<stdArgShift - 1 // mask out argument
)
 */
const (
	//"7/13/2020 20:32:00" / "1/2/2006 15:04:05"
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
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
