package api

import (
	"net/http"

	db "github.com/akkahshh24/go-dutch/db/sqlc"
	"github.com/akkahshh24/go-dutch/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=20"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type createUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

func (s *Server) createUser(c *gin.Context) {
	var req createUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	user, err := s.store.CreateUser(c.Request.Context(), arg)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok {
			switch pgErr.Code {
			case "23505": // unique_violation:
				c.JSON(http.StatusConflict, errorResponse(err))
				return
			}
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := createUserResponse{
		Username: user.Username,
		Email:    user.Email,
	}
	c.JSON(http.StatusCreated, rsp)
}

type getUserByIDRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (s *Server) getUserByID(c *gin.Context) {
	var req getUserByIDRequest
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUserById(c.Request.Context(), req.ID)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, user)
}

type getUserByNameRequest struct {
	Name string `uri:"name" binding:"required,min=3,max=20"`
}

// func (s *Server) getUserByName(c *gin.Context) {
// 	var req getUserByNameRequest
// 	if err := c.ShouldBindUri(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, errorResponse(err))
// 		return
// 	}

// 	user, err := s.store.GetUserByName(c.Request.Context(), req.Name)
// 	if err != nil {
// 		if err == pgx.ErrNoRows {
// 			c.JSON(http.StatusNotFound, errorResponse(err))
// 			return
// 		}
// 		c.JSON(http.StatusInternalServerError, errorResponse(err))
// 		return
// 	}

// 	c.JSON(http.StatusOK, user)
// }

type listUsersRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (s *Server) listUsers(c *gin.Context) {
	var req listUsersRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	users, err := s.store.ListUsers(c.Request.Context(), arg)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, users)
}
