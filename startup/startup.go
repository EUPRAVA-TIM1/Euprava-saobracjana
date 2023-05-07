package startup

import (
	"EuprvaSaobracajna/config"
	"EuprvaSaobracajna/handlers"
	"EuprvaSaobracajna/repo"
	"EuprvaSaobracajna/service"
	"context"
	"fmt"
	muxHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

func (server Server) setup() handlers.SaobracjanaHandler {
	secretRepo := repo.NewSecretRepoSql(server.config.MysqlPort, server.config.MySqlRootPass, server.config.MySqlHost)
	ssoService := service.NewSsoService(ConstructServiceUrl(server.config.SsoServiceHost, server.config.SsoServicePort), server.config.SSOIssuer)
	mupService := service.NewMupService(ConstructServiceUrl(server.config.MupServiceHost, server.config.MupServicePort))
	sudService := service.NewSudService(ConstructServiceUrl(server.config.SudServiceHost, server.config.SudServicePort))
	saobracajnaRepo := repo.NewSaobracjanaRepoSql(server.config.MysqlPort, server.config.MySqlRootPass, server.config.MySqlHost)
	saobracjanaService := service.NewSaobracjanaService(saobracajnaRepo, mupService, sudService)
	jwtService := service.NewJwtService(secretRepo, saobracajnaRepo, ssoService)
	return handlers.NewSaobracajnaHandler(saobracjanaService, jwtService)
}

func ConstructServiceUrl(host, port string) string {
	return fmt.Sprintf("http://%s:%s", host, port)
}

func (server Server) Start() {

	r := mux.NewRouter()

	corsHandler := muxHandlers.CORS(
		muxHandlers.AllowedOrigins([]string{"*"}),
		muxHandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		muxHandlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	h := server.setup()
	h.Init(r)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", server.config.Port),
		Handler: corsHandler(r),
	}

	wait := time.Second * 15
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)

	log.Printf("Listening on port = %s\n", server.config.Port)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("error shutting down server %s", err)
	}
	log.Println("server gracefully stopped")

}
