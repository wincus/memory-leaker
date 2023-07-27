package main

import (
	"flag"
	"fmt"
	"time"
)

const (
	MB = 1 << 20 // bytes in a megabyte
)

var memRate int

func init() {
	// Parse the memory allocation rate from command line arguments
	flag.IntVar(&memRate, "mb-per-second", 10, "Memory allocation rate (MB/second)")
	flag.Parse()
}

func main() {
	// Convert the rate to bytes per second
	ratePerSecond := float64(memRate) * MB

	fmt.Printf("Allocating memory at a rate of %d MB/second\n", memRate)

	// Create a slice to hold the memory
	var mem []byte

	// Start the allocation loop
	ticker := time.NewTicker(1 * time.Second)

	for range ticker.C {

		// Allocate the amount of memory for this second
		bytesToAllocate := int(ratePerSecond)
		mem = append(mem, make([]byte, bytesToAllocate)...)
	}
}
