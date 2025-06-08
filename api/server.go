package api

import (
	"maps"

	db "github.com/akkahshh24/go-dutch/db/sqlc"
	"github.com/akkahshh24/go-dutch/token"
	"github.com/akkahshh24/go-dutch/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store *db.Store, tokenMaker token.Maker) *Server {
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	router := gin.Default()
	server.router = router

	server.setupRouter()

	return server
}

func (s *Server) setupRouter() {
	router := s.router

	router.POST("/users", s.createUser)
	router.POST("/users/login", s.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(s.tokenMaker))
	authRoutes.POST("/groups", s.createGroup)
	authRoutes.POST("/groups/:group_id/members", s.addGroupMember)
}

func (s *Server) Start(address string) error {
	return s.router.Run(address)
}

func successResponse(msg string) gin.H {
	return gin.H{"message": msg}
}

func successResponseWithData(msg string, data map[string]any) gin.H {
	response := gin.H{"message": msg}
	maps.Copy(response, data)
	return response
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
