package controllers

import (
	"01cloud-payment/internal/models"
	"01cloud-payment/internal/util"
	"net/http"

	"github.com/DATA-DOG/go-sqlmock"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Server struct {
	config util.Config
	DB     models.Store
	Router *mux.Router
	Mock   sqlmock.Sqlmock
}

func NewServer(config util.Config, store models.Store) (*Server, error) {

	server := &Server{
		config: config,
		DB:     store,
		Router: mux.NewRouter(),
	}
	server.initializeRoutes()
	return server, nil
}

func (server *Server) Run(addr string) {
	log.Info("Listening to port http://localhost:8080")
	log.Fatal(http.ListenAndServe(addr, server.Router))
}
