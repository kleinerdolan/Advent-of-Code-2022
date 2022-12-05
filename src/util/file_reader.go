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

func ReadFileUntrimmed(day string, mode string) []string {
	data, err := os.ReadFile("input/" + day + mode + ".txt")
	if err != nil {
		return nil
	}

	lines := strings.Split(string(data), "\n")
	return lines[:len(lines)-1]
}
