package models

import (
	"gorm.io/gorm"
)

// Patient represents a patient in the healthcare application.
type Patient struct {
	gorm.Model
	PatientID    string
	Name         string
	PhoneNo      string
	Age          int
	Gender       string
	Appointments []Appointment // Slice to store associated Appointment entities
}
