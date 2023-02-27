package main

import (
	"fmt"
	"log"

	"github.com/karchx/memory"
)

func main() {
  mem, err := memory.GemRam()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(mem)
}
