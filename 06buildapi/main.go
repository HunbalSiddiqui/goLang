package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// model for courses - file
type Course struct {
	CourseId     string  `json:"courseId"`
	CourseName   string  `json:"courseName"`
	CoursePrice  int     `json:"coursePrice"`
	CourseAuthor *Author `json:"courseAuthor"`
}

type Author struct {
	FullName string `json:"fullName"`
	Website  string `json:"website"`
}

// temp DB
var courses []Course

// middleware, helper - file
func (c *Course) isEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("Build Apis - hunbalsiddiqui.com")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJs", CoursePrice: 299, CourseAuthor: &Author{
		FullName: "Hunbal",
		Website:  "hunbalsiddiqui.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Mern", CoursePrice: 399, CourseAuthor: &Author{
		FullName: "Hunbal",
		Website:  "hunbalsiddiqui.com"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	r.HandleFunc("/courses/", createCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteCourse).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))

}

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API using Go.</h1>"))
}

func getCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")
	// grab id from request
	params := mux.Vars(r)
	fmt.Println("Get one course Params: ", params)
	// find matching id
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found.")
	return
}

func createCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
		return
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.isEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON.")
		return
	}
	// generate unqiue id, string
	rand.Seed(time.Now().UnixNano())
	// append course into courses
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
		return
	}
	// grab id from request
	params := mux.Vars(r)
	fmt.Println("Get one course Params: ", params)
	// find matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			var course Course
			courses = append(courses[:index], courses[index+1:]...) // remove existing
			_ = json.NewDecoder(r.Body).Decode((&course))
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found.")
	return
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create course")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data.")
		return
	}
	// grab id from request
	params := mux.Vars(r)
	fmt.Println("Get one course Params: ", params)
	// find matching id
	for index, course := range courses {
		if course.CourseId == params["id"] {
			var course Course
			courses = append(courses[:index], courses[index+1:]...) // remove existing
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found.")
	return
}
