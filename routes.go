package main

import (
	"./app/http/handlers"
)

func (i *App) initializeRoutes() {
	i.Router.HandleFunc("/activities", handlers.getActivities).Methods("GET")
	i.Router.HandleFunc("/activity", handlers.createActivity).Methods("POST")
	i.Router.HandleFunc("/activity/{id:[0-9]+}", handlers.getActivity).Methods("GET")
	i.Router.HandleFunc("/activity/{id:[0-9]+}", handlers.updateActivity).Methods("PUT")
	i.Router.HandleFunc("/activity/{id:[0-9]+}", handlers.deleteActivity).Methods("DELETE")
}