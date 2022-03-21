package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"name"`
	TaskDetail string `json:"task_detail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:         "1",
		TaskName:   "Task 1",
		TaskDetail: "Task 1 Detail",
		Date:       "2020-01-01"}
	tasks = append(tasks, task)
	task2 := Tasks{
		ID:         "2",
		TaskName:   "Task 2",
		TaskDetail: "Task 2 Detail",
		Date:       "2020-02-01"}
	tasks = append(tasks, task2)
	// fmt.Println(tasks)

}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
	fmt.Println("Endpoint Hit: getTasks")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	taskID := mux.Vars(r)
	flag := false
	for i := 0; i < len(tasks); i++ {
		if taskID["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			flag = true
			break
		}
	}
	if !flag {
		json.NewEncoder(w).Encode(map[string]string{"error": "task not found"})
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func delete(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func update(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func handleRoutes() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/gettask/{id}", getTask).Methods("GET")
	router.HandleFunc("/gettasks", getTasks).Methods("GET")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", delete).Methods("DELETE")
	router.HandleFunc("/update/{id}", update).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
func main() {
	allTasks()
	fmt.Println("Server up and running!")
	handleRoutes()
}
