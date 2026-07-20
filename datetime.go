package gofakeit

import (
	"time"
)

var currentYear = time.Now().Year()

func Date() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (f *Faker) Date() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func date(f *Faker) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func PastDate() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (f *Faker) PastDate() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func pastDate(f *Faker) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func FutureDate() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (f *Faker) FutureDate() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func futureDate(f *Faker) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func DateRange(start, end time.Time) time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (f *Faker) DateRange(start, end time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func dateRange(f *Faker, start time.Time, end time.Time) time.Time {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func NanoSecond() int { _ = "STUB: not implemented"; return 0 }

func (f *Faker) NanoSecond() int { _ = "STUB: not implemented"; return 0 }

func nanoSecond(f *Faker) int { _ = "STUB: not implemented"; return 0 }

func Second() int { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Second() int { _ = "STUB: not implemented"; return 0 }

func second(f *Faker) int { _ = "STUB: not implemented"; return 0 }

func Minute() int { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Minute() int { _ = "STUB: not implemented"; return 0 }

func minute(f *Faker) int { _ = "STUB: not implemented"; return 0 }

func Hour() int { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Hour() int { _ = "STUB: not implemented"; return 0 }

func hour(f *Faker) int { _ = "STUB: not implemented"; return 0 }

func Day() int { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Day() int { _ = "STUB: not implemented"; return 0 }

func day(f *Faker) int { _ = "STUB: not implemented"; return 0 }

func WeekDay() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) WeekDay() string { _ = "STUB: not implemented"; return "" }

func weekDay(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Month() int { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Month() int { _ = "STUB: not implemented"; return 0 }

func month(f *Faker) int { _ = "STUB: not implemented"; return 0 }

func MonthString() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) MonthString() string { _ = "STUB: not implemented"; return "" }

func monthString(f *Faker) string { _ = "STUB: not implemented"; return "" }

func Year() int { _ = "STUB: not implemented"; return 0 }

func (f *Faker) Year() int { _ = "STUB: not implemented"; return 0 }

func year(f *Faker) int { _ = "STUB: not implemented"; return 0 }

func TimeZone() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) TimeZone() string { _ = "STUB: not implemented"; return "" }

func timeZone(f *Faker) string { _ = "STUB: not implemented"; return "" }

func TimeZoneFull() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) TimeZoneFull() string { _ = "STUB: not implemented"; return "" }

func timeZoneFull(f *Faker) string { _ = "STUB: not implemented"; return "" }

func TimeZoneRegion() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) TimeZoneRegion() string { _ = "STUB: not implemented"; return "" }

func timeZoneRegion(f *Faker) string { _ = "STUB: not implemented"; return "" }

func TimeZoneAbv() string { _ = "STUB: not implemented"; return "" }

func (f *Faker) TimeZoneAbv() string { _ = "STUB: not implemented"; return "" }

func timeZoneAbv(f *Faker) string { _ = "STUB: not implemented"; return "" }

func TimeZoneOffset() float32 { _ = "STUB: not implemented"; return 0 }

func (f *Faker) TimeZoneOffset() float32 { _ = "STUB: not implemented"; return 0 }

func timeZoneOffset(f *Faker) float32 { _ = "STUB: not implemented"; return 0 }

func javaDateTimeFormatToGolangFormat(format string) string { _ = "STUB: not implemented"; return "" }

func addDateTimeLookup() { _ = "STUB: not implemented"; return }
