package guards

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var LogEntryRegex = regexp.MustCompile(`\[((\d{4}-\d{2}-\d{2}) (\d{2}:(\d{2})))\] (.*)`)
var GuardBeginsShiftRegex = regexp.MustCompile(`Guard #(\d+) begins shift`)

func SleepyFromShiftsInput(reader io.Reader) int {
	scanner := bufio.NewScanner(reader)

	var fullEntries []string
	for scanner.Scan() {
		fullEntry := scanner.Text()
		fullEntries = append(fullEntries, fullEntry)
	}

	sort.Strings(fullEntries)

	guardsSleepTime := make(map[int]int)
	guardsSleepByMinute := make(map[int][]int)

	var fellAsleepAtMin int
	var wokeUpAtMin int
	var guardInShift int

	for _, fullEntry := range fullEntries {
		entryParts := LogEntryRegex.FindStringSubmatch(fullEntry)
		entryMsg := entryParts[5]

		entryMinute, _ := strconv.ParseInt(entryParts[4], 10, 64)

		switch {
		case strings.Contains(entryMsg, "begins shift"):
			shiftBeginsEntry := GuardBeginsShiftRegex.FindStringSubmatch(entryMsg)
			guardInShift64, _ := strconv.ParseInt(shiftBeginsEntry[1], 10, 64)
			guardInShift = int(guardInShift64)
			break

		case strings.Contains(entryMsg, "falls asleep"):
			fellAsleepAtMin = int(entryMinute)
			break

		case strings.Contains(entryMsg, "wakes up"):
			wokeUpAtMin = int(entryMinute)

			guardsSleepTime[guardInShift] += wokeUpAtMin - fellAsleepAtMin
			if guardsSleepByMinute[guardInShift] == nil {
				guardsSleepByMinute[guardInShift] = make([]int, 60, 60)
			}

			for i := fellAsleepAtMin; i < wokeUpAtMin; i++ {
				guardsSleepByMinute[guardInShift][i]++
			}

			break
		}
	}

	var sleepiestGuard int
	var sleepiestMinute int
	sleepiestNTimes := 0
	for guardID, sleepByMinute := range guardsSleepByMinute {
		for minute, nTimesSleeping := range sleepByMinute {
			if nTimesSleeping > sleepiestNTimes {
				sleepiestNTimes = nTimesSleeping
				sleepiestMinute = minute
				sleepiestGuard = guardID
			}
		}
	}

	return sleepiestGuard * sleepiestMinute
}
