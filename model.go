package main

import (
	"time"

	"gorm.io/gorm"
)

type TimePrecisionExperience struct {
	gorm.Model
	ExpTime time.Time
}

