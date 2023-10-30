package models

import (
	"gorm.io/gorm"
)

// TestResult represents a test result in the healthcare application.
type TestResult struct {
	gorm.Model
	TestReading string
	Condition   string
	Appointment Appointment // Reference to the associated appointment
}
