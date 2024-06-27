// main_test.go
package main

import (
	"testing"
)

func TestParse(t *testing.T) {
	parser := &StandardCronParser{}

	tests := []struct {
		cronString string
		expected   *CronSchedule
	}{
		{
			"*/15 0 1,15 * 1-5 /usr/bin/find",
			&CronSchedule{
				Minutes:     []int{0, 15, 30, 45},
				Hours:       []int{0},
				DaysOfMonth: []int{1, 15},
				Months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DaysOfWeek:  []int{1, 2, 3, 4, 5},
				Command:     "/usr/bin/find",
			},
		},
		{
			"0 12 * * 0 /usr/bin/find",
			&CronSchedule{
				Minutes:     []int{0},
				Hours:       []int{12},
				DaysOfMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				Months:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
				DaysOfWeek:  []int{0},
				Command:     "/usr/bin/find",
			},
		},
		{
			"5 0 * 8 * /usr/bin/find",
			&CronSchedule{
				Minutes:     []int{5},
				Hours:       []int{0},
				DaysOfMonth: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31},
				Months:      []int{8},
				DaysOfWeek:  []int{0, 1, 2, 3, 4, 5, 6, 7},
				Command:     "/usr/bin/find",
			},
		},
	}

	for _, test := range tests {
		schedule, err := parser.Parse(test.cronString)
		if err != nil {
			t.Fatalf("Parse failed: %v", err)
		}

		if !equal(schedule, test.expected) {
			t.Fatalf("expected %+v, got %+v", test.expected, schedule)
		}
	}
}

func equal(a, b *CronSchedule) bool {
	if len(a.Minutes) != len(b.Minutes) ||
		len(a.Hours) != len(b.Hours) ||
		len(a.DaysOfMonth) != len(b.DaysOfMonth) ||
		len(a.Months) != len(b.Months) ||
		len(a.DaysOfWeek) != len(b.DaysOfWeek) ||
		a.Command != b.Command {
		return false
	}

	for i := range a.Minutes {
		if a.Minutes[i] != b.Minutes[i] {
			return false
		}
	}

	for i := range a.Hours {
		if a.Hours[i] != b.Hours[i] {
			return false
		}
	}

	for i := range a.DaysOfMonth {
		if a.DaysOfMonth[i] != b.DaysOfMonth[i] {
			return false
		}
	}

	for i := range a.Months {
		if a.Months[i] != b.Months[i] {
			return false
		}
	}

	for i := range a.DaysOfWeek {
		if a.DaysOfWeek[i] != b.DaysOfWeek[i] {
			return false
		}
	}

	return true
}
