package api

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/ajtroup1/platinum-trophy-tracker/service/account"
	"github.com/ajtroup1/platinum-trophy-tracker/service/game"
	"github.com/ajtroup1/platinum-trophy-tracker/service/user"
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

func clearConsole() {
    cmd := exec.Command("cmd", "/c", "cls")
    cmd.Stdout = os.Stdout
    if err := cmd.Run(); err != nil {
        log.Fatalf("failed to clear console: %v", err)
    }
}

func (s *APIServer) Run() error {
	clearConsole()
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	accountStore := account.NewStore(s.db)
	accountHandler := account.NewHandler(accountStore)
	accountHandler.RegisterRoutes(subrouter)

	gameStore := game.NewStore(s.db)
	gameHandler := game.NewHandler(gameStore)
	gameHandler.RegisterRoutes(subrouter)

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
