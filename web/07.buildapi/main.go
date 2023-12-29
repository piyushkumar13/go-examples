package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Course struct {
	CourseId    string  `json:"courseId"`
	CourseName  string  `json:"courseName"`
	CoursePrice int     `json:"coursePrice"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var seedCourses = []Course{

	{"1", "Data Structures", 23, &Author{"Piyush Kumar", "www.pk.com"}},
	{"2", "System Design", 50, &Author{"Kumar Sir", "www.kumar.com"}},
	{"3", "Algorithms", 23, &Author{"ABC Kumar", "www.abc.com"}},
}

func main() {

	fmt.Println("Starting server")

	router := mux.NewRouter()

	router.HandleFunc("/courses", getCourses).Methods("GET")
	router.HandleFunc("/courses/{id}", getCourse).Methods("GET")
	router.HandleFunc("/courses", createCourse).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func getCourse(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)

	receivedId := params["id"]
	fmt.Printf("Getting a course with id=%s\n", receivedId)

	w.Header().Set("Content-Type", "application/json")

	var selectedCourse Course
	for _, course := range seedCourses {

		if course.CourseId == receivedId {
			selectedCourse = course
		}
	}
	json.NewEncoder(w).Encode(selectedCourse)
}

func getCourses(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Getting all courses")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(seedCourses)
}

func createCourse(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Adding course to the list")
	w.Header().Set("Content-Type", "application/json")

	body := r.Body

	if body == nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Please provide some body")
		return
	}

	var course Course
	json.NewDecoder(body).Decode(&course)

	if course.CourseName == "" {
		w.WriteHeader(412)

		json.NewEncoder(w).Encode("Please provide some data in the body")
		return
	}

	seedCourses = append(seedCourses, course)

	w.WriteHeader(202)
	json.NewEncoder(w).Encode(seedCourses)
}

func deleteCourse(w http.ResponseWriter, r *http.Request) {

}
