package cmd

import "time"

type filter struct {
	since      time.Time
	keywords   string
	categories string
}
