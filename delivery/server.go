package delivery

import (
	"final-project-kelompok-1/config"
	"final-project-kelompok-1/delivery/controller"
	"final-project-kelompok-1/manager"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uc     manager.UseCaseManager
	engine *gin.Engine
	host   string
}

func (s *Server) setupControllers() {
	rg := s.engine.Group("/api/v1")
	controller.NewAdminTrainerController(s.uc.AdminTrainerUseCase(), rg).Route()
	controller.NewRoleController(s.uc.RoleUseCase(), rg).Route()
	controller.NewStudentController(s.uc.StudentUseCase(), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repoManager := manager.NewRepoManager(infraManager)
	useCaseManager := manager.NewUseCaseManager(repoManager)

	return &Server{
		uc:     useCaseManager,
		engine: engine,
		host:   fmt.Sprintf(":%s", cfg.ApiPort),
	}
}
