package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/yamaniyuda/simplebank/db/sqlc"
)

type Server struct {
	store  *db.SQLStore
	router *gin.Engine
}

func NewServer(store *db.SQLStore) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router
	return server
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}

// Start runs the https server on a specific router
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
