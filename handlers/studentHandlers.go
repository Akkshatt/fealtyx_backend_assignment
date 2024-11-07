package handlers

import (
	"encoding/json"
	"fealtyx_backend_assignment/models"
	"fealtyx_backend_assignment/repo"
	"fealtyx_backend_assignment/services"
	"fealtyx_backend_assignment/utils"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateStudent handles the route for creating a new student
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		utils.SendErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate name and age
	if student.Name == "" {
		utils.SendErrorResponse(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}
	if student.Age < 0 {
		utils.SendErrorResponse(w, "Age cannot be negative", http.StatusBadRequest)
		return
	}

	// Validate email
	if !isValidEmail(student.Email) {
		utils.SendErrorResponse(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Add the student to the database (repo layer)
	createdStudent, err := repo.CreateStudent(student)
	if err != nil {
		utils.SendErrorResponse(w, "Failed to create student", http.StatusInternalServerError)
		return
	}

	// Respond with the created student details
	utils.SendSuccessResponse(w, "Student created successfully", createdStudent)
}

// GetAllStudents handles the route to fetch all students
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	students, err := repo.GetAllStudents()
	if err != nil {
		utils.SendErrorResponse(w, "Failed to fetch students", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(w, "Students fetched successfully", students)
}

// GetStudentByID handles fetching a single student by ID
func GetStudentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	student, err := repo.GetStudentByID(id)
	if err != nil {
		utils.SendErrorResponse(w, "Student not found", http.StatusNotFound)
		return
	}

	utils.SendSuccessResponse(w, "Student fetched successfully", student)
}

// UpdateStudentByID handles updating a student by ID
func UpdateStudentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	var student models.Student
	if err := json.NewDecoder(r.Body).Decode(&student); err != nil {
		utils.SendErrorResponse(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate name and age
	if student.Name == "" {
		utils.SendErrorResponse(w, "Name cannot be empty", http.StatusBadRequest)
		return
	}
	if student.Age < 0 {
		utils.SendErrorResponse(w, "Age cannot be negative", http.StatusBadRequest)
		return
	}

	// Validate email
	if !isValidEmail(student.Email) {
		utils.SendErrorResponse(w, "Invalid email format", http.StatusBadRequest)
		return
	}

	// Update student in the repo
	updatedStudent, err := repo.UpdateStudentByID(id, student)
	if err != nil {
		utils.SendErrorResponse(w, "Failed to update student", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(w, "Student updated successfully", updatedStudent)
}

// DeleteStudentByID handles deleting a student by ID
func DeleteStudentByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := repo.DeleteStudentByID(id); err != nil {
		utils.SendErrorResponse(w, "Failed to delete student", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(w, "Student deleted successfully", nil)
}

// GenerateStudentSummary handles the route for generating a student summary using Ollama
func GenerateStudentSummary(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	student, err := repo.GetStudentByID(id)
	log.Print(err)
	if err != nil {
		utils.SendErrorResponse(w, "Student not found", http.StatusNotFound)
		return
	}

	// Create a summary prompt based on student information
	prompt := fmt.Sprintf("Summarize this student in a few words as a sentence : Name: %s, Age: %d, Email: %s", student.Name, student.Age, student.Email)

	// Call Ollama API for summary
	summary, err := services.GetStudentSummaryFromOllama(prompt)
	log.Print(err)
	if err != nil {
		utils.SendErrorResponse(w, "Failed to generate summary", http.StatusInternalServerError)
		return
	}

	utils.SendSuccessResponse(w, "Summary generated successfully", summary)
}

// Helper function to validate email format using regex
func isValidEmail(email string) bool {
	// Basic regex for validating email format
	// Improved email validation to handle more cases
	regex := `^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`
	matched, _ := regexp.MatchString(regex, email)
	return matched
}
