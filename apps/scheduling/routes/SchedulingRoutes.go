package routes

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"scheduling/controllers"
)

func ScheduleRoutes() http.Handler {
	router := chi.NewRouter()

	router.Post("/create", controllers.CreateSchedule)
	router.Get("/", controllers.GetAllSchedules)
	router.Get("/{id}", controllers.GetSchedule)
	router.Put("/update/{id}", controllers.UpdateSchedule)
	router.Delete("/delete/{id}", controllers.DeleteSchedule)

	return router
}
