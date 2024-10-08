package main

import "fmt"

type StudentInfo struct {
	ID        string
	FirstName string
	LastName  string
	Grade     int
}

type StudentManagement struct {
	Students map[string]StudentInfo
}

// AddStudent adds a new student to the student list.
func (sm *StudentManagement) AddStudent(student StudentInfo) {
	sm.Students[student.ID] = student
	fmt.Printf("Student with ID %s added successfully.\n", student.ID)
}

// UpdateStudent updates an existing student's information.
func (sm *StudentManagement) UpdateStudent(student StudentInfo) {
	if _, exists := sm.Students[student.ID]; exists {
		sm.Students[student.ID] = student
		fmt.Printf("Student with ID %s updated successfully.\n", student.ID)
	} else {
		fmt.Printf("Student with ID %s does not exist.\n", student.ID)
	}
}

// DeleteStudent removes a student from the student list by ID.
func (sm *StudentManagement) DeleteStudent(id string) {
	if _, exists := sm.Students[id]; exists {
		delete(sm.Students, id)
		fmt.Printf("Student with ID %s deleted successfully.\n", id)
	} else {
		fmt.Printf("Student with ID %s does not exist.\n", id)
	}
}

// GetStudent retrieves a student's information by ID.
func (sm *StudentManagement) GetStudent(id string) {
	if student, exists := sm.Students[id]; exists {
		fmt.Printf("Student Info - ID: %s, Name: %s %s, Grade: %d\n", student.ID, student.FirstName, student.LastName, student.Grade)
	} else {
		fmt.Printf("Student with ID %s not found.\n", id)
	}
}

// ListStudents prints the details of all students.
func (sm *StudentManagement) ListStudents() {
	if len(sm.Students) == 0 {
		fmt.Println("No students found.")
		return
	}

	for _, student := range sm.Students {
		fmt.Printf("ID: %s, Name: %s %s, Grade: %d\n", student.ID, student.FirstName, student.LastName, student.Grade)
	}
}
