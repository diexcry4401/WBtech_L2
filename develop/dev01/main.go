package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	t, err := get_time()
	if err != nil {
		os.Exit(1)
	}
	fmt.Println("Time:", t)
}

func get_time() (time.Time, error) {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return time.Time{}, err
	}
	return t, err
}
