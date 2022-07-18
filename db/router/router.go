package router

import (
	"github.com/AbhinavAchha/golang-tutorial/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	var router = mux.NewRouter()

	router.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")
	router.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")
	router.HandleFunc("/api/movie/{id}", controller.DeleteOneMovie).Methods("DELETE")
	router.HandleFunc("/api/movies", controller.DeleteAllMovie).Methods("DELETE")

	return router

}
