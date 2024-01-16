package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start of month: ", StartOfMonth(time.Now()))
	fmt.Println("End of month: ", EndOfMonth(time.Now()))
	fmt.Println("Start of day of week: ", StartOfDayOfWeek(time.Now()))
	fmt.Println("End of day of week: ", EndOfDayOfWeek(time.Now()))
	fmt.Println("Start and end of weeks of month: ", StartAndEndOfWeeksOfMonth(2024, 1))
	fmt.Println("Week number in month: ", WeekNumberInMonth(time.Now()))
	fmt.Println("Start of year: ", StartOfYear(time.Now()))
	fmt.Println("End of year: ", EndOfYear(time.Now()))
	fmt.Println("Start of quarter: ", StartOfQuarter(time.Now()))
	fmt.Println("End of quarter: ", EndOfQuarter(time.Now()))
	start, end := CurrentWeekRange(time.Now().Location())
	fmt.Println("Start: ", start, " and End: ", end, " of current week")
	fmt.Println("Duration between: ", DurationBetween(time.Now(), time.Now().AddDate(0, 0, 7)))
	date, _ := ParseDateStringWithFormat("25-Jan-2022", "02-Jan-2006")
	fmt.Println("Date with format: ", date)
	fmt.Println("Business days: ", AddBusinessDays(time.Now(), 50))
	fmt.Println("Format duration: ", FormatDuration(time.Hour*24*3+time.Hour*4+time.Minute*15+time.Second*30))

	// // Take current time and load it into “America/Newyork” timezone
	// a()
	// // Take current millis and convert it into timestamp (use “Africa/Nairobi”)
	// b()
	// // Get date-time after 2 months, before 1 week and after 1 year
	// c()
	// // Get start of the day and end of the day from current time
	// d()
}

func StartOfMonth(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
}

func EndOfMonth(date time.Time) time.Time {
	firstDayOfNextMonth := StartOfMonth(date).AddDate(0, 1, 0)
	return firstDayOfNextMonth.Add(-time.Second)
}

func StartOfDayOfWeek(date time.Time) time.Time {
	daysSinceSunday := int(date.Weekday())
	return date.AddDate(0, 0, -daysSinceSunday)
}

func EndOfDayOfWeek(date time.Time) time.Time {
	daysUntilSaturday := 6 - int(date.Weekday())
	return date.AddDate(0, 0, daysUntilSaturday)
}

func StartAndEndOfWeeksOfMonth(year, month int) []struct{ Start, End time.Time } {
	loc, err := time.LoadLocation("Asia/Kolkata")
	if err != nil {
		fmt.Println("Error loading location:", err)
	}

	startOfMonth := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, loc)
	weeks := make([]struct{ Start, End time.Time }, 0)
	for current := startOfMonth; current.Month() == time.Month(month); current = current.AddDate(0, 0, 7) {
		startOfWeek := StartOfDayOfWeek(current)
		endOfWeek := EndOfDayOfWeek(current)
		if endOfWeek.Month() != time.Month(month) {
			endOfWeek = EndOfMonth(current)
		}
		weeks = append(weeks, struct{ Start, End time.Time }{startOfWeek, endOfWeek})
	}
	return weeks
}

func WeekNumberInMonth(date time.Time) int {
	startOfMonth := StartOfMonth(date)
	_, week := date.ISOWeek()
	_, startWeek := startOfMonth.ISOWeek()
	return week - startWeek + 1
}

func StartOfYear(date time.Time) time.Time {
	return time.Date(date.Year(), time.January, 1, 0, 0, 0, 0, date.Location())
}

func EndOfYear(date time.Time) time.Time {
	startOfNextYear := StartOfYear(date).AddDate(1, 0, 0)
	return startOfNextYear.Add(-time.Second)
}

func StartOfQuarter(date time.Time) time.Time {
	quarter := (int(date.Month()) - 1) / 3
	startMonth := time.Month(quarter*3 + 1)
	return time.Date(date.Year(), startMonth, 1, 0, 0, 0, 0, date.Location())
}

