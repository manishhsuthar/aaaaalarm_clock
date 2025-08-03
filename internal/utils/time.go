package utils

import (
    "time"
)

// GetCurrentTime returns the current local time.
func GetCurrentTime() time.Time {
    return time.Now()
}

// FormatTime formats the given time.Time into a string.
func FormatTime(t time.Time) string {
    return t.Format("15:04:05")
}