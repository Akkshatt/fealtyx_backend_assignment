package repo

import (
	"fealtyx_backend_assignment/models"
	"errors"
	"sync"
)

var students []models.Student
var idCounter = 1
var mu sync.Mutex


func init() {
	
	students = []models.Student{
		{ID: 1, Name: "Alice", Age: 20, Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Age: 22, Email: "bob@example.com"},
	}
	idCounter = 3
}


func CreateStudent(student models.Student) (models.Student, error) {
	mu.Lock()
	defer mu.Unlock()

	student.ID = idCounter
	idCounter++
	students = append(students, student)

	return student, nil
}


func GetAllStudents() ([]models.Student, error) {
	mu.Lock()
	defer mu.Unlock()

	return students, nil
}


func GetStudentByID(id int) (models.Student, error) {
	mu.Lock()
	defer mu.Unlock()

	for _, student := range students {
		if student.ID == id {
			return student, nil
		}
	}
	return models.Student{}, errors.New("student not found")
}


func UpdateStudentByID(id int, updatedStudent models.Student) (models.Student, error) {
	mu.Lock()
	defer mu.Unlock()

	for i, student := range students {
		if student.ID == id {
			students[i] = updatedStudent
			students[i].ID = id
			return students[i], nil
		}
	}
	return models.Student{}, errors.New("student not found")
}


func DeleteStudentByID(id int) error {
	mu.Lock()
	defer mu.Unlock()

	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			return nil
		}
	}
	return errors.New("student not found")
}
