package main

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
)

func main() {
	cl := ContactList{}
	actions := map[string]func(*ContactList){
		"1": addContact,
		"2": removeContact,
		"3": findContact,
		"4": listContacts,
		"5": saveContacts,
		"6": loadContacts,
		"7": func(cl *ContactList) {
			fmt.Println("Exiting...")
			os.Exit(0)
		},
	}

	for {
		fmt.Println("\nContact Manager Menu:")
		fmt.Println("1. Add Contact")
		fmt.Println("2. Remove Contact")
		fmt.Println("3. Find Contact")
		fmt.Println("4. List Contacts")
		fmt.Println("5. Save Contacts to File")
		fmt.Println("6. Load Contacts from File")
		fmt.Println("7. Exit")
		fmt.Print("Choose an option: ")

		var choice string
		fmt.Scanln(&choice)

		action, exists := actions[choice]
		if exists {
			action(&cl)
		} else {
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func addContact(cl *ContactList) {
	var name, phone, email string
	fmt.Print("Enter Name: ")
	fmt.Scanln(&name)

	phone = enterPhone()
	email = enterEmail()

	cl.AddContact(Contact{name, phone, email})
	fmt.Println("Contact added.")
}

func enterPhone() string {
	var phone string
	for {
		fmt.Print("Enter Phone: ")
		fmt.Scanln(&phone)

		// Validate phone format for Ukraine
		phoneRegex := `^\+380\d{9}$`
		matched, _ := regexp.MatchString(phoneRegex, phone)
		if matched {
			break
		} else {
			fmt.Println("Invalid phone format. Please enter a valid Ukrainian phone number (e.g., +380XXXXXXXXX).")
		}
	}
	return phone
}

func enterEmail() string {
	var email string
	for {
		fmt.Print("Enter Email: ")
		fmt.Scanln(&email)

		// Validate email format
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		matched, _ := regexp.MatchString(emailRegex, email)
		if matched {
			break
		} else {
			fmt.Println("Invalid email format. Please try again.")
		}
	}
	return email
}

func removeContact(cl *ContactList) {
	var name string
	fmt.Print("Enter Name of Contact to Remove: ")
	fmt.Scanln(&name)
	cl.RemoveContact(name)
}

func findContact(cl *ContactList) {
	var name string
	fmt.Print("Enter Name of Contact to Find: ")
	fmt.Scanln(&name)
	contact := cl.FindContact(name)
	if contact != nil {
		fmt.Printf("Found: %v", contact)
	} else {
		fmt.Println("Contact not found.")
	}
}

func listContacts(cl *ContactList) {
	fmt.Println("\nAll Contacts:")
	fmt.Printf("%v", cl)
}

const fileName = "contacts.json"

func saveContacts(cl *ContactList) {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cl)
	if err != nil {
		fmt.Println("Error saving contacts:", err)
	} else {
		fmt.Println("Contacts saved successfully.")
	}
}

func loadContacts(cl *ContactList) {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(cl)
	if err != nil {
		fmt.Println("Error loading contacts:", err)
	} else {
		fmt.Println("Contacts loaded successfully.")
	}
}
