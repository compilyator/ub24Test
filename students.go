package main

type StudentInfo struct {
	ID        string
	FirstName string
	LastName  string
	Grade     int
}

type StudentManagement struct {
	Students map[string]StudentInfo
}

func (records *StudentManagement) AddStudent(id string, firstName string, lastName string, grade int) {
	info := StudentInfo{ID: id, FirstName: firstName, LastName: lastName, Grade: grade}
	records.Students[id] = info
}

func (records *StudentManagement) GetStudent(id string) StudentInfo {
	return records.Students[id]
}
