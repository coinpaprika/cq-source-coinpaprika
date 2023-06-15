package client

import (
	"regexp"
	"strconv"
	"time"
)

// WithCustomDurations add custom intervals
// parser:
//   - d (day) equal to 24h only integer value
func WithCustomDurations(func(s string) (time.Duration, error)) func(s string) (time.Duration, error) {
	return func(s string) (time.Duration, error) {
		re := regexp.MustCompile(`([-+]?[0-9]+)d`)
		daysString := re.FindStringSubmatch(s)
		if len(daysString) == 0 {
			return time.ParseDuration(s)
		}
		days, err := strconv.ParseInt(daysString[1], 10, 64)
		if err != nil {
			return 0, err
		}

		return time.Duration(days) * 24 * time.Hour, nil
	}
}
