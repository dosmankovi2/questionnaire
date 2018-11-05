package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/dosmankovi2/questionnaire/controllers"

	"github.com/dosmankovi2/questionnaire/app"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.Use(app.TokenAuthentication)

	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")

	router.HandleFunc("/api/questionnaire/{id}", controllers.GetQuestionnaire).Methods("GET")
	router.HandleFunc("/api/questionnaire", controllers.CreateQuestionnaire).Methods("POST")
	router.HandleFunc("/api/questionnaire", controllers.UpdateQuestionnaire).Methods("PUT")
	router.HandleFunc("/api/questionnaire/{id}", controllers.DeleteQuestionnaire).Methods("DELETE")

	router.HandleFunc("/api/question/{id}", controllers.GetQuestion).Methods("GET")
	router.HandleFunc("/api/question", controllers.CreateQuestion).Methods("POST")
	router.HandleFunc("/api/question", controllers.UpdateQuestion).Methods("PUT")
	router.HandleFunc("/api/question/{id}", controllers.DeleteQuestion).Methods("DELETE")

	router.HandleFunc("/api/type/{id}", controllers.GetQuestionType).Methods("GET")
	router.HandleFunc("/api/type", controllers.CreateQuestionType).Methods("POST")
	router.HandleFunc("/api/type", controllers.UpdateQuestionType).Methods("PUT")
	router.HandleFunc("/api/type/{id}", controllers.DeleteQuestionType).Methods("DELETE")

	router.HandleFunc("/api/choice/{id}", controllers.GetChoice).Methods("GET")
	router.HandleFunc("/api/choice", controllers.CreateChoice).Methods("POST")
	router.HandleFunc("/api/choice", controllers.UpdateChoice).Methods("PUT")
	router.HandleFunc("/api/choice/{id}", controllers.DeleteChoice).Methods("DELETE")

	router.HandleFunc("/api/answer/{id}", controllers.GetAnswer).Methods("GET")
	router.HandleFunc("/api/answer", controllers.CreateAnswer).Methods("POST")
	router.HandleFunc("/api/answer", controllers.UpdateAnswer).Methods("PUT")
	router.HandleFunc("/api/answer/{id}", controllers.DeleteAnswer).Methods("DELETE")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8001"
	}

	fmt.Println("Started server...")

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Print(err)
	}
}
