package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/oliveira-a-rafael/mycareer-api/controllers"
)

type Server struct {
	Config            ServerConfig
	AccountController controllers.AccountHandler
}

type ServerConfig struct {
	Port         string
	AllowedHosts []string
}

func (s *Server) Run() {

	router := s.SetupRoutes()

	router.Use(mux.CORSMethodMiddleware(router))

	log.Println("Server started at port", s.Config.Port)

	log.Fatal(http.ListenAndServe(":"+s.Config.Port,
		handlers.CORS(
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Accept", "Accept-Endcoding", "Content-Length"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins(s.Config.AllowedHosts))(router)))

}

func (s *Server) SetupRoutes() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/health", TempHealthCheck).Methods("GET")

	s.routesAccount(router)

	return router
}

func (s *Server) routesAccount(router *mux.Router) {
	accountRouter := router.PathPrefix("/user").Subrouter()
	// @TODO remover
	accountRouter.HandleFunc("/teste", s.AccountController.GetAccount).Methods("GET")
	accountRouter.HandleFunc("/new", s.AccountController.CreateAccount).Methods("POST", "OPTIONS")
	// r.HandleFunc("/user/login", controllers.Authenticate).Methods("POST", "OPTIONS")
}

var TempHealthCheck = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode("ok"); err != nil {
		log.Println(err.Error())
	}
}
