package main

import (
	"web-tracking"
)

func main() {
	m, err := web_tracking.NewWebTracking()
	if err != nil {
		panic(err)
	}

	if err := m.Start(); err != nil {
		panic(err)
	}
}
