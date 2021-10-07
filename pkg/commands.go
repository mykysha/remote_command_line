package cline

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// date handles date command.
func date() []byte {
	now := time.Now()

	date := fmt.Sprint(now.Format("Mon Jan _2 15:04:05 MST 2006"))

	return []byte(date)
}

// echo handles echo command.
func echo(message []string) []byte {
	echo := strings.Join(message, " ")

	return []byte(echo)
}

// timeout handles timeout command.
func timeout(times []string) []byte {
	sleepFor := time.Second

	for _, val := range times {
		var currNumber int

		for i := 0; i < len(val)-1; i++ {
			if unicode.IsDigit(rune(val[i])) {
				currNumber *= 10

				digit, err := strconv.Atoi(string(val[i]))
				if err != nil {
					return []byte("unknown time type")
				}

				currNumber += digit
			} else {
				return []byte("unknown time type")
			}
		}

		timeType := string(val[len(val)-1])

		switch timeType {
		case "d":
			secondsInDay := 86400
			currNumber *= secondsInDay
			sleepFor += time.Second * time.Duration(currNumber)

		case "h":
			secondsInHour := 3600
			currNumber *= secondsInHour
			sleepFor += time.Second * time.Duration(currNumber)

		case "m":
			secondsInMinute := 60
			currNumber *= secondsInMinute

			sleepFor += time.Second * time.Duration(currNumber)

		case "s":
			sleepFor += time.Second * time.Duration(currNumber)

		default:
			return []byte("unknown time type")
		}
	}

	time.Sleep(sleepFor)

	return []byte("done sleeping")
}

// help handles help command.
func help() []byte {
	help := fmt.Sprintf(
		"\n\tAvailable commands are:\n" +
			"\tdate - shows current time\n" +
			"\techo - echoes text that follows\n" +
			"\ttimeout (1d 2h 3m 5s) - sleeps for given amount of time\n" +
			"\thelp / h - lists commands\n" +
			"\twho - gives info about creator")

	return []byte(help)
}

// who handles who command.
func who() []byte {
	who := fmt.Sprintf(
		"\n\tName: Derhunov Mykyta\n" +
			"\tTheme number: 3\n" +
			"\tTopic: remote command line\n" +
			"\tGroup: K27\n")

	return []byte(who)
}
