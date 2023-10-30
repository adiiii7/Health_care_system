package models

import (
	"gorm.io/gorm"
)

// DiagnosticCenter represents a diagnostic center in the healthcare application.
type DiagnosticCenter struct {
	gorm.Model
	Name           string
	ContactNo      string
	Address        string
	ContactEmail   string
	ServiceOffered string
	Tests          []DiagnosticTest // Slice to store DiagnosticTest entities
}
