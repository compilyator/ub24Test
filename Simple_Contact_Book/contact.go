package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
	Email string
}

func (c Contact) Format(f fmt.State, verb rune) {
	fmt.Fprintf(f, "Name: %s, Phone: %s, Email: %s\n", c.Name, c.Phone, c.Email)
}

type ContactList struct {
	Contacts []Contact
}

// Format contacts for printing
func (cl *ContactList) Format(f fmt.State, verb rune) {
	for _, contact := range cl.Contacts {
		fmt.Fprintf(f, "%s", contact)
	}
}

// Add a new contact to the list
func (cl *ContactList) AddContact(contact Contact) {
	cl.Contacts = append(cl.Contacts, contact)
}

// Find a contact by name
func (cl *ContactList) FindContact(name string) *Contact {
	for i := range cl.Contacts {
		if cl.Contacts[i].Name == name {
			return &cl.Contacts[i]
		}
	}
	return nil
}

// Remove a contact by name
func (cl *ContactList) RemoveContact(name string) {
	contact := cl.FindContact(name)
	if contact != nil {
		for i, c := range cl.Contacts {
			if &c == contact {
				cl.Contacts = append(cl.Contacts[:i], cl.Contacts[i+1:]...)
				fmt.Println("Contact removed.")
				return
			}
		}
	} else {
		fmt.Println("Contact not found.")
	}
}
