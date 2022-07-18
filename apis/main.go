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

type Course struct {
	Id     string  `json:"id"`
	Name   string  `json:"name"`
	Price  string  `json:"price"`
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

var courses []Course

func (c *Course) IsEmpty() bool {
	return c.Name == ""
}

func main() {

	fmt.Println("Starting the application...")
	var r = mux.NewRouter()

	courses = append(
		courses,
		Course{Id: "1", Name: "Go", Price: "100", Author: &Author{Fullname: "Golang", Website: "golang.org"}},
		Course{Id: "2", Name: "Python", Price: "200", Author: &Author{Fullname: "Python", Website: "python.org"}},
		Course{Id: "3", Name: "Java", Price: "300", Author: &Author{Fullname: "Java", Website: "java.org"}},
	)

	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", r))

}

func serveHome(w http.ResponseWriter, r *http.Request) {
	var _, e = w.Write([]byte("<p>Welcome to the Home Page!</p>"))
	if e != nil {
		panic(e)
	}
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println(courses)
	w.Header().Set("Content-Type", "application/json")
	var e = json.NewEncoder(w).Encode(courses)
	if e != nil {
		panic(e)
	}
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var params = mux.Vars(r)

	for _, course := range courses {
		if course.Id == params["id"] {
			var e = json.NewEncoder(w).Encode(course)
			if e != nil {
				panic(e)
			}
			return
		}
	}
	var e = json.NewEncoder(w).Encode("No course found")
	if e != nil {
		panic(e)
	}
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Body == nil {
		var e = json.NewEncoder(w).Encode("Please send a request body")
		if e != nil {
			panic(e)
		}
	}

	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)

	if course.IsEmpty() {
		var e = json.NewEncoder(w).Encode("Please send a valid course")
		if e != nil {
			panic(e)
		}
		return
	}

	for _, c := range courses {
		if course.Name == c.Name {
			var _ = json.NewEncoder(w).Encode("COurse already exists")
			return
		}
	}

	rand.Seed(time.Now().UnixNano())
	course.Id = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	_ = json.NewEncoder(w).Encode(course)
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	for index, course := range courses {
		if course.Id == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.Id = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w.Header().Get("Content-Type"))
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	for i, c := range courses {
		if c.Id == params["id"] {
			courses = append(courses[:i], courses[i+1:]...)
			var e = json.NewEncoder(w).Encode(c)
			if e != nil {
				panic(e)
			}
			_ = json.NewEncoder(w).Encode("Course deleted")
			return

		}
	}

}
