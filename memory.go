package memory

import (
	"fmt"
	"strconv"
	"strings"
)

type Ram struct {
	Total     float64
	Free      int
	Available int
	SwapTotal int
	SwapFree  int
}

const (
	MemInfo = "/proc/meminfo"
)

const (
	_  = iota //ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

// parseMeValue parse value from line in /proc/meminfo
func parseMeValue(v string) int {
	in, _ := strconv.Atoi(strings.Fields(v)[1])
	return in
}

// get Ram from /proc/meminfo
func GetRam() (Ram, error) {
	lines, err := readLines(MemInfo, 16)
	if err != nil {
		return Ram{}, err
	}

	ram := Ram{
		Total:     ConverBytes(parseMeValue(lines[0]), "GB"),
		Free:      parseMeValue(lines[1]),
		Available: parseMeValue(lines[2]),
		SwapTotal: parseMeValue(lines[14]),
		SwapFree:  parseMeValue(lines[15]),
	}

	return ram, nil
}

func GetTotalRam() (int, error) {
	line, err := readFirstLine(MemInfo)
	if err != nil {
		return 0, err
	}

	return parseMeValue(line), nil
}

// GetFreeRam ram amount
func GetFreeRam() (int, error) {
	line, err := readLine(MemInfo, 2)
	if err != nil {
		return 0, err
	}

	return parseMeValue(line), nil
}

// GetAvalibleRam ram amount
func GetAvalibleRam() (int, error) {
	line, err := readLine(MemInfo, 3)
	if err != nil {
		return 0, err
	}

	return parseMeValue(line), nil
}

func GetSwapFree() (int, error) {
	fmt.Printf("[DEBUG]: %v", MemInfo)
	line, err := readLine(MemInfo, 16)
	if err != nil {
		return 0, err
	}
	return parseMeValue(line), nil
}

// ConverBytes return memory in MG|GB with base in bytes
func ConverBytes(n int, prefix string) float64 {
	switch prefix {
	case "MB":
		return truncateResult(float64(n) / (1 << 10))
	case "GB":
		return truncateResult(float64(n) / (1 << 20))
	default:
		return float64(0.0)
	}

}
