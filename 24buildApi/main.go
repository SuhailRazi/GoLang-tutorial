package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Model for course - seperate file'
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"-"`
	Author      *Author `json:"auhtor"`
}

type Author struct {
	FullName string `json:"fullname"`
	Website  string `json:"website"`
}

// fakeDB
var courses []Course

// middleware - seperate file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API learning")

	r := mux.NewRouter()

	// seeding of data
	courses = append(courses, Course{
		CourseId:    "CS101",
		CourseName:  "Introduction to Computer Science",
		CoursePrice: 49,
		Author: &Author{
			FullName: "John Doe",
			Website:  "http://www.johndoe.com",
		},
	})

	courses = append(courses, Course{
		CourseId:    "3243",
		CourseName:  "ReactJs",
		CoursePrice: 2222,
		Author: &Author{
			FullName: "Suhail",
			Website:  "go.dev",
		},
	})
	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/courses/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/courses", createOneCourse).Methods("POST")
	r.HandleFunc("/courses/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}", deleteOneCourse).Methods("DELETE")

	// listening to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API in GO lang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(r)

	// loop throught the courses, func matching id and return the response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}
	// check if title is dupicate
	// loop, stoping point of is matches with course.coursename

	// what about - {}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside")
		return
	}

	// generateUniqueId , convert to string
	// append new course to course

	// rand.Seed(time.Now().UnixNano()) // old method
	// course.CourseId = strconv.Itoa(rand.Intn(100))

	course.CourseId = uuid.New().String()
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from req
	params := mux.Vars(r)

	// loop throught the value, get id, remove the value, add the new values with id from params

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		} else {
			// Todo send a response when id is not found
			json.NewEncoder(w).Encode("Item not availabe")
		}
	}

}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// loop throught the courses, find the id, remove the index, index+1

	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode("The selected record have been deleted")
	}

}
