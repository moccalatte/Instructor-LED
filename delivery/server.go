package delivery

import (
	"final-project-kelompok-1/config"
	"final-project-kelompok-1/delivery/controller"
	"final-project-kelompok-1/delivery/middleware"
	"final-project-kelompok-1/manager"
	"final-project-kelompok-1/usecase"
	"final-project-kelompok-1/utils/common"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	uc         manager.UseCaseManager
	engine     *gin.Engine
	host       string
	logService common.MyLogger
	csvService common.CvsCommon
	auth       usecase.AuthUseCase
	jwtService common.JwtToken
}

func (s *Server) setupControllers() {
	s.engine.Use(middleware.NewLogMiddleware(s.logService).LogRequest())
	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)
	rg := s.engine.Group("/api/v1")
	controller.NewStudentController(s.uc.StudentUseCase(), rg, authMiddleware).Route()
	controller.NewCourseController(s.uc.CourseCase(), rg, authMiddleware).Route()
	controller.NewUserController(s.uc.UserUseCase(), rg, authMiddleware).Route()
	controller.NewAuthController(s.auth, rg, s.jwtService).Route()
	controller.NewQuestionController(s.uc.QuestionUseCase(), rg, authMiddleware).Route()
	controller.NewCourseDetailController(s.uc.CourseDetailUseCase(), rg, authMiddleware).Route()
	controller.NewSessionController(s.uc.SessionCaseUseCase(), rg, authMiddleware).Route()
	controller.NewAttendanceController(s.uc.AttendanceUseCase(), rg, authMiddleware).Route()
	controller.NewCsvController(s.uc.CsvCaseUseCase(s.csvService), rg).Route()
}

func (s *Server) Run() {
	s.setupControllers()
	// s.csvService.CreateFile()
	if err := s.engine.Run(s.host); err != nil {
		log.Fatal("server can't run")
	}
}

func NewServer() *Server {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	// gin.SetMode(gin.ReleaseMode)
	// engine := gin.New()

	infraManager, err := manager.NewInfraManager(cfg)
	if err != nil {
		log.Fatal(err)
	}

	repoManager := manager.NewRepoManager(infraManager)
	cvsService := common.NewCsvCommon(cfg.CsvFileConfig)
	useCaseManager := manager.NewUseCaseManager(repoManager, cvsService)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	logService := common.NewMyLogger(cfg.LogFileConfig)
	// cvsService := common.NewCsvCommon(cfg.CsvFileConfig)
	jwtService := common.NewJwtToken(cfg.TokenConfig)

	return &Server{
		uc:         useCaseManager,
		engine:     engine,
		host:       host,
		logService: logService,
		csvService: cvsService,
		auth:       usecase.NewAuthUseCase(useCaseManager.UserUseCase(), useCaseManager.StudentUseCase(), jwtService),
		jwtService: jwtService,
	}
}
