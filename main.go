package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

func main() {
	// Initialize student management system
	management := StudentManagement{
		Students: make(map[string]StudentInfo),
	}

	// Command-line interface to interact with the system
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
			// Add a new student
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
			management.AddStudent(student)

		case 2:
			// Update an existing student
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
			management.UpdateStudent(student)

		case 3:
			// Delete a student by ID
			var id string
			fmt.Print("Enter Student ID to delete: ")
			fmt.Scan(&id)
			management.DeleteStudent(id)

		case 4:
			// Get a student's details by ID
			var id string
			fmt.Print("Enter Student ID to get details: ")
			fmt.Scan(&id)
			management.GetStudent(id)

		case 5:
			// List all students
			fmt.Println("\nListing all students:")
			management.ListStudents()

		case 6:
			// Exit the program
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Invalid choice, please select a valid option.")
		}

		fmt.Println("Press any key to continue")
		keyboard.GetKey()
	}
}
