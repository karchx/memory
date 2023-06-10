package memory

import (
	"bufio"
	"math"
	"os"
)

func truncateResult(n float64) float64 {
	return math.Trunc(n*10) / 10
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

func readLine(s string, n int) (string, error) {
	f, err := os.Open(s)
	if err != nil {
		return "", err
	}

	defer f.Close()

	var line string
	sc := bufio.NewScanner(f)
	for i := 0; i < n && sc.Scan(); i++ {
		if i == n-1 {
			line = sc.Text()
		}
	}
	return line, nil
}
