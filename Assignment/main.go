// Cron Parser main.go
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type StandardCronParser struct{}

// CronSchedule represents the parsed fields of a cron string
type CronSchedule struct {
	Minutes     []int
	Hours       []int
	DaysOfMonth []int
	Months      []int
	DaysOfWeek  []int
	Command     string
}

// expandRange expands the cron range string to a list of numbers
func expandRange(field string, min int, max int) []int {
	if field == "*" {
		var result []int
		for i := min; i <= max; i++ {
			result = append(result, i)
		}
		return result
	}

	var result []int
	parts := strings.Split(field, ",")
	for _, part := range parts {
		rangeMin := min
		rangeMax := max

		if strings.Contains(part, "/") {
			subParts := strings.Split(part, "/")
			rangePart := subParts[0]
			step, _ := strconv.Atoi(subParts[1])

			if rangePart != "*" {
				rangeBounds := strings.Split(rangePart, "-")
				rangeMin, _ = strconv.Atoi(rangeBounds[0])
				if len(rangeBounds) > 1 {
					rangeMax, _ = strconv.Atoi(rangeBounds[1])
				}
			}

			for i := rangeMin; i <= rangeMax; i += step {
				result = append(result, i)
			}
		} else if strings.Contains(part, "-") {
			rangeBounds := strings.Split(part, "-")
			rangeMin, _ := strconv.Atoi(rangeBounds[0])
			rangeMax, _ = strconv.Atoi(rangeBounds[1])
			for i := rangeMin; i <= rangeMax; i++ {
				result = append(result, i)
			}
		} else {
			val, _ := strconv.Atoi(part)
			result = append(result, val)
		}
	}
	return result
}

// Parse parses the given cron string into a CronSchedule struct
func (p *StandardCronParser) Parse(cronString string) (*CronSchedule, error) {
	fields := strings.Fields(cronString)

	if len(fields) != 6 {
		return nil, fmt.Errorf("invalid cron string")
	}

	minuteField := fields[0]
	hourField := fields[1]
	dayOfMonthField := fields[2]
	monthField := fields[3]
	dayOfWeekField := fields[4]
	commandField := fields[5]

	minutes := expandRange(minuteField, 0, 59)
	hours := expandRange(hourField, 0, 23)
	daysOfMonth := expandRange(dayOfMonthField, 1, 31)
	months := expandRange(monthField, 1, 12)
	daysOfWeek := expandRange(dayOfWeekField, 0, 7)

	return &CronSchedule{
		Minutes:     minutes,
		Hours:       hours,
		DaysOfMonth: daysOfMonth,
		Months:      months,
		DaysOfWeek:  daysOfWeek,
		Command:     commandField,
	}, nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide valid number of arguments")
		return
	}

	cronString := os.Args[1]
	parser := &StandardCronParser{}
	schedule, err := parser.Parse(cronString)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%-14s", "minute")
	for _, val := range schedule.Minutes {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Printf("%-14s", "hour")
	for _, val := range schedule.Hours {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Printf("%-14s", "day of month")
	for _, val := range schedule.DaysOfMonth {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Printf("%-14s", "month")
	for _, val := range schedule.Months {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Printf("%-14s", "day of week")
	for _, val := range schedule.DaysOfWeek {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	fmt.Printf("%-14s%s\n", "command", schedule.Command)
}
