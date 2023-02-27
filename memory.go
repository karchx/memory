package memory

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Ram struct {
	Total     int
	Free      int
	Available int
	SwapTotal int
	SwapFree  int
}

const (
  MemInfo = "/proc/meminfo"
)

// parseMeValue parse value from line in /proc/meminfo
func parseMeValue(v string) int {
  in, _ := strconv.Atoi(strings.Fields(v)[1])
  return in
}

func GemRam() (Ram, error) {
  lines, err := readLines(MemInfo, 16)
  if err != nil {
    return Ram{}, err
  }
 
  ram := Ram{
    Total: parseMeValue(lines[0]),
    Free: parseMeValue(lines[1]),
    Available: parseMeValue(lines[2]),
    SwapTotal: parseMeValue(lines[14]),
    SwapFree: parseMeValue(lines[15]),
  }

  return ram, nil
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