func EndOfQuarter(date time.Time) time.Time {
	startOfNextQuarter := StartOfQuarter(date).AddDate(0, 3, 0)
	return startOfNextQuarter.Add(-time.Second)
}

func CurrentWeekRange(timeZone *time.Location) (startOfWeek, endOfWeek time.Time) {
	now := time.Now().In(timeZone)
	startOfWeek = StartOfDayOfWeek(now)
	endOfWeek = EndOfDayOfWeek(now)
	return startOfWeek, endOfWeek
}

func DurationBetween(start, end time.Time) time.Duration {
	return end.Sub(start)
}

func ParseDateStringWithFormat(dateString, format string) (time.Time, error) {
	parsedTime, err := time.Parse(format, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return parsedTime, nil
}

func AddBusinessDays(startDate time.Time, daysToAdd int) time.Time {
	currentDate := startDate
	for i := 0; i < daysToAdd; {
		currentDate = currentDate.AddDate(0, 0, 1)
		if currentDate.Weekday() != time.Saturday && currentDate.Weekday() != time.Sunday {
			i++
		}
	}
	return currentDate
}

func FormatDuration(duration time.Duration) string {
	days := int(duration.Hours() / 24)
	hours := int(duration.Hours()) % 24
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60
	return fmt.Sprintf("%dd %02dh %02dm %02ds", days, hours, minutes, seconds)
}

func a() {
	// Get the current time in the UTC timezone
	currentTime := time.Now()

	// Load the "America/New_York" timezone
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Convert the current time to the "America/New_York" timezone
	currentTimeInNewYork := currentTime.In(loc)

	fmt.Println("Current Time:", currentTime)
	fmt.Println("Current Time (New York):", currentTimeInNewYork)
}

func b() {
	// Get the current time in milliseconds
	currentMillis := time.Now().UnixMilli()

	// Load the "Africa/Nairobi" timezone
	loc, err := time.LoadLocation("Africa/Nairobi")
	if err != nil {
		fmt.Println("Error loading location:", err)
		return
	}

	// Convert milliseconds to a time.Time value
	currentTimeInNairobi := time.Unix(0, currentMillis*int64(time.Millisecond)).In(loc)

	// Format the timestamp
	timestamp := currentTimeInNairobi.Format("2006-01-02 15:04:05")

	fmt.Println("Current Milliseconds:", currentMillis)
	fmt.Println("Timestamp in Africa/Nairobi:", timestamp)
}

func c() {
	// Get the current date and time
	currentTime := time.Now()

	// Calculate date after 2 months
	dateAfter2Months := currentTime.AddDate(0, 2, 0)

	// Calculate date before 1 week
	dateBefore1Week := currentTime.AddDate(0, 0, -7)

	// Calculate date after 1 year
	dateAfter1Year := currentTime.AddDate(1, 0, 0)

	// Format dates as strings
	formatDate := func(date time.Time) string {
		return date.Format("2006-01-02 15:04:05")
	}

	fmt.Println("Current Date and Time:", formatDate(currentTime))
	fmt.Println("Date After 2 Months:", formatDate(dateAfter2Months))
	fmt.Println("Date Before 1 Week:", formatDate(dateBefore1Week))
	fmt.Println("Date After 1 Year:", formatDate(dateAfter1Year))
}

func d() {
	// Get the current time
	currentTime := time.Now()

	// Get the start of the day (midnight)
	startOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())

	// Get the end of the day (just before midnight)
	endOfDay := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 999999999, currentTime.Location())

	// Format times as strings
	formatTime := func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05.999999999")
	}

	fmt.Println("Current Time:", formatTime(currentTime))
	fmt.Println("Start of the Day (Midnight):", formatTime(startOfDay))
	fmt.Println("End of the Day (Just Before Midnight):", formatTime(endOfDay))
}
