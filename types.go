package isocalutils

import (
  "fmt"
  "time"
	"github.com/snabb/isoweek"
)

var date_format = "2006-01-02-mon"

// MyString returns a "time.Time" as "2006-01-02-mon"
func MyString(t time.Time) string {
  return t.Format(date_format)
}

var days_of_week = []string {
  "sun", "mon", "tue", "wed", "thu", "fri", "sat", "sun" }

// IsoYear is range 1.. 32,000+ A.D.
type IsoYear int16
// IsoWeek is range 1..53
type IsoWeek int8
// IsoDay is range 1..7, or 0 to refer to an entire IsoWeek
type IsoDay int8

// IsoWD - if IsoDay == 0 it means IsoWeek refers
// to the whole week, not a particular day
type IsoWD struct {
  IsoWeek
  IsoDay
}

// IsoYWD completely specifies a day
// (or a week, if IsoYWD.IsoWD.IsoDay == 0)
type IsoYWD struct {
  IsoYear
  IsoWD
}

// NewYMD 
func NewYMD(y,m,d int) time.Time {
  CheckYMD(y,m,d)
  // Since we are in Europe, we set the time to 12 noon so that it stays
  // the same day anywhere away from the International Date Line.
	p := time.Date(y,time.Month(m),d, 12,0,0,0, time.UTC)
  return p
}

func NewIsoYWD(y,w,d int) *IsoYWD{
  CheckYWD(y,w,d)
  p := new(IsoYWD)
  if y <= 200 {
    y += 1900
  }
  p.IsoYear = IsoYear(y)
  p.IsoWeek = IsoWeek(w)
  p.IsoDay  = IsoDay(d)
  return p
}

func CheckYMD(y,m,d int) {
  if y <= 200 {
    y += 1900
  }
  if y < 1900 || y > 2100 || m < 0 || m > 12 || d < 0 || d > 31 {
    panic(fmt.Sprintf("YWD: bad arg(s): y,m,d=%d,%d,%d \n", y,m,d))
  }
}

func CheckYWD(y,w,d int) {
  if y <= 200 {
    y += 1900
  }
  if y < 1900 || y > 2100 || w < 0 || w > 53 || d < 0 || d > 31 {
    panic(fmt.Sprintf("YWD: bad arg(s): y,w,d=%d,%d,%d \n", y,w,d))
  }
}

func NewIsoYWDfromYMD(y,im,d int) *IsoYWD{
  CheckYMD(y,im,d)
  if y <= 200 {
    y += 1900
  }
  m := time.Month(im)
  // FromDate returns ISO 8601 week number of a date.
  yy,ww := isoweek.FromDate(y,m,d)
  // ISOWeekday returns the ISO 8601 weekday number of given day.
  // (1 = Mon, 2 = Tue,.. 7 = Sun)
  dd := isoweek.ISOWeekday(y,m,d)
  fmt.Printf("NewIsoYWDfromYMD: %04d-%02d-%02d => %04d-W%02d-%02d \n", y,im,d, yy,ww,dd)
  return NewIsoYWD(yy,ww,dd)
}

func (id IsoWD) IsWeek() bool {
  return id.IsoDay == 0
}

func (id IsoWD) String() string {
  return fmt.Sprintf("W%02d-%02d", id.IsoWeek, id.IsoDay)
}

func (id IsoYWD) String() string {
  if id.IsWeek() {
    return fmt.Sprintf("%4dw%02d", id.IsoYear, id.IsoWeek)
  }
  return fmt.Sprintf("%4d-W%02d-%d", id.IsoYear, id.IsoWeek, id.IsoDay)
}

type IsoYearDescription struct {
  IsoYear
  LeapWeek      bool
  LeapDay       bool
  Easter        IsoWeek // Always Sunday, i.e. IsoDay == 7
  Midsummer     IsoWD   // The day celebrated, can be Fri..Mon
  Thanksgiving  IsoWeek // Always Thursday, i.e. IsoDay == 4
  GregorianYMDs []time.Time // [7]; can be nil
}

var LongYears1900to2099 []IsoYear = []IsoYear {
  1903, 1931, 1959, 1987,
  1908, 1936, 1964, 1992,
  1914, 1942, 1970, 1998,
  1920, 1948, 1976,
  1925, 1953, 1981,
  2004, 2032, 2060, 2088,
  2009, 2037, 2065, 2093,
  2015, 2043, 2071, 2099,
  2020, 2048, 2076,
  2026, 2054, 2082,
}

var IsoYearDescriptions1900to2099 []IsoYearDescription

func init() {
  IsoYearDescriptions1900to2099 = make([]IsoYearDescription, 200)
  var iYear IsoYear
  for iYear = 1900; iYear < 2100; iYear++ {
    idx := iYear - 1900
    pYD := (IsoYearDescriptions1900to2099[idx])
    pYD.LeapWeek = iYear.HasLeapWeek()
    pYD.LeapDay  = iYear.IsLeapYear()
    if pYD.LeapDay != iYear.IsLeapYearTime() {
      panic("Long Year Leap Week OOPS")
    }
    // TODO:510 Leap Years
    // TODO:570 Movable holidays
  }
}

func (iy IsoYear) IsLeapYear() bool {
  return iy%400 == 0 || (iy%4 == 0 && iy%100 != 0)
}

// IsLongYear is valid for 1900 .. 2099
func (iy IsoYear) HasLeapWeek() bool {
  return isoYearIsInSlice(iy, LongYears1900to2099)
}

func isoYearIsInSlice(iy IsoYear, iys []IsoYear) bool {
  for i := range iys {
    if iy == iys[i] {
      return true
    }
  }
  return false
}

// https://www.socketloop.com/tutorials/golang-how-to-determine-if-a-year-is-leap-year
func (iy IsoYear) IsLeapYearTime() bool {
  // convert int to Time - use the last day of the year, which is 31st December
  year := time.Date(int(iy), time.December, 31, 0, 0, 0, 0, time.Local)
  days := year.YearDay()
	return (days > 365)
 }
