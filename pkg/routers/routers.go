package routers

import (
	"github.com/arkiralor/GoURLShortener/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/", controllers.Index).Methods("GET")
}
