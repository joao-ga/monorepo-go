package controllers

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"scheduling/models"
	"scheduling/services"
)

func CreateSchedule(w http.ResponseWriter, r *http.Request) {
	var schedule models.Schedule

	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := services.CreateSchedule(context.Background(), schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(result)
}

func GetSchedule(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	schedule, err := services.GetScheduleByid(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(schedule)
}

func GetAllSchedules(w http.ResponseWriter, r *http.Request) {
	schedules, err := services.GetAllSchedules(context.Background())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(schedules)
}

func UpdateSchedule(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var schedule map[string]interface{}

	if err := json.NewDecoder(r.Body).Decode(&schedule); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := services.UpdateSchedule(context.Background(), id, schedule)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"schedule": "updated"})
}

func DeleteSchedule(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	err := services.DeleteSchedule(context.Background(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"schedule": "deleted"})
}
