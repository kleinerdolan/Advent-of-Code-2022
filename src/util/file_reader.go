package util

import (
	"os"
	"strings"
)

func ReadFile(day string, mode string) []string {
	data, err := os.ReadFile("input/" + day + mode + ".txt")
	if err != nil {
		return nil
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	return lines
}
