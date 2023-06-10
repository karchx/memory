package main

import (
	"fmt"
	"log"

	"github.com/karchx/memory"
)

func main() {
  mem, err := memory.GetRam()
  totalMem, err := memory.GetTotalRam()
  freeMem, err := memory.GetFreeRam()
  availableMem, err := memory.GetAvalibleRam()

  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("RAM: ", mem)
  fmt.Println("TOTAL RAM: ", memory.ConverBytes(totalMem, "GB"))
  fmt.Println("FREE RAM: ", memory.ConverBytes(freeMem, "GB"))
  fmt.Println("AVAILABLE RAM: ", memory.ConverBytes(availableMem, "GB"))
}
