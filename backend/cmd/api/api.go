package api

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/ajtroup1/speakeasy/service/user"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr   string
	db     *sql.DB
	Router *mux.Router
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	s.Router = router

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	addr := listener.Addr().String()
	ip := strings.Split(addr, ":")[0]
	if ip == "::" || ip == "" {
		ip = "127.0.0.1"
	}
	log.Printf("Server listening on %s:%s", ip, s.addr)

	return http.Serve(listener, s.Router)
}
