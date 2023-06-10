package main

import (
	"fmt"
	"log"

	"github.com/karchx/memory"
)

func main() {
  mem, err := memory.GetRam()
  totalMem, err := memory.GetTotalRam()

  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("RAM: ", mem)
  fmt.Println("TOTAL RAM: ", memory.ConverBytes(totalMem, "GB"))
}
