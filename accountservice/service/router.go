package service

import (
	"github.com/gorilla/mux"
)
func NewRouter() *mux.Router{

	//Create an instance of the Gorilla router
	router := mux.NewRouter().StrictSlash(true)

	for _,route := range routes{
		// Attach each route, uses a Builder-like pattern to set each route up.
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}
