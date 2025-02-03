package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"users/controllers"
)

func UserRoutes() http.Handler {
	router := chi.NewRouter()

	router.Post("/create", controllers.CreateUser)
	router.Get("/", controllers.GetUsers)
	router.Get("/{id}", controllers.GetUser)
	router.Put("/update/{id}", controllers.UpdateUser)
	router.Delete("/delete/{id}", controllers.DeleteUser)

	return router
}
