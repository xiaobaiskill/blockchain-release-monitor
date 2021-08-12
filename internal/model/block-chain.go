package model

import "time"

type ReleaseInfo struct {
	Name    string
	Project string
	Version string
	Url     string
	Time    time.Time
}
