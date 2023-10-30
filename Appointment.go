package models

import (
	"gorm.io/gorm"
)

// Appointment represents an appointment in the healthcare application.
type Appointment struct {
	gorm.Model
	AppointmentDate  int
	ApprovalStatus   string
	DiagnosticTests  []DiagnosticTest // Slice to store DiagnosticTest entities
	Patient          Patient
	DiagnosticCenter DiagnosticCenter
	TestResults      []TestResult // Slice to store TestResult entities
}
