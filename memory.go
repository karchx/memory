package memory

import (
	"bufio"
	"math"
	"os"
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

func readLines(s string, n int) ([]string, error) {
	f, err := os.Open(s)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	lines := make([]string, 0, n)
	sc := bufio.NewScanner(f)
	for i := 0; i < n && sc.Scan(); i++ {
		lines = append(lines, sc.Text())
	}
	return lines, nil
}

func readFirstLine(s string) (string, error) {
  f, err := os.Open(s)
  if err != nil {
    return "", err
  }
  defer f.Close()
  sc := bufio.NewScanner(f)
  sc.Scan()
  return sc.Text(), nil
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

func truncateResult(n float64) float64 {
	return math.Trunc(n*10) / 10
}
