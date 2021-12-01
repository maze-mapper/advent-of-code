// Advent of Code 2018 - Day 4 - Tests
package day4

import (
	"reflect"
	"testing"
	"time"
)

// sleeping is a nested map of guard to minutes with a count of how often they are asleep at that minute
var sleeping = map[int]map[int]int{
	10: map[int]int{
		5: 1, 6: 1, 7: 1, 8: 1, 9: 1,
		10: 1, 11: 1, 12: 1, 13: 1, 14: 1, 15: 1, 16: 1, 17: 1, 18: 1, 19: 1,
		20: 1, 21: 1, 22: 1, 23: 1, 24: 2, 25: 1, 26: 1, 27: 1, 28: 1,
		30: 1, 31: 1, 32: 1, 33: 1, 34: 1, 35: 1, 36: 1, 37: 1, 38: 1, 39: 1,
		40: 1, 41: 1, 42: 1, 43: 1, 44: 1, 45: 1, 46: 1, 47: 1, 48: 1, 49: 1,
		50: 1, 51: 1, 52: 1, 53: 1, 54: 1,
	},
	99: map[int]int{
		36: 1, 37: 1, 38: 1, 39: 1,
		40: 2, 41: 2, 42: 2, 43: 2, 44: 2, 45: 3, 46: 2, 47: 2, 48: 2, 49: 2,
		50: 1, 51: 1, 52: 1, 53: 1, 54: 1,
	},
}

func TestParseEvent(t *testing.T) {
	tests := map[string]struct {
		input string
		want  event
	}{
		"Begin shift": {
			input: "[1518-11-01 00:00] Guard #10 begins shift",
			want: event{
				timestamp: time.Date(1518, time.November, 01, 0, 0, 0, 0, time.UTC),
				guard:     10,
				id:        BeginShift,
			},
		},
		"Fall asleep": {
			input: "[1518-11-01 00:05] falls asleep",
			want: event{
				timestamp: time.Date(1518, time.November, 01, 0, 5, 0, 0, time.UTC),
				guard:     0,
				id:        FallAsleep,
			},
		},
		"Wake up": {
			input: "[1518-11-01 00:25] wakes up",
			want: event{
				timestamp: time.Date(1518, time.November, 01, 0, 25, 0, 0, time.UTC),
				guard:     0,
				id:        WakeUp,
			},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := parseEvent(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("Got %v, want %v", got, tc.want)
			}
		})
	}
}

func TestPart1(t *testing.T) {
	wantGuard := 10
	wantMinute := 24

	gotGuard, gotMinute := part1(sleeping)
	if gotGuard != wantGuard {
		t.Errorf("Got guard %d, want guard %d", gotGuard, wantGuard)
	}
	if gotMinute != wantMinute {
		t.Errorf("Got minute %d, want minute %d", gotMinute, wantMinute)
	}
}

func TestPart2(t *testing.T) {
	wantGuard := 99
	wantMinute := 45

	gotGuard, gotMinute := part2(sleeping)
	if gotGuard != wantGuard {
		t.Errorf("Got guard %d, want guard %d", gotGuard, wantGuard)
	}
	if gotMinute != wantMinute {
		t.Errorf("Got minute %d, want minute %d", gotMinute, wantMinute)
	}
}
