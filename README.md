-- Create the "healthcare" database
CREATE DATABASE healthcare;

-- Use the "healthcare" database
USE healthcare;

-- Create the "User" table for user authentication
CREATE TABLE User (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(20) NOT NULL
);

-- Insert sample data into the "User" table
INSERT INTO User (username, password, role) VALUES
('Adi', 'adi@123', 'admin'),
('Rahul', 'ra@123', 'patient'),
('Siddhant', 'sid@123', 'patient');

-- Create the "Patient" table for patient information
CREATE TABLE Patient (
    patientid INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phoneno VARCHAR(20) NOT NULL,
    age INT NOT NULL,
    gender VARCHAR(10) NOT NULL
);

INSERT INTO Patient (patientid, name, phoneno,age,gender) VALUES
(11,'Rahul',987,12,'M'),
(12,'Siddhant',786,14,'M'),
(13,'soni',127,20,'F');

-- Create the "Appointment" table for appointments
CREATE TABLE Appointment (
    id INT AUTO_INCREMENT PRIMARY KEY,
    appointmentdate DATETIME NOT NULL,
    approvalstatus VARCHAR(20) NOT NULL,
    diagnosticTests TEXT,
    patient INT NOT NULL,
    diagnosticCenter INT NOT NULL,
    testResult INT
);

-- Create the "TestResult" table for test results
CREATE TABLE TestResult (
    id INT AUTO_INCREMENT PRIMARY KEY,
    testreading VARCHAR(255) NOT NULL,
    condition TEXT,
    appointment INT NOT NULL
);

-- Create the "DiagnosticTest" table for diagnostic tests
CREATE TABLE DiagnosticTest (
    id INT AUTO_INCREMENT PRIMARY KEY,
    testname VARCHAR(255) NOT NULL,
    testprice DECIMAL(10, 2) NOT NULL,
    normalvalue VARCHAR(255) NOT NULL,
    units VARCHAR(20) NOT NULL,
    diagnosticCenter INT NOT NULL
);

-- Create the "DiagnosticCenter" table for diagnostic centers
CREATE TABLE DiagnosticCenter (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    contactno VARCHAR(20) NOT NULL,
    address TEXT,
    contactemail VARCHAR(255) NOT NULL,
    serviceoffered TEXT,
    tests TEXT
);
