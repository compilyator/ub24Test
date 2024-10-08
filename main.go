package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// AddStudent adds a new student to the student list.
func AddStudent(management *StudentManagement) {
	var id, firstName, lastName string
	var grade int
	fmt.Print("Enter Student ID: ")
	fmt.Scan(&id)
	fmt.Print("Enter First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter Last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter Grade: ")
	fmt.Scan(&grade)

	student := StudentInfo{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Grade:     grade,
	}
	management.Students[student.ID] = student
	fmt.Printf("Student with ID %s added successfully.\n", student.ID)
	management.SaveToFile()
}

// UpdateStudent updates an existing student's information.
func UpdateStudent(management *StudentManagement) {
	var id, firstName, lastName string
	var grade int
	fmt.Print("Enter Student ID to update: ")
	fmt.Scan(&id)
	fmt.Print("Enter Updated First Name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter Updated Last Name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter Updated Grade: ")
	fmt.Scan(&grade)

	student := StudentInfo{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Grade:     grade,
	}

	if _, exists := management.Students[student.ID]; exists {
		management.Students[student.ID] = student
		fmt.Printf("Student with ID %s updated successfully.\n", student.ID)
		management.SaveToFile()
	} else {
		fmt.Printf("Student with ID %s does not exist.\n", student.ID)
	}
}

// DeleteStudent removes a student from the student list by ID.
func DeleteStudent(management *StudentManagement) {
	var id string
	fmt.Print("Enter Student ID to delete: ")
	fmt.Scan(&id)
	if _, exists := management.Students[id]; exists {
		delete(management.Students, id)
		fmt.Printf("Student with ID %s deleted successfully.\n", id)
		management.SaveToFile()
	} else {
		fmt.Printf("Student with ID %s does not exist.\n", id)
	}
}

// GetStudent retrieves a student's information by ID.
func GetStudent(management *StudentManagement) {
	var id string
	fmt.Print("Enter Student ID to get details: ")
	fmt.Scan(&id)
	if student, exists := management.Students[id]; exists {
		fmt.Printf("Student Info - ID: %s, Name: %s %s, Grade: %d\n", student.ID, student.FirstName, student.LastName, student.Grade)
	} else {
		fmt.Printf("Student with ID %s not found.\n", id)
	}
}

// ListStudents prints the details of all students.
func ListStudents(management *StudentManagement) {
	if len(management.Students) == 0 {
		fmt.Println("No students found.")
		return
	}

	for _, student := range management.Students {
		fmt.Printf("ID: %s, Name: %s %s, Grade: %d\n", student.ID, student.FirstName, student.LastName, student.Grade)
	}
}

// DisplayMenu displays the main menu and handles user input.
func DisplayMenu(management *StudentManagement) {
	for {
		fmt.Println("\n--- Student Management ---")
		fmt.Println("1. Add Student")
		fmt.Println("2. Update Student")
		fmt.Println("3. Delete Student")
		fmt.Println("4. Get Student by ID")
		fmt.Println("5. List All Students")
		fmt.Println("6. Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			AddStudent(management)
		case 2:
			UpdateStudent(management)
		case 3:
			DeleteStudent(management)
		case 4:
			GetStudent(management)
		case 5:
			ListStudents(management)
		case 6:
			fmt.Println("Exiting the program.")
			return
		default:
			fmt.Println("Invalid choice, please select a valid option.")
		}

		fmt.Println("Press any key to continue")
		keyboard.GetKey()
	}
}

func main() {
	// Initialize student management system
	management := StudentManagement{
		Students: make(map[string]StudentInfo),
	}

	// Load student data from file at the beginning
	management.LoadFromFile()

	// Display the main menu
	DisplayMenu(&management)

	fmt.Println("Hello world")
}
