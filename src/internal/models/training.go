package models

import (
	"fmt"
	"time"
)

type Training struct {
	ID       int       `json:"id"`
	Date     time.Time `json:"date"`
	Type     string    `json:"type"`
	Duration int       `json:"duration"`
	Notes    string    `json:"notes"`
}

func (t Training) String() string {
	return fmt.Sprintf("[%d] %s — %s (%d min)| %s", t.ID, t.Date.Format("2006-01-02 15:04"), t.Type, t.Duration, t.Notes)
}
