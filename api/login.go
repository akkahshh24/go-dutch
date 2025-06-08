package api

import (
	"net/http"

	"github.com/akkahshh24/go-dutch/util"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum,min=3,max=20"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
	Email       string `json:"email"`
}

func (s *Server) loginUser(c *gin.Context) {
	var req loginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := s.store.GetUserByName(c.Request.Context(), req.Username)
	if err != nil {
		if err == pgx.ErrNoRows {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if err := util.CheckPassword(req.Password, user.Password); err != nil {
		c.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	accessToken, err := s.tokenMaker.CreateToken(user.ID, s.config.AccessTokenDuration)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := loginResponse{
		AccessToken: accessToken,
		Username:    user.Username,
		Email:       user.Email,
	}

	c.JSON(http.StatusOK, rsp)
}
