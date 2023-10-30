package models

import (
	"gorm.io/gorm"
)

// DiagnosticTest represents a diagnostic test in the healthcare application.
type DiagnosticTest struct {
	gorm.Model
	TestName          string
	TestPrice         float64
	NormalValue       string
	Units             string
	DiagnosticCenters []DiagnosticCenter `gorm:"many2many:diagnostic_center_tests;"` // Slice to store associated DiagnosticCenter entities
}
