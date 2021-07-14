package isocalutils

import(
	"fmt"
	"time"
	"github.com/snabb/isoweek" // MIT license
)

// StartDateCalcs is TBS.
type StartDateCalcs struct {
	SpecDate     time.Time
	SpecDateIso  IsoYWD
	SpecMonth1st time.Time
	SpecMonth1stIso IsoYWD
	StartDate    time.Time
	StartDateIso IsoYWD
}

// Call NewStartDate(1957, 01, 27)
// As I recall, it was a Sunday...
// We'll hafta fudge the time of day
// (i.e. chg 1030pm => noon) to avoid
// Jan.28 in the Eastern hemisphere.

// NewStartDate takes a Gregorian YMD (such as
// a date of birth) and sets the StartDate to:
// the start of the ISO week that contains
// the start of the calendar month that contains
// the specified (argument) date.
func NewStartDate(Y,iM,D int) *StartDateCalcs {
	CheckYMD(Y,iM,D)
	M := time.Month(iM)
  if Y <= 200 {
    Y += 1900
  }
	p := new(StartDateCalcs)
	p.SpecDate = NewYMD(Y,iM,D)
	p.SpecDateIso = *NewIsoYWDfromYMD(Y,iM,D)
	fmt.Printf("Args ints: %04d-%02d-%02d %s \n",
		Y,iM,D, days_of_week[p.SpecDate.Weekday()])
	println("Date spec:", MyString(p.SpecDate))
	println("(iso)spec:", p.SpecDateIso.String())

	// Rewind to the 1st of the spec month.
	// func (t Time) Date() (year int, month Month, day int)
	p.SpecMonth1st = time.Date(Y,M,1, 12,0,0,0, time.UTC)
	println("Month 1st:", MyString(p.SpecMonth1st))

	// Now identify the ISO week of it.
	y,m,d := p.SpecMonth1st.Date()
	// FromDate returns ISO 8601 week number of a date.
	yy,ww := isoweek.FromDate(y,m,d)
	// ISOWeekday returns the ISO 8601 weekday number of given day.
	// (1 = Mon, 2 = Tue,.. 7 = Sun)
	dd := isoweek.ISOWeekday(y,m,d)
	fmt.Printf("(iso) 1st: %04d-w%02d-%02d \n", yy,ww,dd)
	p.SpecMonth1stIso = *NewIsoYWD(yy,ww,dd)

	// Now rewind to the start of that ISO week of the 1st (i.e. the Monday).
	p.StartDateIso = *NewIsoYWD(yy,ww,1)
	// Or, better: StartTime returns Monday 00:00 of the given ISO 8601 week.
	p.StartDate = isoweek.StartTime(yy,ww,time.UTC).Add(12*time.Hour)
	println("(iso)Mon.:", p.StartDateIso.String())
	println("=> Monday:", MyString(p.StartDate))

	return p
}
