package util

import (
	"fmt"
	"strconv"
)

func ParseTime(uptimeSeconds string) (string, error) {
	seconds, err := strconv.Atoi(uptimeSeconds)
	if err != nil {
		return "", err
	}

	days := seconds / 86400
	hours := (seconds % 86400) / 3600
	minutes := (seconds % 3600) / 60
	secs := seconds % 60

	switch {
	case days > 0:
		return fmt.Sprintf("%dd%dh", days, hours), nil
	case hours > 0:
		return fmt.Sprintf("%dh%dm", hours, minutes), nil
	case minutes > 0:
		return fmt.Sprintf("%dm%ds", minutes, secs), nil
	default:
		return fmt.Sprintf("%ds", secs), nil
	}
}
